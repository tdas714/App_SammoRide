package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/App-SammoRide/package/client"
)

// ride "github.com/App-SammoRide/chaincodes/Ride"

// "github.com/App-SammoRide/node"

func main() {
	InputYml := flag.String("in", "InputFile", "InfoFile")
	flag.Parse()

	inputinfo := client.Parse(*InputYml)

	node := client.NewNode(inputinfo.Path, inputinfo.ClientInfo)
	fmt.Println(node.RootCertificate)
	fmt.Println(node.Info.City, node.Info.Country, node.Info.Name)

	for {
		fmt.Println("Enter Your Command: ")
		var in string
		fmt.Scanln(&in)
		if strings.Contains(in, "join") {
			go node.JoinNetwork()
		} else if strings.Contains(in, "enroll") {
			node.CreateNode()
		} else if client.Contains([]string{"1", "2", "3"}, in) {
			var c *client.ClientInfo
			var i *client.InputInfo
			if in == "1" {
				i = client.Parse("ClientInfo/client_1.yml")
			} else if in == "2" {
				i = client.Parse("ClientInfo/client_2.yml")
			} else if in == "3" {
				i = client.Parse("ClientInfo/client_3.yml")
			} else {
				continue
			}
			c = i.ClientInfo
			if node.Info.Port == c.Port {
				continue
			}
			node.SendProposalToRider(*c, "Here", "There")
		} else if strings.Contains(in, "annon") {
			node.AnnounceAvailability()
		}
	}

}
