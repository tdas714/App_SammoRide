package main

import (
	"flag"
	"fmt"

	"github.com/App-SammoRide/client"
)

func main() {
	InputYml := flag.String("in", "InputFile", "InfoFile")
	flag.Parse()

	node := NewNode(*InputYml)

	fmt.Println(node.Info.City, node.Info.Country, node.Info.Name)
	// client.SendEnrollRequest(node.Info.Country, node.Info.Name,
	// 	node.Info.Province, node.Info.City, node.Info.Postalcode, client.GetIP())

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)

	go node.JoinNetwork()

	client.SendData(node.Info.IP, "CAs/rootCa.crt",
		cPath, kPath, "127.0.0.1", "8443", "hello", []byte(node.Info.Name))
}
