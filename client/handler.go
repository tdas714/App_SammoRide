package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GossipHandler(w http.ResponseWriter, r *http.Request, name string) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	CheckErr(err, "gossipRequest")

	// Write "Hello, world!" to the response body
	fmt.Println("Gossip from: ", r.RemoteAddr)
	io.WriteString(w, "Lets gossip!\n")
}

func RiderAHandler(w http.ResponseWriter, r *http.Request,
	c *ClientInfo, gossipList []string) []string {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	CheckErr(err, "RiderAnnonreq")
	gossip := GossipDeserialize(bodyBytes)
	if contains(gossipList, gossip.Header) {
		return
	}
	riderA := RADeserialize(gossip.Data)

	Gossip(bodyBytes, c, "Announcement/rider")
	gossipList = append(gossipList, gossip.Header)
	return gossipList
	// This will give info about showing drivers on map
}

func OrderProposalHandler(w http.ResponseWriter, resp *http.Request,
	name, arriveTime string, rideFair float32, info *ClientInfo) {

	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		info.Country, info.Name, info.Province,
		info.City)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	CheckErr(err, "OrderPrposal")

	keyPem, err := ioutil.ReadFile(kPath)
	CheckErr(err, "orderProposalHandler")

	info.PublicKey = nil

	contract := ContractDeserialize(bodyBytes)

	contract.Driver = info
	contract.RideFair = rideFair
	contract.ArrivalTime = arriveTime
	contract.TravelerSig = nil
	contract.DriverSig = nil

	hash := sha256.Sum256(contract.ContractSerialize())
	r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash[:])

	info.PublicKey = &LoadPrivateKey(keyPem).PublicKey
	contract.DriverSig = &Sig{r, s}
	w.Write(contract.ContractSerialize())
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
