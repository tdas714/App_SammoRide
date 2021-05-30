package main

import (
	"flag"
	"fmt"

	"github.com/App-SammoRide/client"
)

func main() {
	InputYml := flag.String("in", "InputFile", "InfoFile")
	flag.Parse()

	var c client.ClientInfo
	c.GetConf(*InputYml)

	fmt.Println(c.City, c.Country, c.Name)
	// client.SendEnrollRequest(c.Country, c.Name,
	// 	c.Province, c.City, c.Postalcode, client.GetIP())
	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		c.Country, c.Name, c.Province,
		c.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		c.Country, c.Name, c.Province,
		c.City)

	client.SendData(c.IP, "CAs/rootCa.crt",
		cPath, kPath, "127.0.0.1", "8443", "hello", []byte(c.Name))
}
