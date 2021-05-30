package main

import (
	"crypto/x509"

	"github.com/App-SammoRide/client"
)

type Node struct {
	Info        *client.ClientInfo
	PeerList    []string
	Certificate *x509.Certificate
}

func (node *Node) CreateNode() {
	// This will return Peer-list
	client.SendEnrollRequest(node.Info.Country, node.Info.Name,
		node.Info.Province, node.Info.City, node.Info.Postalcode,
		node.Info.IP, "127.0.0.1", "8080") //<--This will be dynamic
}

func (node *Node) JoinNetwork() {

}
