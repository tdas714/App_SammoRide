package node

import (
	"fmt"
	"strings"

	"github.com/App-SammoRide/client"
)

type Node struct {
	Info            *client.ClientInfo
	Connection      *client.Connections
	GossipData      []string
	Certificatepath string
	KeyPath         string
}

func NewNode(InputYml string) *Node {
	var c client.ClientInfo
	c.GetConf(InputYml)

	conn := client.LoadConnections(fmt.Sprintf("../database/%s/connections.bin", c.Name))
	defer conn.Close()

	var gossipData []string

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		c.Country, c.Name, c.Province,
		c.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		c.Country, c.Name, c.Province,
		c.City)

	node := Node{&c, conn, gossipData, cPath, kPath}
	return &node
}

func (node *Node) CreateNode() {
	// This will return Peer-list
	plist := client.SendEnrollRequest(node.Info.Country, node.Info.Name,
		node.Info.Province, node.Info.City, node.Info.Postalcode,
		node.Info.IP, node.Info.Port, "127.0.0.1", "8080") //<--This will be dynamic

	node.Connection.Add(plist)
}

func (node *Node) JoinNetwork() {

	client.StartPeerServer("CAs/interCa.crt", "CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, node.Info, node.GossipData)
	fmt.Println("Server Listing in port: ", node.Info.Port)
}

func (node *Node) AnnounceAvailability() {

	pList := node.Connection.GetRandom(1)

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)

	riderA := client.RiderAnnouncement{"lat", "long", "avail", node.Info}

	for _, p := range pList {
		splited := strings.Split(p, ":")

		client.SendData("CAs/rootCa.crt",
			cPath, kPath, splited[0], splited[1],
			"Announcement/rider", riderA.RASerialize(), 3)

	}

	client.SendData("CAs/rootCa.crt",
		cPath, kPath, "127.0.0.1", "8443",
		"/Announcement/rider", riderA.RASerialize(), 20)

}

func (node *Node) SendProposal(c client.ClientInfo) {
	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)

	orderProp := client.OrderContract{PickupLoc: "Currect Loc",
		DestLoc: "Destination Loc", Traveler: node.Info}
	// Have to get response
	resp := client.SendData("CAs/rootCa.crt",
		cPath, kPath, c.IP, c.Port,
		"OrderProposal", orderProp.ContractSerialize(), 10) // This will be real contract

	fmt.Println(client.ContractDeserialize(resp))
}
