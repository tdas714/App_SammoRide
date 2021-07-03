package peer

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"time"
)

// When an endorser receives a SignedProposal message, it should verify the
// signature over the proposal bytes. This verification requires the following
// steps:
// 1. Verification of the validity of the certificate that was used to produce
//    the signature.  The certificate will be available once proposalBytes has
//    been unmarshalled to a Proposal message, and Proposal.header has been
//    unmarshalled to a Header message. While this unmarshalling-before-verifying
//    might not be ideal, it is unavoidable because i) the signature needs to also
//    protect the signing certificate; ii) it is desirable that Header is created
//    once by the client and never changed (for the sake of accountability and
//    non-repudiation). Note also that it is actually impossible to conclusively
//    verify the validity of the certificate included in a Proposal, because the
//    proposal needs to first be endorsed and ordered with respect to certificate
//    expiration transactions. Still, it is useful to pre-filter expired
//    certificates at this stage.
// 2. Verification that the certificate is trusted (signed by a trusted CA) and
//    that it is allowed to transact with us (with respect to some ACLs);
// 3. Verification that the signature on proposalBytes is valid;
// 4. Detect replay attacks;
type SignedProposal struct {
	// The bytes of Proposal
	ProposalBytes []byte
	// Signaure over proposalBytes; this signature is to be verified against
	// the creator identity contained in the header of the Proposal message
	// marshaled as proposalBytes
	DriverSignature   []byte
	TravelerSignature []byte
	DriverPublicKey   string
	TravelerPublicKey string
}

func (m *SignedProposal) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "signedProposal/Serialize")
	}
	return js
}

func DeSerializeSignedProposal(data io.Reader) *SignedProposal {
	var m *SignedProposal
	json.NewDecoder(data).Decode(&m)
	return m
}

func (m *SignedProposal) GetProposalBytes() []byte {
	if m != nil {
		return m.ProposalBytes
	}
	return nil
}

func (m *SignedProposal) GetDriverSignature() *Sig {
	if m != nil {
		return DeSerializeSig(m.DriverSignature)
	}
	return nil
}

func (m *SignedProposal) GetTravelerSignature() *Sig {
	if m != nil {
		return DeSerializeSig(m.DriverSignature)
	}
	return nil
}

type Sig struct {
	R *big.Int
	S *big.Int
}

func (m *Sig) GetR() *big.Int {
	return m.R
}

func (m *Sig) GetS() *big.Int {
	return m.S
}

func (m *Sig) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "sig/Serialize")
	}
	return js
}

func DeSerializeSig(data []byte) *Sig {
	var m *Sig
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}

// A Proposal is sent to an endorser for endorsement.  The proposal contains:
// 1. A header which should be unmarshaled to a Header message.  Note that
//    Header is both the header of a Proposal and of a Transaction, in that i)
//    both headers should be unmarshaled to this message; and ii) it is used to
//    compute cryptographic hashes and signatures.  The header has fields common
//    to all proposals/transactions.  In addition it has a type field for
//    additional customization. An example of this is the ChaincodeHeaderExtension
//    message used to extend the Header for type CHAINCODE.
// 2. A payload whose type depends on the header's type field.
// 3. An extension whose type depends on the header's type field.
//
// Let us see an example. For type CHAINCODE (see the Header message),
// we have the following:
// 1. The header is a Header message whose extensions field is a
//    ChaincodeHeaderExtension message.
// 2. The payload is a ChaincodeProposalPayload message.
// 3. The extension is a ChaincodeAction that might be used to ask the
//    endorsers to endorse a specific ChaincodeAction, thus emulating the
//    submitting peer model.
type Proposal struct {
	// The header of the proposal. It is the bytes of the Header
	Header []byte
	// The payload of the proposal as defined by the type in the proposal
	// header.
	Payload []byte
}

func (m *Proposal) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "proposal/Serialize")
	}
	return js
}

func DeSerializeProposal(data []byte) *Proposal {
	var m *Proposal
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}

func (m *Proposal) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Proposal) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	// The ID of the chaincode to target.
	ChaincodeId *ChaincodeID
}

func (m *ChaincodeHeaderExtension) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

// ChaincodeProposalPayload is the Proposal's payload message to be used when
// the Header's type is CHAINCODE.  It contains the arguments for this
// invocation.
type ChaincodeProposalPayload struct {
	// Input contains the arguments for this invocation. If this invocation
	// deploys a new chaincode, ESCC/VSCC are part of this field.
	// This is usually a marshaled ChaincodeInvocationSpec

	ChaincodeId *ChaincodeID
	Input       *ChaincodeInput
	Timeout     int32
}

func (m *ChaincodeProposalPayload) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "ChainCodproppayload/Serialize")
	}
	return js
}

func DeSerializeChaincodeProposalPayload(data []byte) *ChaincodeProposalPayload {
	var m *ChaincodeProposalPayload
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}

func (m *ChaincodeProposalPayload) GetChaincodeID() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeProposalPayload) GetChaincodeInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeProposalPayload) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

//ChaincodeID contains the path as specified by the deploy transaction
//that created it as well as the hashCode that is generated by the
//system for the path. From the user level (ie, CLI, REST API and so on)
//deploy transaction is expected to provide the path and other requests
//are expected to provide the hashCode. The other value will be ignored.
//Internally, the structure could contain both values. For instance, the
//hashCode will be set when first generated using the path
type ChaincodeID struct {
	//all other requests will use the name (really a hashcode) generated by
	//the deploy transaction
	Name string
	//user friendly version name for the chaincode
	Version string
}

func (m *ChaincodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.Args are inputs
type ChaincodeInput struct {
	PickupLocation      string
	DestinationLocation string
	RideFair            float32
	ArrivalTime         time.Duration
	Decorations         map[string]interface{}
	Hash                []byte
}

func (m *ChaincodeInput) GetPickupLocation() string {
	if m != nil {
		return m.PickupLocation
	}
	return ""
}
func (m *ChaincodeInput) GetDestinationLocation() string {
	if m != nil {
		return m.DestinationLocation
	}
	return ""
}
func (m *ChaincodeInput) GetRideFair() float32 {
	if m != nil {
		return m.RideFair
	}
	return 0
}
func (m *ChaincodeInput) GetArrivalTime() time.Duration {
	if m != nil {
		return m.ArrivalTime
	}
	return time.Minute * -10
}
func (m *ChaincodeInput) GetDecoration() interface{} {
	if m != nil {
		return m.Decorations
	}
	return nil
}
