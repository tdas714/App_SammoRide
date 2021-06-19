package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/App-SammoRide/client"
	"github.com/App-SammoRide/node"
)

func main() {
	InputYml := flag.String("in", "InputFile", "InfoFile")
	flag.Parse()

	node := node.NewNode(*InputYml)

	fmt.Println(node.Info.City, node.Info.Country, node.Info.Name)
	// client.SendEnrollRequest(node.Info.Country, node.Info.Name,
	// 	node.Info.Province, node.Info.City, node.Info.Postalcode, client.GetIP())

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		node.Info.Country, node.Info.Name, node.Info.Province,
		node.Info.City)

	// node.CreateNode()

	for {
		fmt.Println("Enter Your Command: ")
		var in string
		fmt.Scanln(&in)
		if strings.Contains(in, "join") {
			go node.JoinNetwork()
		} else if strings.Contains(in, "enroll") {
			node.CreateNode()
		} else if strings.Contains(in, ":") {
			splited := strings.Split(in, ":")
			client.SendData("CAs/rootCa.crt",
				cPath, kPath, "127.0.0.1", splited[1], "gossip",
				[]byte(node.Info.IP+":"+node.Info.Port), 10)

		} else if strings.Contains(in, "annon") {
			node.AnnounceAvailability()
		}
	}

}
