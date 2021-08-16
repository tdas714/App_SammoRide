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

	"github.com/App-SammoRide/package/client/struct/common"
	"github.com/App-SammoRide/package/client/struct/ledger"
	"github.com/App-SammoRide/package/client/struct/peer"
)

type ActivityType int32

const (
	FREE        ActivityType = 0
	NEGOTIATING ActivityType = 1
	ENGAGED     ActivityType = 2
)

type Node struct {
	Info                *ClientInfo
	Paths               *FilePath
	Connection          *Connections
	RootCertificate     []byte
	InterCertificate    []byte
	Certificatepath     string
	KeyPath             string
	GossipSentList      map[int64]string
	SentProposal        map[time.Time]*peer.SignedProposal
	ReceivedEndorsement map[time.Time][]*peer.Endorsement
	CommittmentCounter  int
	WorldState          *ledger.WorldState
	Blockchain          *ledger.Blockchain
	EndorsementPolicy   *common.ImplicitMetaPolicy
	NumberOfEndorsers   *common.SignaturePolicy
	ActivityStatus      ActivityType
}

func InitNode(paths *FilePath, c *ClientInfo) *Node {

	conn := LoadConnections(fmt.Sprintf("%s/connections.bin", paths.UtilsPath))
	defer conn.Close()

	gSL := make(map[int64]string)
	gSL[time.Now().Unix()] = "Orderer"

	fmt.Println("Cert Pamth: ", paths.CertificatePath)
	node := Node{Info: c, Connection: conn, Certificatepath: paths.CertificatePath, KeyPath: paths.KeyPath, Paths: paths,
		GossipSentList: gSL, SentProposal: make(map[time.Time]*peer.SignedProposal), RootCertificate: []byte(paths.RootCAPath), InterCertificate: []byte(paths.InterCAPath),
		ReceivedEndorsement: make(map[time.Time][]*peer.Endorsement), WorldState: ledger.Init(), Blockchain: ledger.LoadDatabase(paths.ChainPath),
		EndorsementPolicy: &common.ImplicitMetaPolicy{SubPolicy: common.Policy_PolicyType_name[1], Rule: common.ImplicitMetaPolicy_MAJORITY},
		NumberOfEndorsers: &common.SignaturePolicy{Type: &common.SignaturePolicy_SignedBy{SignedBy: 3}}, ActivityStatus: FREE}

	return &node
}

func NewNode(paths *FilePath, c *ClientInfo) *Node {

	conn := InitConnections(fmt.Sprintf("%s/connections.bin", paths.UtilsPath))
	defer conn.Close()

	gSL := make(map[int64]string)
	gSL[time.Now().Unix()] = "Orderer"

	blockchain := ledger.Blockchain{}
	blockchain.InitBlockchain(paths.ChainPath)
	defer blockchain.Close()
	node := Node{Info: c, Connection: conn, Certificatepath: paths.CertificatePath, KeyPath: paths.KeyPath, Paths: paths,
		GossipSentList: gSL, SentProposal: make(map[time.Time]*peer.SignedProposal),
		ReceivedEndorsement: make(map[time.Time][]*peer.Endorsement), WorldState: ledger.Init(), Blockchain: &blockchain,
		EndorsementPolicy: &common.ImplicitMetaPolicy{SubPolicy: common.Policy_PolicyType_name[1], Rule: common.ImplicitMetaPolicy_MAJORITY},
		NumberOfEndorsers: &common.SignaturePolicy{Type: &common.SignaturePolicy_SignedBy{SignedBy: 3}}, ActivityStatus: FREE}
	return &node
}

func (node *Node) Close() {
	node.Connection.Close()
	node.Blockchain.Close()
	node.WorldState.Close(node.Paths.UtilsPath + "/WorldState.json")
}

func (node *Node) CreateNode() {
	// This will return Peer-list and ordererlist
	plist, olist := SendEnrollRequest(node.Info.Country, node.Info.Name,
		node.Info.Province, node.Info.City, node.Info.Postalcode,
		node.Info.IP, node.Info.Port, "202.142.113.190", "8080", node.Paths, node) //<--This will be dynamic

	node.Connection.AddPeer(plist)
	node.Connection.AddOrderer(olist)
	node.Connection.Close()

}

func (node *Node) JoinNetwork() {

	StartPeerServer(node.Paths.InterCAPath, node.Paths.RootCAPath,
		node.Certificatepath, node.KeyPath, node)
	fmt.Println("Server Listing in port: ", node.Info.Port)
	node.ActivityStatus = FREE
}

func (node *Node) AnnounceAvailability() {
	// var splited []string
	// pList := node.Connection.GetRandom(1)

	riderA := RiderAnnouncement{Header: time.Now().Unix(), Latitude: "lat", Longitude: "long", Avalability: "avail", Info: node.Info}
	var splited []string
	if node.ActivityStatus == FREE {
		for _, o := range node.Connection.GetRandomOrderer(1) {
			splited = strings.Split(o, ":")
			SendData(node.Paths.RootCAPath,
				node.Certificatepath, node.KeyPath, splited[0], splited[1],
				"Announcement/rider", riderA.RASerialize(), 2)
		}
	}
	fmt.Println("Annoncement Sent")

}

func (node *Node) SendProposalToRider(c ClientInfo, pickup, des string) {

	node.Connection.AddPeer([]string{c.IP})

	fmt.Println("Sending proposal to: ", c.IP, " ", c.Port)

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	certPem, err := ioutil.ReadFile(node.Certificatepath)
	CheckErr(err, "Test")

	tbyte, err := time.Now().MarshalBinary()
	CheckErr(err, "SendProposal/tbyte")
	h := string(hash(tbyte))
	random.Seed(time.Now().Unix())

	cheader := common.ChannelHeader{ChannelId: string(common.Ride), TxId: h, Epoch: node.Blockchain.LastHeader.Number} // Change epoch based on block height
	sheader := common.SignatureHeader{Traveler: certPem, TravelerNonce: IntToByteArray(random.Int63())}
	Header := common.Header{ChannelHeader: &cheader, SignatureHeader: &sheader}
	propPayload := peer.ChaincodeProposalPayload{ChaincodeId: &peer.ChaincodeID{Name: "ride", Version: "1.0"},
		Input: &peer.ChaincodeInput{PickupLocation: pickup, DestinationLocation: des}, IPs: []string{node.Info.IP}, Ports: []string{node.Info.Port}}

	proposal := peer.Proposal{Header: Header.Serialize(), Payload: propPayload.Serialize()}

	r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash(proposal.Serialize()))
	CheckErr(err, "sendtorider/sign")
	sig := peer.Sig{R: r, S: s}

	signedProp := peer.SignedProposal{ProposalBytes: proposal.Serialize(),
		TravelerSignature: sig.Serialize(),
		TravelerPublicKey: Keyencode(&LoadPrivateKey(keyPem).PublicKey)}

	resp := SendData(node.Paths.RootCAPath, node.Certificatepath, node.KeyPath,
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
		node.ActivityStatus = NEGOTIATING
	}

}

//
// SEND TO ENDROSERS , RUN CHAIN CODE BASED ON decorations , ADD STRUCTURES FOR PROPOSAL RESPONSE
//

func (node *Node) SendProposalToEndorser(signedProp *peer.SignedProposal, c ClientInfo) {
	fmt.Println("Sending to endorser: ", c.IP+" "+"4000") // Have to change that to perform organically
	if int32(node.Connection.len()) >= node.NumberOfEndorsers.GetSignedBy() {

		var splited []string
		endorsList := node.Connection.GetRandomPeer(int(node.NumberOfEndorsers.GetSignedBy()))
		for _, endor := range endorsList {
			splited = strings.Split(endor, ":")
			fmt.Println(splited)
			SendData(node.Paths.RootCAPath,
				node.Certificatepath, node.KeyPath, c.IP, "4000",
				"Endorcers/SignedProposal", signedProp.Serialize(), 2)
		}
		node.ActivityStatus = ENGAGED
	}
}

func (node *Node) GetSnapshot(numSnap int64) {

	var resp []byte
	var splited []string
	for _, o := range node.Connection.GetRandomOrderer(1) {
		splited = strings.Split(o, ":")

		resp = SendData(node.Paths.RootCAPath,
			node.Certificatepath, node.KeyPath, splited[0], "8443",
			"Request/BlockSnapshot", IntToByteArray(numSnap), 2)

	}

	snapsEnv := common.DeSerializeSnapshotEnvelop(resp)

	if snapsEnv.Verify() {
		snapsBlocks := common.DeSerializeSnapshotBlocks(snapsEnv.Data)
		for _, block := range snapsBlocks.Blocks {
			if node.WorldState.UpdateBlock(block.GetData(), int(node.Blockchain.LastHeader.Number)) {
				node.Blockchain.Update(*block)
			}
		}
	}
}

func (node *Node) Gossip(header int64, num int, data []byte, ipAddr, domain string) {
	_, ok := node.GossipSentList[header]
	if !ok {
		var splited []string
		pList := node.Connection.GetRandomPeer(num)
		for _, p := range pList {
			if !strings.Contains(p, node.Info.IP+":"+node.Info.Port) {
				splited = strings.Split(p, ":")
				SendData(node.Paths.RootCAPath,
					node.Certificatepath, node.KeyPath, splited[0], splited[1],
					domain, data, 1)
			}
		}
		for _, o := range node.Connection.OrdererList {
			splited = strings.Split(o, ":")
			SendData(node.Paths.RootCAPath,
				node.Certificatepath, node.KeyPath, splited[0], "8443",
				domain, data, 1)

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
