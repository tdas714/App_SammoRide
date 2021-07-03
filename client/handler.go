package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	random "math/rand"
	"net/http"
	"time"

	ride "github.com/App-SammoRide/chaincodes/Ride"
	"github.com/App-SammoRide/struct/common"
	"github.com/App-SammoRide/struct/peer"
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
	// Have an interrested rider field to keep tract of interested rider

}

// Handles proposal from traveler
func TravelerOrderProposalHandler(w http.ResponseWriter, resp *http.Request,
	arriveTime int, rideFair float32, node *Node) {

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	certPem, err := ioutil.ReadFile(node.Certificatepath)
	CheckErr(err, "OrderProposalhandler/CertPem")

	signedprop := peer.DeSerializeSignedProposal(resp.Body)
	travelerSig := peer.DeSerializeSig(signedprop.TravelerSignature)
	proposalbytes := signedprop.GetProposalBytes()
	v := ecdsa.Verify(Keydecode(signedprop.TravelerPublicKey), hash(proposalbytes), travelerSig.R, travelerSig.S)

	if v {
		proposal := peer.DeSerializeProposal(proposalbytes)
		// Check Pickup location and Destination location + arival diatance
		// Note txId
		headerbytes := proposal.GetHeader()
		header := common.DeSerializeHeader(headerbytes)
		if header.ChannelHeader.ChannelId != string(common.Ride) {
			return
		}
		header.SignatureHeader.Driver = certPem
		header.SignatureHeader.DriverNonce = IntToByteArray(random.Int63())

		chaincodepayload := peer.DeSerializeChaincodeProposalPayload(proposal.GetPayload())
		chaincodepayload.Input.ArrivalTime = time.Minute * time.Duration(arriveTime)
		chaincodepayload.Input.RideFair = rideFair

		chaincodepayload.Input.Decorations = ride.GetRide()

		proposalbytes = proposal.Serialize()

		r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash(proposalbytes))
		CheckErr(err, "sendtorider/sign")
		sig := peer.Sig{R: r, S: s}

		signedprop = &peer.SignedProposal{ProposalBytes: proposalbytes,
			DriverSignature: sig.Serialize(),
			DriverPublicKey: Keyencode(&LoadPrivateKey(keyPem).PublicKey)}
	}
	fmt.Println(signedprop)
	w.Write(signedprop.Serialize())
}

func Endorse(w http.ResponseWriter, resp *http.Request, node *Node) {
	signedProposal := peer.DeSerializeSignedProposal(resp.Body)
	travelerSig := peer.DeSerializeSig(signedProposal.TravelerSignature)
	driverSig := peer.DeSerializeSig(signedProposal.DriverSignature)

	travelerV := ecdsa.Verify(Keydecode(signedProposal.TravelerPublicKey), hash(signedProposal.GetProposalBytes()), travelerSig.R, travelerSig.S)
	driverV := ecdsa.Verify(Keydecode(signedProposal.DriverPublicKey), hash(signedProposal.GetProposalBytes()), driverSig.R, driverSig.S)

	if travelerV && driverV {
		proposal := peer.DeSerializeProposal(signedProposal.ProposalBytes)
		chaincodeProp := peer.DeSerializeChaincodeProposalPayload(proposal.GetPayload())
		header := common.DeSerializeHeader(proposal.GetHeader())
		if time.Now().Before(header.ChannelHeader.Timestamp.Add(time.Duration(chaincodeProp.GetTimeout()))) {
			if autheticate(node.RootCertificate, header.SignatureHeader.Driver) && autheticate(node.RootCertificate, header.SignatureHeader.Traveler) {
				driver := LoadCertificate(header.SignatureHeader.Driver)
				if driver.Subject.CommonName == "Driver" {
					// Run smart convtract
				}
			}
		}
	}
}

func autheticate(rootCa, peerCa []byte) bool {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootCa))
	if !ok {
		panic("failed to parse root certificate")
	}

	cert := LoadCertificate(peerCa)
	// CheckErr(err, "VerifyOrderer/ParseCert")

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		return false
		// panic("failed to verify certificate: " + err.Error())
	}
	log.Print("Peer Verified")
	return true
}

func Eq(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
