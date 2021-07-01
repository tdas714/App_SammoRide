package gateway

import (
	"github.com/App-SammoRide/structure/common"
	"github.com/App-SammoRide/structure/peer"
)

// EndorseRequest contains the details required to obtain sufficient endorsements for a
// transaction to be committed to the ledger.
type EndorseRequest struct {
	// The unique identifier for the transaction.
	TransactionId string
	// Identifier of the channel this request is bound for.
	ChannelId string
	// The signed proposal ready for endorsement.
	ProposedTransaction *peer.SignedProposal
	// If targeting the peers of specific organizations (e.g. for private data scenarios),
	// the list of organizations' MSPIDs should be supplied here.
	EndorsingOrganizations []string
}

func (m *EndorseRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *EndorseRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *EndorseRequest) GetProposedTransaction() *peer.SignedProposal {
	if m != nil {
		return m.ProposedTransaction
	}
	return nil
}

func (m *EndorseRequest) GetEndorsingOrganizations() []string {
	if m != nil {
		return m.EndorsingOrganizations
	}
	return nil
}

// EndorseResponse returns the result of endorsing a transaction.
type EndorseResponse struct {
	// The response that is returned by the transaction function, as defined
	// in peer/proposal_response.proto.
	Result *peer.Response
	// The unsigned set of transaction responses from the endorsing peers for signing by the client
	// before submitting to ordering service (via gateway).
	PreparedTransaction *common.Envelope
}

func (m *EndorseResponse) GetResult() *peer.Response {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EndorseResponse) GetPreparedTransaction() *common.Envelope {
	if m != nil {
		return m.PreparedTransaction
	}
	return nil
}

// SubmitRequest contains the details required to submit a transaction (update the ledger).
type SubmitRequest struct {
	// Identifier of the transaction to submit.
	TransactionId string
	// Identifier of the channel this request is bound for.
	ChannelId string
	// The signed set of endorsed transaction responses to submit.
	PreparedTransaction *common.Envelope
}

func (m *SubmitRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *SubmitRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SubmitRequest) GetPreparedTransaction() *common.Envelope {
	if m != nil {
		return m.PreparedTransaction
	}
	return nil
}

// SubmitResponse returns the result of submitting a transaction.
type SubmitResponse struct {
}

// SignedCommitStatusRequest contains a serialized CommitStatusRequest message, and a digital signature for the
// serialized request message.
type SignedCommitStatusRequest struct {
	// Serialized CommitStatusRequest message.
	Request []byte
	// Signature for request message.
	Signature []byte
}

func (m *SignedCommitStatusRequest) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SignedCommitStatusRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// CommitStatusRequest contains the details required to check whether a transaction has been
// successfully committed.
type CommitStatusRequest struct {
	// Identifier of the transaction to check.
	TransactionId string
	// Identifier of the channel this request is bound for.
	ChannelId string
	// Client requestor identity.
	Identity []byte
}

func (m *CommitStatusRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *CommitStatusRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *CommitStatusRequest) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

// CommitStatusResponse returns the result of committing a transaction.
type CommitStatusResponse struct {
	// The result of the transaction commit, as defined in peer/transaction.proto.
	Result peer.TxValidationCode
	// Block number that contains the transaction.
	BlockNumber uint64
}

func (m *CommitStatusResponse) GetResult() peer.TxValidationCode {
	if m != nil {
		return m.Result
	}
	return peer.TxValidationCode_VALID
}

func (m *CommitStatusResponse) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

// EvaluateRequest contains the details required to evaluate a transaction (query the ledger).
type EvaluateRequest struct {
	// Identifier of the transaction to evaluate.
	TransactionId string
	// Identifier of the channel this request is bound for.
	ChannelId string
	// The signed proposal ready for evaluation.
	ProposedTransaction *peer.SignedProposal
	// If targeting the peers of specific organizations (e.g. for private data scenarios),
	// the list of organizations' MSPIDs should be supplied here.
	TargetOrganizations []string
}

func (m *EvaluateRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *EvaluateRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *EvaluateRequest) GetProposedTransaction() *peer.SignedProposal {
	if m != nil {
		return m.ProposedTransaction
	}
	return nil
}

func (m *EvaluateRequest) GetTargetOrganizations() []string {
	if m != nil {
		return m.TargetOrganizations
	}
	return nil
}

// EvaluateResponse returns the result of evaluating a transaction.
type EvaluateResponse struct {
	// The response that is returned by the transaction function, as defined
	// in peer/proposal_response.proto.
	Result *peer.Response
}

func (m *EvaluateResponse) GetResult() *peer.Response {
	if m != nil {
		return m.Result
	}
	return nil
}

// SignedChaincodeEventsRequest contains a serialized ChaincodeEventsRequest message, and a digital signature for the
// serialized request message.
type SignedChaincodeEventsRequest struct {
	// Serialized ChaincodeEventsRequest message.
	Request []byte
	// Signature for request message.
	Signature []byte
}

func (m *SignedChaincodeEventsRequest) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SignedChaincodeEventsRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// ChaincodeEventsRequest contains details of the chaincode events that the caller wants to receive.
type ChaincodeEventsRequest struct {
	// Identifier of the channel this request is bound for.
	ChannelId string
	// Name of the chaincode for which events are requested.
	ChaincodeId string
	// Client requestor identity.
	Identity []byte
}

func (m *ChaincodeEventsRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChaincodeEventsRequest) GetChaincodeId() string {
	if m != nil {
		return m.ChaincodeId
	}
	return ""
}

func (m *ChaincodeEventsRequest) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

// ChaincodeEventsResponse returns chaincode events emitted from a specific block.
type ChaincodeEventsResponse struct {
	// Chaincode events emitted by the requested chaincode. The events are presented in the same order that the
	// transactions that emitted them appear within the block.
	Events []*peer.ChaincodeEvent
	// Block number in which the chaincode events were emitted.
	BlockNumber uint64
}

func (m *ChaincodeEventsResponse) GetEvents() []*peer.ChaincodeEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ChaincodeEventsResponse) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

// If any of the functions in the Gateway service returns an error, then it will be in the format of
// a google.rpc.Status message. The 'details' field of this message will be populated with extra
// information if the error is a result of one or more failed requests to remote peers or orderer nodes.
// EndpointError contains details of errors that are received by any of the endorsing peers
// as a result of processing the Evaluate or Endorse services, or from the ordering node(s) as a result of
// processing the Submit service.
type EndpointError struct {
	// The address of the endorsing peer or ordering node that returned an error.
	Address string
	// The MSP Identifier of this endpoint.
	MspId string
	// The error message returned by this endpoint.
	Message string
}

func (m *EndpointError) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EndpointError) GetMspId() string {
	if m != nil {
		return m.MspId
	}
	return ""
}

func (m *EndpointError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// ProposedTransaction contains the details required for offline signing prior to evaluating or endorsing
// a transaction.
type ProposedTransaction struct {
	// Identifier of the proposed transaction.
	TransactionId string
	// The signed proposal.
	Proposal *peer.SignedProposal
	// The list of endorsing organizations.
	EndorsingOrganizations []string
}

func (m *ProposedTransaction) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *ProposedTransaction) GetProposal() *peer.SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *ProposedTransaction) GetEndorsingOrganizations() []string {
	if m != nil {
		return m.EndorsingOrganizations
	}
	return nil
}

// PreparedTransaction contains the details required for offline signing prior to submitting a transaction.
type PreparedTransaction struct {
	// Identifier of the prepared transaction.
	TransactionId string
	// The transaction envelope.
	Envelope *common.Envelope
	// The response that is returned by the transaction function during endorsement, as defined
	// in peer/proposal_response.proto
	Result *peer.Response
}

func (m *PreparedTransaction) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *PreparedTransaction) GetEnvelope() *common.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

func (m *PreparedTransaction) GetResult() *peer.Response {
	if m != nil {
		return m.Result
	}
	return nil
}
