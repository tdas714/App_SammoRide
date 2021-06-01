package node

import (
	"fmt"
	"strings"

	"github.com/App-SammoRide/client"
)

type Node struct {
	Info *client.ClientInfo
}

func NewNode(InputYml string) *Node {
	var c client.ClientInfo
	c.GetConf(InputYml)

	node := Node{&c}
	return &node
}

func (node *Node) CreateNode() {
	// This will return Peer-list
	client.SendEnrollRequest(node.Info.Country, node.Info.Name,
		node.Info.Province, node.Info.City, node.Info.Postalcode,
		node.Info.IP, node.Info.Port, "127.0.0.1", "8080") //<--This will be dynamic
}

func (node *Node) JoinNetwork() {

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)

	client.StartPeerServer(node.Info.Name, node.Info.IP, node.Info.Port, "CAs/interCa.crt",
		cPath, kPath)
	fmt.Println("Server Listing in port: ", node.Info.Port)
}

func (node *Node) AnnounceAvailability() {

	pList := client.GetPeerList(1, "database", node.Info.Name)

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
			"Announcement/rider", riderA.RASerialize())

	}

	client.SendData("CAs/rootCa.crt",
		cPath, kPath, "127.0.0.1", "8443",
		"Announcement/rider", riderA.RASerialize())

}

func (node *Node) SendProposal(c client.ClientInfo) {
	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)

	client.SendData("CAs/rootCa.crt",
		cPath, kPath, c.IP, c.Port,
		"proposal", []byte("Contact")) // This will be real contract
}
