package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
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

// Handles proposal from rider
func RiderOrderProposalHandler(w http.ResponseWriter, resp *http.Request,
	arriveTime string, rideFair float32, node *Node) {

	// kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
	// 	node.Info.Country, node.Info.Name, node.Info.Province,
	// 	node.Info.City)

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	node.Info.PublicKey = nil

	contract := ContractDeserialize(resp.Body)

	fmt.Println("Received Rider Order Proposal from: ",
		contract.Driver.IP+":"+contract.Driver.Port)

	node.Connection.Add([]string{contract.Driver.IP + ":" + contract.Driver.Port})

	contract.Driver = node.Info
	contract.RideFair = rideFair
	contract.ArrivalTime = arriveTime
	contract.TravelerSig = nil
	contract.DriverSig = nil

	hash := sha256.Sum256(contract.ContractSerialize())
	r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash[:])

	node.Info.PublicKey = &LoadPrivateKey(keyPem).PublicKey
	contract.DriverSig = &Sig{r, s}
	w.Write(contract.ContractSerialize())

	// Have to add gossip
}
