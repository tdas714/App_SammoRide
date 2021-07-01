package peer

import (
	"time"
)

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED             ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER              ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED            ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                  ChaincodeMessage_Type = 3
	ChaincodeMessage_READY                 ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION           ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED             ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR                 ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE             ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE             ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE             ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE      ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE              ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE    ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT      ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT      ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE     ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE             ChaincodeMessage_Type = 18
	ChaincodeMessage_GET_HISTORY_FOR_KEY   ChaincodeMessage_Type = 19
	ChaincodeMessage_GET_STATE_METADATA    ChaincodeMessage_Type = 20
	ChaincodeMessage_PUT_STATE_METADATA    ChaincodeMessage_Type = 21
	ChaincodeMessage_GET_PRIVATE_DATA_HASH ChaincodeMessage_Type = 22
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "GET_STATE_BY_RANGE",
	15: "GET_QUERY_RESULT",
	16: "QUERY_STATE_NEXT",
	17: "QUERY_STATE_CLOSE",
	18: "KEEPALIVE",
	19: "GET_HISTORY_FOR_KEY",
	20: "GET_STATE_METADATA",
	21: "PUT_STATE_METADATA",
	22: "GET_PRIVATE_DATA_HASH",
}

var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":             0,
	"REGISTER":              1,
	"REGISTERED":            2,
	"INIT":                  3,
	"READY":                 4,
	"TRANSACTION":           5,
	"COMPLETED":             6,
	"ERROR":                 7,
	"GET_STATE":             8,
	"PUT_STATE":             9,
	"DEL_STATE":             10,
	"INVOKE_CHAINCODE":      11,
	"RESPONSE":              13,
	"GET_STATE_BY_RANGE":    14,
	"GET_QUERY_RESULT":      15,
	"QUERY_STATE_NEXT":      16,
	"QUERY_STATE_CLOSE":     17,
	"KEEPALIVE":             18,
	"GET_HISTORY_FOR_KEY":   19,
	"GET_STATE_METADATA":    20,
	"PUT_STATE_METADATA":    21,
	"GET_PRIVATE_DATA_HASH": 22,
}

type ChaincodeMessage struct {
	Type      ChaincodeMessage_Type
	Timestamp time.Time
	Payload   []byte
	Txid      string
	Proposal  *SignedProposal
	//event emitted by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	//with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent
	//channel id
	ChannelId string
}

func (m *ChaincodeMessage) GetType() ChaincodeMessage_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeMessage_UNDEFINED
}

func (m *ChaincodeMessage) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Now().AddDate(25, 0, 0)
}

func (m *ChaincodeMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChaincodeMessage) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ChaincodeMessage) GetProposal() *SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

func (m *ChaincodeMessage) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// GetState is the payload of a ChaincodeMessage. It contains a key which
// is to be fetched from the ledger. If the collection is specified, the key
// would be fetched from the collection (i.e., private state)
type GetState struct {
	Key        string
	Collection string
}

func (m *GetState) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetState) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type GetStateMetadata struct {
	Key        string
	Collection string
}

func (m *GetStateMetadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetStateMetadata) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

// PutState is the payload of a ChaincodeMessage. It contains a key and value
// which needs to be written to the transaction's write set. If the collection is
// specified, the key and value would be written to the transaction's private
// write set.
type PutState struct {
	Key        string
	Value      []byte
	Collection string
}

func (m *PutState) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutState) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PutState) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type PutStateMetadata struct {
	Key        string
	Collection string
	Metadata   *StateMetadata
}

func (m *PutStateMetadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutStateMetadata) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *PutStateMetadata) GetMetadata() *StateMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// DelState is the payload of a ChaincodeMessage. It contains a key which
// needs to be recorded in the transaction's write set as a delete operation.
// If the collection is specified, the key needs to be recorded in the
// transaction's private write set as a delete operation.
type DelState struct {
	Key        string
	Collection string
}

func (m *DelState) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DelState) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

// GetStateByRange is the payload of a ChaincodeMessage. It contains a start key and
// a end key required to execute range query. If the collection is specified,
// the range query needs to be executed on the private data. The metadata hold
// the byte representation of QueryMetadata.
type GetStateByRange struct {
	StartKey   string
	EndKey     string
	Collection string
	Metadata   []byte
}

func (m *GetStateByRange) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GetStateByRange) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *GetStateByRange) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *GetStateByRange) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// GetQueryResult is the payload of a ChaincodeMessage. It contains a query
// string in the form that is supported by the underlying state database.
// If the collection is specified, the query needs to be executed on the
// private data.  The metadata hold the byte representation of QueryMetadata.
type GetQueryResult struct {
	Query      string
	Collection string
	Metadata   []byte
}

func (m *GetQueryResult) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *GetQueryResult) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *GetQueryResult) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// QueryMetadata is the metadata of a GetStateByRange and GetQueryResult.
// It contains a pageSize which denotes the number of records to be fetched
// and a bookmark.
type QueryMetadata struct {
	PageSize int32
	Bookmark string
}

func (m *QueryMetadata) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryMetadata) GetBookmark() string {
	if m != nil {
		return m.Bookmark
	}
	return ""
}

// GetHistoryForKey is the payload of a ChaincodeMessage. It contains a key
// for which the historical values need to be retrieved.
type GetHistoryForKey struct {
	Key string
}

func (m *GetHistoryForKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryStateNext struct {
	Id string
}

func (m *QueryStateNext) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryStateClose struct {
	Id string
}

func (m *QueryStateClose) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryResultBytes hold the byte representation of a record returned by the peer.
type QueryResultBytes struct {
	ResultBytes []byte
}

func (m *QueryResultBytes) GetResultBytes() []byte {
	if m != nil {
		return m.ResultBytes
	}
	return nil
}

// QueryResponse is returned by the peer as a result of a GetStateByRange,
// GetQueryResult, and GetHistoryForKey. It holds a bunch of records in
// results field, a flag to denote whether more results need to be fetched from
// the peer in has_more field, transaction id in id field, and a QueryResponseMetadata
// in metadata field.
type QueryResponse struct {
	Results  []*QueryResultBytes
	HasMore  bool
	Id       string
	Metadata []byte
}

func (m *QueryResponse) GetResults() []*QueryResultBytes {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *QueryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryResponse) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// QueryResponseMetadata is the metadata of a QueryResponse. It contains a count
// which denotes the number of records fetched from the ledger and a bookmark.
type QueryResponseMetadata struct {
	FetchedRecordsCount int32
	Bookmark            string
}

func (m *QueryResponseMetadata) GetFetchedRecordsCount() int32 {
	if m != nil {
		return m.FetchedRecordsCount
	}
	return 0
}

func (m *QueryResponseMetadata) GetBookmark() string {
	if m != nil {
		return m.Bookmark
	}
	return ""
}

type StateMetadata struct {
	Metakey string
	Value   []byte
}

func (m *StateMetadata) GetMetakey() string {
	if m != nil {
		return m.Metakey
	}
	return ""
}

func (m *StateMetadata) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type StateMetadataResult struct {
	Entries []*StateMetadata
}

func (m *StateMetadataResult) GetEntries() []*StateMetadata {
	if m != nil {
		return m.Entries
	}
	return nil
}
