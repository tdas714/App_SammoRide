package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/App-SammoRide/client"
	// "github.com/App-SammoRide/node"
)

func main() {
	InputYml := flag.String("in", "InputFile", "InfoFile")
	flag.Parse()

	node := client.NewNode(*InputYml, "ClientDatabase")

	fmt.Println(node.Info.City, node.Info.Country, node.Info.Name)
	// client.SendEnrollRequest(node.Info.Country, node.Info.Name,
	// 	node.Info.Province, node.Info.City, node.Info.Postalcode, client.GetIP())

	// cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
	// 	node.Info.Country, node.Info.Name, node.Info.Province,
	// 	node.Info.City)
	// kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
	// 	node.Info.Country, node.Info.Name, node.Info.Province,
	// 	node.Info.City)

	// node.CreateNode()

	for {
		fmt.Println("Enter Your Command: ")
		var in string
		fmt.Scanln(&in)
		if strings.Contains(in, "join") {
			go node.JoinNetwork()
		} else if strings.Contains(in, "enroll") {
			node.CreateNode()
		} else if client.Contains([]string{"1", "2", "3"}, in) {
			var c client.ClientInfo
			if in == "1" {
				c.GetConf("ClientInfo/client_1.yml")
			} else if in == "2" {
				c.GetConf("ClientInfo/client_2.yml")
			} else if in == "3" {
				c.GetConf("ClientInfo/client_3.yml")
			} else {
				continue
			}

			node.RiderSendProposal(c)
		} else if strings.Contains(in, "annon") {
			node.AnnounceAvailability()
		}
	}

}
