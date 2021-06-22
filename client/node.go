package client

import (
	"fmt"
	"strings"
	"time"
	// "github.com/App-SammoRide/client"
)

type Node struct {
	Info            *ClientInfo
	Connection      *Connections
	Certificatepath string
	KeyPath         string
	GossipSentList  map[int64]string
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

	node := Node{Info: &c, Connection: conn, Certificatepath: cPath, KeyPath: kPath,
		GossipSentList: make(map[int64]string)}
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

	// for _, p := range pList {
	// 	splited = strings.Split(p, ":")
	// 	if !strings.Contains(p, node.Info.IP+":"+node.Info.Port) {
	// 		fmt.Println("Gossip Node: ", p)
	// 		SendData("CAs/rootCa.crt",
	// 			node.Certificatepath, node.KeyPath, splited[0], splited[1],
	// 			"Announcement/rider", riderA.RASerialize(), 3)
	// 	}

	// }

	SendData("CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, "127.0.0.1", "8443",
		"Announcement/rider", riderA.RASerialize(), 2)

}

func (node *Node) RiderSendProposal(c ClientInfo) {

	node.Connection.Add([]string{c.IP})

	orderProp := OrderContract{PickupLoc: "Currect Loc",
		DestLoc: "Destination Loc", Traveler: node.Info}
	// Have to get response
	resp := SendData("CAs/rootCa.crt",
		node.Certificatepath, node.KeyPath, c.IP, c.Port,
		"OrderProposal/Rider", orderProp.ContractSerialize(), 10) // This will be real contract

	fmt.Println(ContractDeserialize(resp))
}

func (node *Node) Gossip(header int64, num int, data []byte, ipAddr string) {
	_, ok := node.GossipSentList[header]
	if !ok {
		var splited []string
		pList := node.Connection.GetRandom(num)
		for _, p := range pList {
			splited = strings.Split(p, ":")
			SendData("CAs/rootCa.crt",
				node.Certificatepath, node.KeyPath, splited[0], splited[1],
				"OrderProposal/Rider", data, 1)
		}
		node.GossipSentList[header] = ipAddr
	}
}
