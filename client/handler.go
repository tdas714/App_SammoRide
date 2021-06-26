package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"
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
	arriveTime string, rideFair float32, node *Node) {

	keyPem, err := ioutil.ReadFile(node.KeyPath)
	CheckErr(err, "orderProposalHandler")

	certPem, err := ioutil.ReadFile(node.Certificatepath)
	CheckErr(err, "Test")

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
	contract.DriverCert = certPem

	hash := sha256.Sum256(contract.ContractSerialize())
	r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash[:])

	contract.DriverSig = []big.Int{*r, *s}

	fmt.Println(contract)
	w.Write(contract.ContractSerialize())
}

func TransactionProposalHandler(w http.ResponseWriter, r *http.Request, node *Node) {

	transRes := TransactionProposalResponse{}
	transProp := ContractDeserialize(r.Body)

	// Check for ktime diffrence from now and transaction proposal creation time
	timeDiff := time.Now().Sub(transProp.TimeStamp)
	if timeDiff >= node.EndorsmentPolicy.ProposalTimeDiff {
		transRes.Msg = "Session timed Out, Try again"
	}

	if transProp.TravelerSig != nil {
		transRes.Msg = "Traveler Signature Not Found!"
	}

	if transProp.DriverSig != nil {
		transRes.Msg = "Driver Signature Not Found!"
	}

	demoTransRes := transRes
	hash := sha256.Sum256(demoTransRes.TransResSerialize())

	verify := ecdsa.Verify(Keydecode(transProp.Traveler.PublicKey), hash[:],
		&transProp.TravelerSig[0], &transProp.TravelerSig[1])

	if !verify {
		transRes.Msg = "Signature is Invalid"
	}

	// Create writer policy to check driver has authorized for driving
	if node.WritersPolicy.DriverCheck != string(LoadCertificate(transProp.DriverCert).Subject.CommonName) {
		transRes.Msg = "Drivers Certificate Corrupted"
	}
	wrCheck := []bool{}
	wrCheck = append(wrCheck, autheticate(node.RootCertificate, transProp.DriverCert)) ///make Verify with root cert
	wrCheck = append(wrCheck, autheticate(node.RootCertificate, transProp.TravelerCert))
	if !Eq(node.WritersPolicy.CertificateValidation, wrCheck) {
		transRes.Msg = "Certificates are not valid"
	}

	// Here goes the Chain code evocation
	// Chekc for upi payment apis
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
