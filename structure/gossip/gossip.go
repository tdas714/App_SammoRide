package gosip

import "github.com/App-SammoRide/structure/peer"

type PullMsgType int32

const (
	PullMsgType_UNDEFINED    PullMsgType = 0
	PullMsgType_BLOCK_MSG    PullMsgType = 1
	PullMsgType_IDENTITY_MSG PullMsgType = 2
)

var PullMsgType_name = map[int32]string{
	0: "UNDEFINED",
	1: "BLOCK_MSG",
	2: "IDENTITY_MSG",
}

var PullMsgType_value = map[string]int32{
	"UNDEFINED":    0,
	"BLOCK_MSG":    1,
	"IDENTITY_MSG": 2,
}

type GossipMessage_Tag int32

const (
	GossipMessage_UNDEFINED    GossipMessage_Tag = 0
	GossipMessage_EMPTY        GossipMessage_Tag = 1
	GossipMessage_ORG_ONLY     GossipMessage_Tag = 2
	GossipMessage_CHAN_ONLY    GossipMessage_Tag = 3
	GossipMessage_CHAN_AND_ORG GossipMessage_Tag = 4
	GossipMessage_CHAN_OR_ORG  GossipMessage_Tag = 5
)

var GossipMessage_Tag_name = map[int32]string{
	0: "UNDEFINED",
	1: "EMPTY",
	2: "ORG_ONLY",
	3: "CHAN_ONLY",
	4: "CHAN_AND_ORG",
	5: "CHAN_OR_ORG",
}

var GossipMessage_Tag_value = map[string]int32{
	"UNDEFINED":    0,
	"EMPTY":        1,
	"ORG_ONLY":     2,
	"CHAN_ONLY":    3,
	"CHAN_AND_ORG": 4,
	"CHAN_OR_ORG":  5,
}

// GossipEnvelope contains a marshalled
// GossipMessage and a signature over it.
// It may also contain a SecretEnvelope
// which is a marshalled Secret
type GossipEnvelope struct {
	Payload        []byte
	Signature      []byte
	SecretEnvelope *SecretEnvelope
}

func (m *GossipEnvelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}
func (m *GossipEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}
func (m *GossipEnvelope) GetSecretEnvelope() *SecretEnvelope {
	if m != nil {
		return m.SecretEnvelope
	}
	return nil
}

// SecretEnvelope is a marshalled Secret
// and a signature over it.
// The signature should be validated by the peer
// that signed the GossipEnvelope the SecretEnvelope
// came with
type SecretEnvelope struct {
	Payload   []byte
	Signature []byte
}

func (m *SecretEnvelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}
func (m *SecretEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Secret is an entity that might be omitted
// from an GossipEnvelope when the remote peer that is receiving
// the GossipEnvelope shouldn't know the secret's content.
type Secret struct {
	// Types that are valid to be assigned to Content:
	//	*Secret_InternalEndpoint
	Content isSecret_Content
}

type isSecret_Content interface {
	isSecret_Content()
}

type Secret_InternalEndpoint struct {
	InternalEndpoint string
}

func (*Secret_InternalEndpoint) isSecret_Content() {}
func (m *Secret) GetContent() isSecret_Content {
	if m != nil {
		return m.Content
	}
	return nil
}
func (m *Secret) GetInternalEndpoint() string {
	if x, ok := m.GetContent().(*Secret_InternalEndpoint); ok {
		return x.InternalEndpoint
	}
	return ""
}

// GossipMessage defines the message sent in a gossip network
type GossipMessage struct {
	// used mainly for testing, but will might be used in the future
	// for ensuring message delivery by acking
	Nonce uint64
	// The channel of the message.
	// Some GossipMessages may set this to nil, because
	// they are cross-channels but some may not
	Channel []byte
	// determines to which peers it is allowed
	// to forward the message
	Tag GossipMessage_Tag
	// Types that are valid to be assigned to Content:
	//	*GossipMessage_AliveMsg
	//	*GossipMessage_MemReq
	//	*GossipMessage_MemRes
	//	*GossipMessage_DataMsg
	//	*GossipMessage_Hello
	//	*GossipMessage_DataDig
	//	*GossipMessage_DataReq
	//	*GossipMessage_DataUpdate
	//	*GossipMessage_Empty
	//	*GossipMessage_Conn
	//	*GossipMessage_StateInfo
	//	*GossipMessage_StateSnapshot
	//	*GossipMessage_StateInfoPullReq
	//	*GossipMessage_StateRequest
	//	*GossipMessage_StateResponse
	//	*GossipMessage_LeadershipMsg
	//	*GossipMessage_PeerIdentity
	//	*GossipMessage_Ack
	//	*GossipMessage_PrivateReq
	//	*GossipMessage_PrivateRes
	//	*GossipMessage_PrivateData
	Content isGossipMessage_Content
}

func (m *GossipMessage) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}
func (m *GossipMessage) GetChannel() []byte {
	if m != nil {
		return m.Channel
	}
	return nil
}
func (m *GossipMessage) GetTag() GossipMessage_Tag {
	if m != nil {
		return m.Tag
	}
	return GossipMessage_UNDEFINED
}

type isGossipMessage_Content interface {
	isGossipMessage_Content()
}

type GossipMessage_AliveMsg struct {
	AliveMsg *AliveMessage `protobuf:"bytes,5,opt,name=alive_msg,json=aliveMsg,proto3,oneof"`
}

type GossipMessage_MemReq struct {
	MemReq *MembershipRequest `protobuf:"bytes,6,opt,name=mem_req,json=memReq,proto3,oneof"`
}

type GossipMessage_MemRes struct {
	MemRes *MembershipResponse `protobuf:"bytes,7,opt,name=mem_res,json=memRes,proto3,oneof"`
}

type GossipMessage_DataMsg struct {
	DataMsg *DataMessage `protobuf:"bytes,8,opt,name=data_msg,json=dataMsg,proto3,oneof"`
}

type GossipMessage_Hello struct {
	Hello *GossipHello `protobuf:"bytes,9,opt,name=hello,proto3,oneof"`
}

type GossipMessage_DataDig struct {
	DataDig *DataDigest `protobuf:"bytes,10,opt,name=data_dig,json=dataDig,proto3,oneof"`
}

type GossipMessage_DataReq struct {
	DataReq *DataRequest `protobuf:"bytes,11,opt,name=data_req,json=dataReq,proto3,oneof"`
}

type GossipMessage_DataUpdate struct {
	DataUpdate *DataUpdate `protobuf:"bytes,12,opt,name=data_update,json=dataUpdate,proto3,oneof"`
}

type GossipMessage_Empty struct {
	Empty *Empty `protobuf:"bytes,13,opt,name=empty,proto3,oneof"`
}

type GossipMessage_Conn struct {
	Conn *ConnEstablish `protobuf:"bytes,14,opt,name=conn,proto3,oneof"`
}

type GossipMessage_StateInfo struct {
	StateInfo *StateInfo `protobuf:"bytes,15,opt,name=state_info,json=stateInfo,proto3,oneof"`
}

type GossipMessage_StateSnapshot struct {
	StateSnapshot *StateInfoSnapshot `protobuf:"bytes,16,opt,name=state_snapshot,json=stateSnapshot,proto3,oneof"`
}

type GossipMessage_StateInfoPullReq struct {
	StateInfoPullReq *StateInfoPullRequest `protobuf:"bytes,17,opt,name=state_info_pull_req,json=stateInfoPullReq,proto3,oneof"`
}

type GossipMessage_StateRequest struct {
	StateRequest *RemoteStateRequest `protobuf:"bytes,18,opt,name=state_request,json=stateRequest,proto3,oneof"`
}

type GossipMessage_StateResponse struct {
	StateResponse *RemoteStateResponse `protobuf:"bytes,19,opt,name=state_response,json=stateResponse,proto3,oneof"`
}

type GossipMessage_LeadershipMsg struct {
	LeadershipMsg *LeadershipMessage `protobuf:"bytes,20,opt,name=leadership_msg,json=leadershipMsg,proto3,oneof"`
}

type GossipMessage_PeerIdentity struct {
	PeerIdentity *PeerIdentity `protobuf:"bytes,21,opt,name=peer_identity,json=peerIdentity,proto3,oneof"`
}

type GossipMessage_Ack struct {
	Ack *Acknowledgement `protobuf:"bytes,22,opt,name=ack,proto3,oneof"`
}

type GossipMessage_PrivateReq struct {
	PrivateReq *RemotePvtDataRequest `protobuf:"bytes,23,opt,name=privateReq,proto3,oneof"`
}

type GossipMessage_PrivateRes struct {
	PrivateRes *RemotePvtDataResponse `protobuf:"bytes,24,opt,name=privateRes,proto3,oneof"`
}

type GossipMessage_PrivateData struct {
	PrivateData *PrivateDataMessage `protobuf:"bytes,25,opt,name=private_data,json=privateData,proto3,oneof"`
}

func (*GossipMessage_AliveMsg) isGossipMessage_Content()         {}
func (*GossipMessage_MemReq) isGossipMessage_Content()           {}
func (*GossipMessage_MemRes) isGossipMessage_Content()           {}
func (*GossipMessage_DataMsg) isGossipMessage_Content()          {}
func (*GossipMessage_Hello) isGossipMessage_Content()            {}
func (*GossipMessage_DataDig) isGossipMessage_Content()          {}
func (*GossipMessage_DataReq) isGossipMessage_Content()          {}
func (*GossipMessage_DataUpdate) isGossipMessage_Content()       {}
func (*GossipMessage_Empty) isGossipMessage_Content()            {}
func (*GossipMessage_Conn) isGossipMessage_Content()             {}
func (*GossipMessage_StateInfo) isGossipMessage_Content()        {}
func (*GossipMessage_StateSnapshot) isGossipMessage_Content()    {}
func (*GossipMessage_StateInfoPullReq) isGossipMessage_Content() {}
func (*GossipMessage_StateRequest) isGossipMessage_Content()     {}
func (*GossipMessage_StateResponse) isGossipMessage_Content()    {}
func (*GossipMessage_LeadershipMsg) isGossipMessage_Content()    {}
func (*GossipMessage_PeerIdentity) isGossipMessage_Content()     {}
func (*GossipMessage_Ack) isGossipMessage_Content()              {}
func (*GossipMessage_PrivateReq) isGossipMessage_Content()       {}
func (*GossipMessage_PrivateRes) isGossipMessage_Content()       {}
func (*GossipMessage_PrivateData) isGossipMessage_Content()      {}

func (m *GossipMessage) GetContent() isGossipMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}
func (m *GossipMessage) GetAliveMsg() *AliveMessage {
	if x, ok := m.GetContent().(*GossipMessage_AliveMsg); ok {
		return x.AliveMsg
	}
	return nil
}
func (m *GossipMessage) GetMemReq() *MembershipRequest {
	if x, ok := m.GetContent().(*GossipMessage_MemReq); ok {
		return x.MemReq
	}
	return nil
}
func (m *GossipMessage) GetMemRes() *MembershipResponse {
	if x, ok := m.GetContent().(*GossipMessage_MemRes); ok {
		return x.MemRes
	}
	return nil
}
func (m *GossipMessage) GetDataMsg() *DataMessage {
	if x, ok := m.GetContent().(*GossipMessage_DataMsg); ok {
		return x.DataMsg
	}
	return nil
}
func (m *GossipMessage) GetHello() *GossipHello {
	if x, ok := m.GetContent().(*GossipMessage_Hello); ok {
		return x.Hello
	}
	return nil
}
func (m *GossipMessage) GetDataDig() *DataDigest {
	if x, ok := m.GetContent().(*GossipMessage_DataDig); ok {
		return x.DataDig
	}
	return nil
}
func (m *GossipMessage) GetDataReq() *DataRequest {
	if x, ok := m.GetContent().(*GossipMessage_DataReq); ok {
		return x.DataReq
	}
	return nil
}
func (m *GossipMessage) GetDataUpdate() *DataUpdate {
	if x, ok := m.GetContent().(*GossipMessage_DataUpdate); ok {
		return x.DataUpdate
	}
	return nil
}
func (m *GossipMessage) GetEmpty() *Empty {
	if x, ok := m.GetContent().(*GossipMessage_Empty); ok {
		return x.Empty
	}
	return nil
}
func (m *GossipMessage) GetConn() *ConnEstablish {
	if x, ok := m.GetContent().(*GossipMessage_Conn); ok {
		return x.Conn
	}
	return nil
}
func (m *GossipMessage) GetStateInfo() *StateInfo {
	if x, ok := m.GetContent().(*GossipMessage_StateInfo); ok {
		return x.StateInfo
	}
	return nil
}
func (m *GossipMessage) GetStateSnapshot() *StateInfoSnapshot {
	if x, ok := m.GetContent().(*GossipMessage_StateSnapshot); ok {
		return x.StateSnapshot
	}
	return nil
}
func (m *GossipMessage) GetStateInfoPullReq() *StateInfoPullRequest {
	if x, ok := m.GetContent().(*GossipMessage_StateInfoPullReq); ok {
		return x.StateInfoPullReq
	}
	return nil
}
func (m *GossipMessage) GetStateRequest() *RemoteStateRequest {
	if x, ok := m.GetContent().(*GossipMessage_StateRequest); ok {
		return x.StateRequest
	}
	return nil
}
func (m *GossipMessage) GetStateResponse() *RemoteStateResponse {
	if x, ok := m.GetContent().(*GossipMessage_StateResponse); ok {
		return x.StateResponse
	}
	return nil
}
func (m *GossipMessage) GetLeadershipMsg() *LeadershipMessage {
	if x, ok := m.GetContent().(*GossipMessage_LeadershipMsg); ok {
		return x.LeadershipMsg
	}
	return nil
}
func (m *GossipMessage) GetPeerIdentity() *PeerIdentity {
	if x, ok := m.GetContent().(*GossipMessage_PeerIdentity); ok {
		return x.PeerIdentity
	}
	return nil
}
func (m *GossipMessage) GetAck() *Acknowledgement {
	if x, ok := m.GetContent().(*GossipMessage_Ack); ok {
		return x.Ack
	}
	return nil
}
func (m *GossipMessage) GetPrivateReq() *RemotePvtDataRequest {
	if x, ok := m.GetContent().(*GossipMessage_PrivateReq); ok {
		return x.PrivateReq
	}
	return nil
}
func (m *GossipMessage) GetPrivateRes() *RemotePvtDataResponse {
	if x, ok := m.GetContent().(*GossipMessage_PrivateRes); ok {
		return x.PrivateRes
	}
	return nil
}
func (m *GossipMessage) GetPrivateData() *PrivateDataMessage {
	if x, ok := m.GetContent().(*GossipMessage_PrivateData); ok {
		return x.PrivateData
	}
	return nil
}

// StateInfo is used for a peer to relay its state information
// to other peers
type StateInfo struct {
	Timestamp *PeerTime
	PkiId     []byte
	// channel_MAC is an authentication code that proves
	// that the peer that sent this message knows
	// the name of the channel.
	Channel_MAC []byte
	Properties  *Properties
}

func (m *StateInfo) GetTimestamp() *PeerTime {
	if m != nil {
		return m.Timestamp
	}
	return nil
}
func (m *StateInfo) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}
func (m *StateInfo) GetChannel_MAC() []byte {
	if m != nil {
		return m.Channel_MAC
	}
	return nil
}
func (m *StateInfo) GetProperties() *Properties {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Properties struct {
	LedgerHeight uint64
	LeftChannel  bool
	Chaincodes   []*Chaincode
}

func (m *Properties) GetLedgerHeight() uint64 {
	if m != nil {
		return m.LedgerHeight
	}
	return 0
}
func (m *Properties) GetLeftChannel() bool {
	if m != nil {
		return m.LeftChannel
	}
	return false
}
func (m *Properties) GetChaincodes() []*Chaincode {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// StateInfoSnapshot is an aggregation of StateInfo messages
type StateInfoSnapshot struct {
	Elements []*GossipEnvelope
}

func (m *StateInfoSnapshot) GetElements() []*GossipEnvelope {
	if m != nil {
		return m.Elements
	}
	return nil
}

// StateInfoPullRequest is used to fetch a StateInfoSnapshot
// from a remote peer
type StateInfoPullRequest struct {
	// channel_MAC is an authentication code that proves
	// that the peer that sent this message knows
	// the name of the channel.
	Channel_MAC []byte `protobuf:"bytes,1,opt,name=channel_MAC,json=channelMAC,proto3" json:"channel_MAC,omitempty"`
}

func (m *StateInfoPullRequest) GetChannel_MAC() []byte {
	if m != nil {
		return m.Channel_MAC
	}
	return nil
}

// ConnEstablish is the message used for the gossip handshake
// Whenever a peer connects to another peer, it handshakes
// with it by sending this message that proves its identity
type ConnEstablish struct {
	PkiId       []byte
	Identity    []byte
	TlsCertHash []byte
	Probe       bool
}

func (m *ConnEstablish) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}

func (m *ConnEstablish) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *ConnEstablish) GetTlsCertHash() []byte {
	if m != nil {
		return m.TlsCertHash
	}
	return nil
}

func (m *ConnEstablish) GetProbe() bool {
	if m != nil {
		return m.Probe
	}
	return false
}

// PeerIdentity defines the identity of the peer
// Used to make other peers learn of the identity
// of a certain peer
type PeerIdentity struct {
	PkiId    []byte
	Cert     []byte
	Metadata []byte
}

func (m *PeerIdentity) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}

func (m *PeerIdentity) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *PeerIdentity) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// DataRequest is a message used for a peer to request
// certain data blocks from a remote peer
type DataRequest struct {
	Nonce   uint64
	Digests [][]byte
	MsgType PullMsgType
}

func (m *DataRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *DataRequest) GetDigests() [][]byte {
	if m != nil {
		return m.Digests
	}
	return nil
}

func (m *DataRequest) GetMsgType() PullMsgType {
	if m != nil {
		return m.MsgType
	}
	return PullMsgType_UNDEFINED
}

// GossipHello is the message that is used for the peer to initiate
// a pull round with another peer
type GossipHello struct {
	Nonce    uint64
	Metadata []byte
	MsgType  PullMsgType
}

func (m *GossipHello) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *GossipHello) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GossipHello) GetMsgType() PullMsgType {
	if m != nil {
		return m.MsgType
	}
	return PullMsgType_UNDEFINED
}

// DataUpdate is the final message in the pull phase
// sent from the receiver to the initiator
type DataUpdate struct {
	Nonce   uint64
	Data    []*GossipEnvelope
	MsgType PullMsgType
}

func (m *DataUpdate) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *DataUpdate) GetData() []*GossipEnvelope {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DataUpdate) GetMsgType() PullMsgType {
	if m != nil {
		return m.MsgType
	}
	return PullMsgType_UNDEFINED
}

// DataDigest is the message sent from the receiver peer
// to the initator peer and contains the data items it has
type DataDigest struct {
	Nonce   uint64
	Digests [][]byte
	MsgType PullMsgType
}

func (m *DataDigest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *DataDigest) GetDigests() [][]byte {
	if m != nil {
		return m.Digests
	}
	return nil
}

func (m *DataDigest) GetMsgType() PullMsgType {
	if m != nil {
		return m.MsgType
	}
	return PullMsgType_UNDEFINED
}

// DataMessage is the message that contains a block
type DataMessage struct {
	Payload *Payload
}

func (m *DataMessage) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// PrivateDataMessage message which includes private
// data information to distributed once transaction
// has been endorsed
type PrivateDataMessage struct {
	Payload *PrivatePayload
}

func (m *PrivateDataMessage) GetPayload() *PrivatePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Payload contains a block
type Payload struct {
	SeqNum      uint64
	Data        []byte
	PrivateData [][]byte
}

func (m *Payload) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Payload) GetPrivateData() [][]byte {
	if m != nil {
		return m.PrivateData
	}
	return nil
}

// PrivatePayload payload to encapsulate private
// data with collection name to enable routing
// based on collection partitioning
type PrivatePayload struct {
	CollectionName    string                        `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Namespace         string                        `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TxId              string                        `protobuf:"bytes,3,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	PrivateRwset      []byte                        `protobuf:"bytes,4,opt,name=private_rwset,json=privateRwset,proto3" json:"private_rwset,omitempty"`
	PrivateSimHeight  uint64                        `protobuf:"varint,5,opt,name=private_sim_height,json=privateSimHeight,proto3" json:"private_sim_height,omitempty"`
	CollectionConfigs *peer.CollectionConfigPackage `protobuf:"bytes,6,opt,name=collection_configs,json=collectionConfigs,proto3" json:"collection_configs,omitempty"`
}

func (m *PrivatePayload) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}
func (m *PrivatePayload) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}
func (m *PrivatePayload) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}
func (m *PrivatePayload) GetPrivateRwset() []byte {
	if m != nil {
		return m.PrivateRwset
	}
	return nil
}
func (m *PrivatePayload) GetPrivateSimHeight() uint64 {
	if m != nil {
		return m.PrivateSimHeight
	}
	return 0
}
func (m *PrivatePayload) GetCollectionConfigs() *peer.CollectionConfigPackage {
	if m != nil {
		return m.CollectionConfigs
	}
	return nil
}

// AliveMessage is sent to inform remote peers
// of a peer's existence and activity
type AliveMessage struct {
	Membership *Member
	Timestamp  *PeerTime
	Identity   []byte
}

func (m *AliveMessage) GetMembership() *Member {
	if m != nil {
		return m.Membership
	}
	return nil
}
func (m *AliveMessage) GetTimestamp() *PeerTime {
	if m != nil {
		return m.Timestamp
	}
	return nil
}
func (m *AliveMessage) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

// Leadership Message is sent during leader election to inform
// remote peers about intent of peer to proclaim itself as leader
type LeadershipMessage struct {
	PkiId         []byte
	Timestamp     *PeerTime
	IsDeclaration bool
}

func (m *LeadershipMessage) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}
func (m *LeadershipMessage) GetTimestamp() *PeerTime {
	if m != nil {
		return m.Timestamp
	}
	return nil
}
func (m *LeadershipMessage) GetIsDeclaration() bool {
	if m != nil {
		return m.IsDeclaration
	}
	return false
}

// PeerTime defines the logical time of a peer's life
type PeerTime struct {
	IncNum uint64
	SeqNum uint64
}

func (m *PeerTime) GetIncNum() uint64 {
	if m != nil {
		return m.IncNum
	}
	return 0
}
func (m *PeerTime) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

// MembershipRequest is used to ask membership information
// from a remote peer
type MembershipRequest struct {
	SelfInformation *GossipEnvelope
	Known           [][]byte
}

func (m *MembershipRequest) GetSelfInformation() *GossipEnvelope {
	if m != nil {
		return m.SelfInformation
	}
	return nil
}
func (m *MembershipRequest) GetKnown() [][]byte {
	if m != nil {
		return m.Known
	}
	return nil
}

// MembershipResponse is used for replying to MembershipRequests
type MembershipResponse struct {
	Alive []*GossipEnvelope
	Dead  []*GossipEnvelope
}

func (m *MembershipResponse) GetAlive() []*GossipEnvelope {
	if m != nil {
		return m.Alive
	}
	return nil
}
func (m *MembershipResponse) GetDead() []*GossipEnvelope {
	if m != nil {
		return m.Dead
	}
	return nil
}

// Member holds membership-related information
// about a peer
type Member struct {
	Endpoint string
	Metadata []byte
	PkiId    []byte
}

func (m *Member) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}
func (m *Member) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}
func (m *Member) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}

// Empty is used for pinging and in tests
type Empty struct{}

// RemoteStateRequest is used to ask a set of blocks
// from a remote peer
type RemoteStateRequest struct {
	StartSeqNum uint64
	EndSeqNum   uint64
}

func (m *RemoteStateRequest) GetStartSeqNum() uint64 {
	if m != nil {
		return m.StartSeqNum
	}
	return 0
}
func (m *RemoteStateRequest) GetEndSeqNum() uint64 {
	if m != nil {
		return m.EndSeqNum
	}
	return 0
}

// RemoteStateResponse is used to send a set of blocks
// to a remote peer
type RemoteStateResponse struct {
	Payloads []*Payload
}

func (m *RemoteStateResponse) GetPayloads() []*Payload {
	if m != nil {
		return m.Payloads
	}
	return nil
}

// RemotePrivateDataRequest message used to request
// missing private rwset
type RemotePvtDataRequest struct {
	Digests []*PvtDataDigest
}

func (m *RemotePvtDataRequest) GetDigests() []*PvtDataDigest {
	if m != nil {
		return m.Digests
	}
	return nil
}

// PvtDataDigest defines a digest of private data
type PvtDataDigest struct {
	TxId       string
	Namespace  string
	Collection string
	BlockSeq   uint64
	SeqInBlock uint64
}

func (m *PvtDataDigest) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}
func (m *PvtDataDigest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}
func (m *PvtDataDigest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}
func (m *PvtDataDigest) GetBlockSeq() uint64 {
	if m != nil {
		return m.BlockSeq
	}
	return 0
}
func (m *PvtDataDigest) GetSeqInBlock() uint64 {
	if m != nil {
		return m.SeqInBlock
	}
	return 0
}

// RemotePrivateData message to response on private
// data replication request
type RemotePvtDataResponse struct {
	Elements []*PvtDataElement
}

func (m *RemotePvtDataResponse) GetElements() []*PvtDataElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

type PvtDataElement struct {
	Digest *PvtDataDigest
	// the payload is a marshaled kvrwset.KVRWSet
	Payload [][]byte
}

func (m *PvtDataElement) GetDigest() *PvtDataDigest {
	if m != nil {
		return m.Digest
	}
	return nil
}
func (m *PvtDataElement) GetPayload() [][]byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// PvtPayload augments private rwset data and tx index
// inside the block
type PvtDataPayload struct {
	TxSeqInBlock uint64
	// Encodes marhslaed bytes of rwset.TxPvtReadWriteSet
	// defined in rwset.proto
	Payload []byte
}

func (m *PvtDataPayload) GetTxSeqInBlock() uint64 {
	if m != nil {
		return m.TxSeqInBlock
	}
	return 0
}
func (m *PvtDataPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Acknowledgement struct {
	Error string
}

func (m *Acknowledgement) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// Chaincode represents a Chaincode that is installed
// on a peer
type Chaincode struct {
	Name     string
	Version  string
	Metadata []byte
}

func (m *Chaincode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *Chaincode) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}
func (m *Chaincode) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}
