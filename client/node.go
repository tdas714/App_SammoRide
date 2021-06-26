package client

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	"github.com/App-SammoRide/policy"
	// "github.com/App-SammoRide/client"
)

type Node struct {
	Info             *ClientInfo
	Connection       *Connections
	RootCertificate  []byte
	Certificatepath  string
	KeyPath          string
	GossipSentList   map[int64]string
	EndorsmentPolicy *policy.EndorsmentPolicy
	WritersPolicy    *policy.WritersPolicy
}

func NewNode(InputYml, dir string) *Node {
	var c ClientInfo
	c.GetConf(InputYml)

	filepath := fmt.Sprintf("../%s/%s", dir, strings.Split(c.Name, ".")[0])
	CreateDirIfNotExist(filepath)

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
		GossipSentList: gSL, EndorsmentPolicy: endors, WritersPolicy: writers, RootCertificate: rootca}
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

func (node *Node) SendProposalToRider(c ClientInfo, loc, des string) {

	node.Connection.Add([]string{c.IP})

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	certPem, err := ioutil.ReadFile(node.Certificatepath)
	CheckErr(err, "Test")

	orderProp := TransactionProposal{PickupLoc: "Currect Loc",
		DestLoc: "Destination Loc", Traveler: node.Info, TravelerCert: certPem}

	fmt.Println("Sendin proposal to: ", c.IP+":"+c.Port)
	// Have to get response
	resp := SendData("CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, c.IP, c.Port,
		"OrderProposal/Traveler", orderProp.ContractSerialize(), 10) // This will be real contract
	// HTTP response mto structure
	// respString := fmt.Sprintf("%s", resp)
	transProp := ContractDeserialize(bytes.NewBuffer(resp))
	demoContract := *transProp
	demoContract.DriverSig = nil
	h := hash(demoContract.ContractSerialize())[:]
	publicKey := Keydecode(transProp.Driver.PublicKey)
	// fmt.Println(contract.Driver.PublicKey, h,
	// 	&contract.DriverSig[0], &contract.DriverSig[1])

	verify := ecdsa.Verify(publicKey, h,
		&transProp.DriverSig[0], &transProp.DriverSig[1])

	if verify {
		transProp.PickupLoc = loc
		transProp.DestLoc = des
		node.Info.PublicKey = Keyencode(&LoadPrivateKey(keyPem).PublicKey)
		transProp.Traveler = node.Info
		transProp.TravelerSig = nil
		transProp.Type = StartRide
		transProp.TimeStamp = time.Now()

		hash := sha256.Sum256(transProp.ContractSerialize())
		r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash[:])
		CheckErr(err, "SendriderProposal/Sig")

		transProp.TravelerSig = []big.Int{*r, *s}

		if node.Connection.len() >= node.EndorsmentPolicy.NumberOfEndorsers {
			var splited []string
			endorsList := node.Connection.GetRandom(node.EndorsmentPolicy.NumberOfEndorsers)
			for _, endor := range endorsList {
				splited = strings.Split(endor, ":")

				SendData("CAs/interCa.crt", node.Certificatepath, node.KeyPath,
					splited[0], splited[1], "Transaction/Proposal", transProp.ContractSerialize(),
					3)
			}
		} else {
			SendData("CAs/rootCa.crt",
				node.Certificatepath, node.KeyPath, "127.0.0.1", "8443",
				"Transaction/Proposal", transProp.ContractSerialize(), 2)
		}

	}

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
