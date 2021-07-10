package client

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	random "math/rand"
	"net/http"
	"time"

	"github.com/App-SammoRide/chaincodes/ride_1_0"
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
	fmt.Println("Received from: ", resp.RemoteAddr)
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
		timeStamp := time.Now()
		header.ChannelHeader.Timestamp = timeStamp

		chaincodepayload := peer.DeSerializeChaincodeProposalPayload(proposal.GetPayload())
		chaincodepayload.Input.ArrivalTime = time.Minute * time.Duration(arriveTime)
		chaincodepayload.Input.RideFair = rideFair
		chaincodepayload.Timeout = time.Duration(time.Minute * 2)

		fmt.Println(header)
		fmt.Println(chaincodepayload)
		respProposal := peer.Proposal{Header: header.Serialize(), Payload: chaincodepayload.Serialize()}

		r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash(respProposal.Serialize()))
		CheckErr(err, "sendtorider/sign")
		sig := peer.Sig{R: r, S: s}

		// signedprop.ProposalBytes = proposalbytes
		// signedprop.DriverSignature = sig.Serialize()
		// signedprop.DriverPublicKey = Keyencode(&LoadPrivateKey(keyPem).PublicKey)

		sendSignedprop := peer.SignedProposal{ProposalBytes: respProposal.Serialize(),
			DriverSignature: sig.Serialize(), DriverPublicKey: Keyencode(&LoadPrivateKey(keyPem).PublicKey),
			TravelerSignature: travelerSig.Serialize(), TravelerPublicKey: signedprop.TravelerPublicKey}

		// // fmt.Println(signedprop)
		w.Write(sendSignedprop.Serialize())
	}
}

// RESET SENDING PROPOSAL mAP
func Endorse(w http.ResponseWriter, resp *http.Request, node *Node) {
	fmt.Println("Recieved msg for endorse: ", resp.RemoteAddr)
	signedProposal := peer.DeSerializeSignedProposal(resp.Body)
	travelerSig := peer.DeSerializeSig(signedProposal.TravelerSignature)
	driverSig := peer.DeSerializeSig(signedProposal.DriverSignature)

	travelerV := ecdsa.Verify(Keydecode(signedProposal.TravelerPublicKey), hash(signedProposal.GetProposalBytes()), travelerSig.R, travelerSig.S)
	driverV := ecdsa.Verify(Keydecode(signedProposal.DriverPublicKey), hash(signedProposal.GetProposalBytes()), driverSig.R, driverSig.S)

	if travelerV && driverV {
		fmt.Println("verify")

		// certPem, err := ioutil.ReadFile(node.Certificatepath)
		// CheckErr(err, "OrderProposalhandler/CertPem")

		interPem, err := ioutil.ReadFile("CAs/interCa.crt")
		CheckErr(err, "Interca")

		proposal := peer.DeSerializeProposal(signedProposal.ProposalBytes)
		chaincodeProp := peer.DeSerializeChaincodeProposalPayload(proposal.GetPayload())
		header := common.DeSerializeHeader(proposal.GetHeader())
		fmt.Println("Added time: ", header.ChannelHeader.Timestamp.Add(time.Duration(chaincodeProp.GetTimeout())))
		fmt.Println(header.ChannelHeader.Timestamp)
		if time.Now().Before(header.ChannelHeader.Timestamp.Add(time.Duration(chaincodeProp.GetTimeout()))) {
			fmt.Println("Timeout")
			if VerifyPeer(node.RootCertificate, interPem, header.SignatureHeader.Driver) { //} && autheticate(node.RootCertificate, header.SignatureHeader.Traveler) {
				fmt.Println("Authenticate")
				driver := LoadCertificate(header.SignatureHeader.Driver)
				fmt.Println("Receive from: ", LoadCertificate(header.SignatureHeader.Traveler).IPAddresses)

				if driver.Subject.CommonName == "Driver" && time.Now().Before(driver.NotAfter) {

					fmt.Println("Invoke Chancode: ", chaincodeProp.ChaincodeId.GetName(), chaincodeProp.ChaincodeId.GetVersion())
					if chaincodeProp.ChaincodeId.GetName() == "ride" && chaincodeProp.ChaincodeId.GetVersion() == "1.0" {
						rwset, events, status := ride_1_0.StartRide(chaincodeProp.Input.PickupLocation, chaincodeProp.Input.DestinationLocation,
							chaincodeProp.Input.RideFair, header.ChannelHeader.TxId, signedProposal.DriverPublicKey,
							signedProposal.TravelerPublicKey, chaincodeProp.ChaincodeId)

						chainAction := peer.ChaincodeAction{Results: rwset, Events: events, ChaincodeId: chaincodeProp.ChaincodeId}
						proposalResPayload := peer.ProposalResponsePayload{ProposalHash: hash(signedProposal.GetProposalBytes()), Extension: &chainAction}

						proposalRes := peer.ProposalResponse{Timestamp: header.ChannelHeader.Timestamp, Response: &peer.Response{Status: int32(status)},
							Payload: proposalResPayload.Serialize()}

						certPem, err := ioutil.ReadFile(node.Certificatepath)
						CheckErr(err, "Endorsment/certPem")

						keyPem, err := ioutil.ReadFile(node.KeyPath)
						CheckErr(err, "endorsment/Keypem")

						r, s, err := ecdsa.Sign(rand.Reader, LoadPrivateKey(keyPem), hash(proposalResPayload.Serialize()))
						CheckErr(err, "sendtorider/sign")
						sig := peer.Sig{R: r, S: s}

						endorcer := peer.Endorsement{Endorser: certPem, Signature: &sig, PublicKey: Keyencode(&LoadPrivateKey(keyPem).PublicKey)}
						proposalRes.Endorsement = &endorcer
						fmt.Println("sending signed endorsement: ", chaincodeProp.IP, chaincodeProp.Port)
						SendData("CAs/rootCa.crt",
							node.Certificatepath, node.KeyPath, chaincodeProp.IP, chaincodeProp.Port,
							"Traveler/SignedEndorsement", proposalRes.Serialize(), 2)
					}
				}
			}
		}
	}
}

func EndorsementResponseHandler(w http.ResponseWriter, resp *http.Request, node *Node) {
	fmt.Println("Received endrosements")
	proposalRes := peer.DeSerializeProposalResponse(resp.Body)
	signedProp, ok := node.SentProposal[proposalRes.GetTimestamp()]
	fmt.Println(ok)
	if ok {
		fmt.Println("Activated")
		if proposalRes.Response.GetStatus() == 200 {
			verify := ecdsa.Verify(Keydecode(proposalRes.Endorsement.PublicKey), hash(proposalRes.Payload), proposalRes.Endorsement.Signature.R,
				proposalRes.Endorsement.Signature.S)
			if verify {
				fmt.Println("verified")
				propResPayload := peer.DeSerializeProposalResponsePayload(proposalRes.GetPayload())
				h := hash(signedProp.Serialize())
				if bytes.Compare(propResPayload.ProposalHash, h) == 1 {
					endorsement := node.ReceivedEndorsement[proposalRes.Timestamp]

					for _, e := range endorsement {
						if e != proposalRes.Endorsement {
							fmt.Println("Returned")
							return
						}
					}
					endorsement = append(endorsement, proposalRes.Endorsement)
					if len(endorsement) >= 1 {
						fmt.Println("Got endorsement")
						proposal := peer.DeSerializeProposal(signedProp.GetProposalBytes())

						chaincodeEndorsedAction := peer.ChaincodeEndorsedAction{ProposalResponsePayload: proposalRes.GetPayload(),
							Endorsements: endorsement}

						chaincodeacionPayload := peer.ChaincodeActionPayload{ChaincodeProposalPayload: proposal.GetPayload(),
							Action: &chaincodeEndorsedAction}

						transactionAcion := peer.TransactionAction{Header: proposal.GetHeader(), Payload: chaincodeacionPayload.Serialize()}

						transaction := peer.Transaction{Actions: []*peer.TransactionAction{&transactionAcion}}

						// SENDING TRANSACTION  TO ORDER FOR BLOCK CONSTRACTION
						SendData("CAs/rootCa.crt",
							node.Certificatepath, node.KeyPath, "127.0.0.1", "8443",
							"TransactionCommitmentRequest", transaction.Serialize(), 2)
						fmt.Println("Sent  to orderer")
					}
				}
			}
		}
	}
}

func BlockCommitmentHandler(w http.ResponseWriter, resp *http.Request, node *Node) {
	block := common.DeSerializeBlock(resp.Body)
	blockData := block.GetData()
	// Have mto do some Checking
	ok := node.WorldState.UpdateBlock(blockData, int(node.Blockchan.LastHeader.GetNumber()))
	if ok {
		node.Blockchan.Update(*block)
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

func Eq(a, b []interface{}) bool {
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
