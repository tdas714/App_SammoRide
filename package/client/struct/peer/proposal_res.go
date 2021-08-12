package peer

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"time"
)

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp time.Time
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement
}

func (m *ProposalResponse) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "ProposalResponse/Serialize")
	}
	return js
}

func DeSerializeProposalResponse(data io.Reader) *ProposalResponse {
	var m *ProposalResponse
	json.NewDecoder(data).Decode(&m)
	return m
}

func (m *ProposalResponse) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Now().AddDate(25, 0, 0)
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

type ResponseType int32

var (
	COMMITTED ResponseType = 0
	REJECTED  ResponseType = 1
)

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32
	// A message associated with the response code.
	Message string
	// A payload that can be used to include metadata with this response.
	Payload []byte
}

func (m *Response) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Response/Serialize")
	}
	return js
}

func DeSerializeResponse(data io.Reader) *Response {
	var m *Response
	json.NewDecoder(data).Decode(&m)
	return m
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte).
	ProposalHash []byte
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension *ChaincodeAction
}

func (m *ProposalResponsePayload) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "ProposalResponsePayload/Serialize")
	}
	return js
}

func DeSerializeProposalResponsePayload(data []byte) *ProposalResponsePayload {
	var m *ProposalResponsePayload
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}

func (m *ProposalResponsePayload) GetProposalHash() []byte {
	if m != nil {
		return m.ProposalHash
	}
	return nil
}

func (m *ProposalResponsePayload) GetExtension() *ChaincodeAction {
	if m != nil {
		return m.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature *Sig
	PublicKey string
}

func (m *Endorsement) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Endorsement/Serialize")
	}
	return js
}

func DeSerializeEndorsement(data io.Reader) *Endorsement {
	var m *Endorsement
	json.NewDecoder(data).Decode(&m)
	return m
}

func (m *Endorsement) GetEndorser() []byte {
	if m != nil {
		return m.Endorser
	}
	return nil
}

func (m *Endorsement) GetSignature() *Sig {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Endorsement) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

// FINDA WAY TO MAKE DELAYED PAYMENT, WITH UPI. cAN'T USE UPI IN GOLANG .RETURN AMOUNT IN INT.IN THE END OF RIDE COMPLETE. MONEY GETS PAIED TO
//EACH PERTICIPANTS.aT THE RIDE COMPLETE (RESOLUTION PHASE) DRIVER AND TRAVELER SENDS THE SIGNAL TO ORDRERER SERVICE AND ORDERER SERVICE SENDS THE MONEY
// TO PARTICILANTS

// ChaincodeAction contains the executed chaincode results, response, and event.
type ChaincodeAction struct {
	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte
	// This field contains the event generated by the chaincode.
	// Only a single marshaled ChaincodeEvent is included.
	Events []byte
	// This field contains the ChaincodeID of executing this invocation. Endorser
	// will set it with the ChaincodeID called by endorser while simulating proposal.
	// Committer will validate the version matching with latest chaincode version.
	// Adding ChaincodeID to keep version opens up the possibility of multiple
	// ChaincodeAction per transaction.
	ChaincodeId *ChaincodeID
}

func (m *ChaincodeAction) GetResults() []byte {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ChaincodeAction) GetEvents() []byte {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ChaincodeAction) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

//ChaincodeEvent is used for events and registrations that are specific to chaincode
//string type - "chaincode"
type ChaincodeEvent struct {
	ChaincodeId *ChaincodeID
	TxId        string
	EventName   string
	Payload     []byte
}

func (m *ChaincodeEvent) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeEvent) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChaincodeEvent) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *ChaincodeEvent) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EventsStruct struct {
	ChaincodeEvents []*ChaincodeEvent
}

func (m *EventsStruct) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Events/Serialize")
	}
	return js
}

func DeSerializeEvetns(data io.Reader) *EventsStruct {
	var m *EventsStruct
	json.NewDecoder(data).Decode(&m)
	return m
}
