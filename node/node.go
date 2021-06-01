package node

import (
	"fmt"

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
