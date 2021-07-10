package client

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	random "math/rand"
	"strings"
	"time"

	"github.com/App-SammoRide/policy"
	"github.com/App-SammoRide/struct/common"
	"github.com/App-SammoRide/struct/ledger"
	"github.com/App-SammoRide/struct/peer"
)

type Node struct {
	Info                *ClientInfo
	Connection          *Connections
	RootCertificate     []byte
	Certificatepath     string
	KeyPath             string
	GossipSentList      map[int64]string
	EndorsmentPolicy    *policy.EndorsmentPolicy
	WritersPolicy       *policy.WritersPolicy
	SentProposal        map[time.Time]*peer.SignedProposal
	ReceivedEndorsement map[time.Time][]*peer.Endorsement
	WorldState          *ledger.WorldState
	Blockchan           *ledger.Blockchain
}

func NewNode(InputYml, dir string) *Node {
	var c ClientInfo
	c.GetConf(InputYml)

	filepath := fmt.Sprintf("../%s/%s", dir, strings.Split(c.Name, ".")[0])
	CreateDirIfNotExist(filepath)

	chainPath := fmt.Sprintf("../%s/%s", dir, "Chain")
	CreateDirIfNotExist(chainPath)

	conn := LoadConnections(fmt.Sprintf("%s/connections.bin", filepath))
	defer conn.Close()

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		c.Country, c.Name, c.Province,
		c.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		c.Country, c.Name, c.Province,
		c.City)
	gSL := make(map[int64]string)
	gSL[time.Now().Unix()] = "Orderer"

	rootca, err := ioutil.ReadFile("CAs/rootCa.crt")
	CheckErr(err, "Node/RootCa")

	endors := policy.GetEndorsmentPolicy()
	writers := policy.GetWritersPolicy()
	node := Node{Info: &c, Connection: conn, Certificatepath: cPath, KeyPath: kPath,
		GossipSentList: gSL, EndorsmentPolicy: endors, WritersPolicy: writers, RootCertificate: rootca, SentProposal: make(map[time.Time]*peer.SignedProposal),
		ReceivedEndorsement: make(map[time.Time][]*peer.Endorsement), WorldState: ledger.Init(), Blockchan: ledger.LoadDatabase(chainPath)}
	return &node
}

func (node *Node) Close() {
	node.Connection.Close()
}

func (node *Node) CreateNode() {
	// This will return Peer-list
	plist := SendEnrollRequest(node.Info.Country, node.Info.Name,
		node.Info.Province, node.Info.City, node.Info.Postalcode,
		node.Info.IP, node.Info.Port, "127.0.0.1", "8080") //<--This will be dynamic

	node.Connection.Add(plist)
	node.Connection.Close()
}

func (node *Node) JoinNetwork() {

	StartPeerServer("CAs/interCa.crt", "CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, node)
	fmt.Println("Server Listing in port: ", node.Info.Port)
}

func (node *Node) AnnounceAvailability() {
	// var splited []string
	// pList := node.Connection.GetRandom(1)

	riderA := RiderAnnouncement{Header: time.Now().Unix(), Latitude: "lat", Longitude: "long", Avalability: "avail", Info: node.Info}

	SendData("CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, "127.0.0.1", "8443",
		"Announcement/rider", riderA.RASerialize(), 2)
	fmt.Println("Annoncement Sent")

}

func (node *Node) SendProposalToRider(c ClientInfo, pickup, des string) {

	node.Connection.Add([]string{c.IP})

	fmt.Println("Sending proposal to: ", c.IP, " ", c.Port)

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	certPem, err := ioutil.ReadFile(node.Certificatepath)
	CheckErr(err, "Test")

	tbyte, err := time.Now().MarshalBinary()
	CheckErr(err, "SendProposal/tbyte")
	h := string(hash(tbyte))
	random.Seed(time.Now().Unix())

	cheader := common.ChannelHeader{ChannelId: string(common.Ride), TxId: h, Epoch: 1} // Change epoch based on block height
	sheader := common.SignatureHeader{Traveler: certPem, TravelerNonce: IntToByteArray(random.Int63())}
	Header := common.Header{ChannelHeader: &cheader, SignatureHeader: &sheader}
	propPayload := peer.ChaincodeProposalPayload{ChaincodeId: &peer.ChaincodeID{Name: "ride", Version: "1.0"},
		Input: &peer.ChaincodeInput{PickupLocation: pickup, DestinationLocation: des}, IP: node.Info.IP, Port: node.Info.Port}

	proposal := peer.Proposal{Header: Header.Serialize(), Payload: propPayload.Serialize()}

	r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash(proposal.Serialize()))
	CheckErr(err, "sendtorider/sign")
	sig := peer.Sig{R: r, S: s}

	signedProp := peer.SignedProposal{ProposalBytes: proposal.Serialize(),
		TravelerSignature: sig.Serialize(),
		TravelerPublicKey: Keyencode(&LoadPrivateKey(keyPem).PublicKey)}

	resp := SendData("CAs/interCa.crt", node.Certificatepath, node.KeyPath,
		c.IP, c.Port, "Transaction/Proposal", signedProp.Serialize(),
		10)

	RespSignedProp := peer.DeSerializeSignedProposal(bytes.NewBuffer(resp))

	propb := RespSignedProp.GetProposalBytes()
	prop := peer.DeSerializeProposal(propb)
	header := common.DeSerializeHeader(prop.GetHeader())
	fmt.Println(header.ChannelHeader.Timestamp)

	driverSig := peer.DeSerializeSig(RespSignedProp.DriverSignature)
	v := ecdsa.Verify(Keydecode(RespSignedProp.DriverPublicKey), hash(RespSignedProp.GetProposalBytes()), driverSig.R, driverSig.S)

	if v {
		r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash(RespSignedProp.GetProposalBytes()))
		CheckErr(err, "RespProposal")
		sig := peer.Sig{R: r, S: s}
		RespSignedProp.TravelerSignature = sig.Serialize()
		// We have to unserialize proposalbytes for checking the ride fair and arrival time, chancode decorations
		// Send this to endrosers as per endorcement policy
		node.SentProposal[header.ChannelHeader.Timestamp] = RespSignedProp
		node.SendProposalToEndorser(RespSignedProp, c)
	}

}

//
// SEND TO ENDROSERS , RUN CHAIN CODE BASED ON decorations , ADD STRUCTURES FOR PROPOSAL RESPONSE
//

func (node *Node) SendProposalToEndorser(signedProp *peer.SignedProposal, c ClientInfo) {
	fmt.Println("Sending to endorser: ", c.IP+" "+"4000")
	// if node.Connection.len() >= node.EndorsmentPolicy.NumberOfEndorsers {
	// 	var splited []string
	// 	endorsList := node.Connection.GetRandom(node.EndorsmentPolicy.NumberOfEndorsers)
	// 	for _, endor := range endorsList {
	// splited = strings.Split(endor, ":")
	SendData("CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, c.IP, "4000",
		"Endorcers/SignedProposal", signedProp.Serialize(), 2)
	// 	}
	// }
}

func (node *Node) Gossip(header int64, num int, data []byte, ipAddr, domain string) {
	_, ok := node.GossipSentList[header]
	if !ok {
		var splited []string
		pList := node.Connection.GetRandom(num)
		for _, p := range pList {
			if !strings.Contains(p, node.Info.IP+":"+node.Info.Port) {
				splited = strings.Split(p, ":")
				SendData("CAs/rootCa.crt",
					node.Certificatepath, node.KeyPath, splited[0], splited[1],
					domain, data, 1)
			}
		}
		node.GossipSentList[header] = ipAddr
	}
}

func hash(b []byte) []byte {
	h := sha256.New()
	// hash the body bytes
	h.Write(b)
	// compute the SHA256 hash
	return h.Sum(nil)
}
