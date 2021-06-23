package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

// func GossipHandler(w http.ResponseWriter, r *http.Request, name string) {
// 	bodyBytes, err := ioutil.ReadAll(r.Body)
// 	CheckErr(err, "gossipRequest")

// 	// Write "Hello, world!" to the response body
// 	fmt.Println("Gossip from: ", r.RemoteAddr)
// 	io.WriteString(w, "Lets gossip!\n")
// }

func RiderAHandler(w http.ResponseWriter, r *http.Request,
	node *Node) {

	riderA := RADeserialize(r.Body)
	node.Connection.Add([]string{riderA.Info.IP + ":" + riderA.Info.Port})
	node.Gossip(riderA.Header, 1, riderA.RASerialize(), riderA.Info.IP, "Announcement/rider")

	fmt.Print("Rider Announcment from: ", riderA.Info.IP+":"+riderA.Info.Port)

}

// Handles proposal from traveler
func TravelerOrderProposalHandler(w http.ResponseWriter, resp *http.Request,
	arriveTime string, rideFair float32, node *Node) {

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	contract := ContractDeserialize(resp.Body)

	fmt.Println("Received Rider Order Proposal from: ",
		contract.Traveler.IP+":"+contract.Traveler.Port)

	node.Connection.Add([]string{contract.Traveler.IP + ":" + contract.Traveler.Port})

	contract.Driver = node.Info
	contract.RideFair = rideFair
	contract.ArrivalTime = arriveTime
	contract.TravelerSig = nil
	contract.DriverSig = nil
	node.Info.PublicKey = Keyencode(&LoadPrivateKey(keyPem).PublicKey)
	contract.Driver = node.Info

	hash := sha256.Sum256(contract.ContractSerialize())
	r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash[:])

	contract.DriverSig = []big.Int{*r, *s}

	fmt.Println(contract)
	w.Write(contract.ContractSerialize())

	// HAVE TO ADD GOSSIP HERE <<<<<<===============
}
