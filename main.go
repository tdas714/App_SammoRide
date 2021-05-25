package main

import "github.com/App-SammoRide/client"

func main() {
	// client.SendEnrollRequest("India", "Tapas.Das", "west Bengal", "kolkata", "700028", client.GetIP())
	client.SendData("127.0.0.1", "CAs/rootCa.crt", "PeerCerts/Cert.crt", "PeerCerts/Cert.key")
}
