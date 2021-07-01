package discovery

// SignedRequest contains a serialized Request in the payload field
// and a signature.
// The identity that is used to verify the signature
// can be extracted from the authentication field of type AuthInfo
// in the Request itself after deserializing it.
type SignedRequest struct {
	Payload   []byte
	Signature []byte
}

func (m *SignedRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SignedRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Request contains authentication info about the client that sent the request
// and the queries it wishes to query the service
type Request struct {
	// authentication contains information that the service uses to check
	// the client's eligibility for the queries.
	Authentication *AuthInfo
	// queries
	Queries []*Query
}

func (m *Request) GetAuthentication() *AuthInfo {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Request) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Response struct {
	// The results are returned in the same order of the queries
	Results []*QueryResult
}

func (m *Response) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// AuthInfo aggregates authentication information that the server uses
// to authenticate the client
type AuthInfo struct {
	// This is the identity of the client that is used to verify the signature
	// on the SignedRequest's payload.
	// It is a msp.SerializedIdentity in bytes form
	ClientIdentity []byte
	// This is the hash of the client's TLS cert.
	// When the network is running with TLS, clients that don't include a certificate
	// will be denied access to the service.
	// Since the Request is encapsulated with a SignedRequest (which is signed),
	// this binds the TLS session to the enrollement identity of the client and
	// therefore both authenticates the client to the server,
	// and also prevents the server from relaying the request message to another server.
	ClientTlsCertHash []byte
}

func (m *AuthInfo) GetClientIdentity() []byte {
	if m != nil {
		return m.ClientIdentity
	}
	return nil
}

func (m *AuthInfo) GetClientTlsCertHash() []byte {
	if m != nil {
		return m.ClientTlsCertHash
	}
	return nil
}

// Query asks for information in the context of a specific channel
type Query struct {
	Channel string
	// Types that are valid to be assigned to Query:
	//	*Query_ConfigQuery
	//	*Query_PeerQuery
	//	*Query_CcQuery
	//	*Query_LocalPeers
	Query isQuery_Query
}

func (m *Query) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type isQuery_Query interface {
	isQuery_Query()
}

type Query_ConfigQuery struct {
	ConfigQuery *ConfigQuery
}

type Query_PeerQuery struct {
	PeerQuery *PeerMembershipQuery
}

type Query_CcQuery struct {
	CcQuery *ChaincodeQuery
}

type Query_LocalPeers struct {
	LocalPeers *LocalPeerQuery
}

func (*Query_ConfigQuery) isQuery_Query() {}

func (*Query_PeerQuery) isQuery_Query() {}

func (*Query_CcQuery) isQuery_Query() {}

func (*Query_LocalPeers) isQuery_Query() {}

func (m *Query) GetQuery() isQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Query) GetConfigQuery() *ConfigQuery {
	if x, ok := m.GetQuery().(*Query_ConfigQuery); ok {
		return x.ConfigQuery
	}
	return nil
}

func (m *Query) GetPeerQuery() *PeerMembershipQuery {
	if x, ok := m.GetQuery().(*Query_PeerQuery); ok {
		return x.PeerQuery
	}
	return nil
}

func (m *Query) GetCcQuery() *ChaincodeQuery {
	if x, ok := m.GetQuery().(*Query_CcQuery); ok {
		return x.CcQuery
	}
	return nil
}

func (m *Query) GetLocalPeers() *LocalPeerQuery {
	if x, ok := m.GetQuery().(*Query_LocalPeers); ok {
		return x.LocalPeers
	}
	return nil
}

// QueryResult contains a result for a given Query.
// The corresponding Query can be inferred by the index of the QueryResult from
// its enclosing Response message.
// QueryResults are ordered in the same order as the Queries are ordered in their enclosing Request.
type QueryResult struct {
	// Types that are valid to be assigned to Result:
	//	*QueryResult_Error
	//	*QueryResult_ConfigResult
	//	*QueryResult_CcQueryRes
	//	*QueryResult_Members
	Result isQueryResult_Result
}

type isQueryResult_Result interface {
	isQueryResult_Result()
}

type QueryResult_Error struct {
	Error *Error
}

type QueryResult_ConfigResult struct {
	ConfigResult *ConfigResult
}

type QueryResult_CcQueryRes struct {
	CcQueryRes *ChaincodeQueryResult
}

type QueryResult_Members struct {
	Members *PeerMembershipResult
}

func (*QueryResult_Error) isQueryResult_Result() {}

func (*QueryResult_ConfigResult) isQueryResult_Result() {}

func (*QueryResult_CcQueryRes) isQueryResult_Result() {}

func (*QueryResult_Members) isQueryResult_Result() {}

func (m *QueryResult) GetResult() isQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *QueryResult) GetError() *Error {
	if x, ok := m.GetResult().(*QueryResult_Error); ok {
		return x.Error
	}
	return nil
}

func (m *QueryResult) GetConfigResult() *ConfigResult {
	if x, ok := m.GetResult().(*QueryResult_ConfigResult); ok {
		return x.ConfigResult
	}
	return nil
}

func (m *QueryResult) GetCcQueryRes() *ChaincodeQueryResult {
	if x, ok := m.GetResult().(*QueryResult_CcQueryRes); ok {
		return x.CcQueryRes
	}
	return nil
}

func (m *QueryResult) GetMembers() *PeerMembershipResult {
	if x, ok := m.GetResult().(*QueryResult_Members); ok {
		return x.Members
	}
	return nil
}

// ConfigQuery requests a ConfigResult
type ConfigQuery struct{}

type ConfigResult struct {
	// msps is a map from MSP_ID to FabricMSPConfig
	Msps map[string]*msp.FabricMSPConfig
	// orderers is a map from MSP_ID to endpoint lists of orderers
	Orderers map[string]*Endpoints
}

func (m *ConfigResult) GetMsps() map[string]*msp.FabricMSPConfig {
	if m != nil {
		return m.Msps
	}
	return nil
}

func (m *ConfigResult) GetOrderers() map[string]*Endpoints {
	if m != nil {
		return m.Orderers
	}
	return nil
}

// PeerMembershipQuery requests PeerMembershipResult.
// The filter field may be optionally populated in order
// for the peer membership to be filtered according to
// chaincodes that are installed on peers and collection
// access control policies.
type PeerMembershipQuery struct {
	Filter *ChaincodeInterest
}

func (m *PeerMembershipQuery) GetFilter() *ChaincodeInterest {
	if m != nil {
		return m.Filter
	}
	return nil
}

// PeerMembershipResult contains peers mapped by their organizations (MSP_ID)
type PeerMembershipResult struct {
	PeersByOrg map[string]*Peers
}

func (m *PeerMembershipResult) GetPeersByOrg() map[string]*Peers {
	if m != nil {
		return m.PeersByOrg
	}
	return nil
}

// ChaincodeQuery requests ChaincodeQueryResults for a given
// list of chaincode invocations.
// Each invocation is a separate one, and the endorsement policy
// is evaluated independantly for each given interest.
type ChaincodeQuery struct {
	Interests []*ChaincodeInterest
}

func (m *ChaincodeQuery) GetInterests() []*ChaincodeInterest {
	if m != nil {
		return m.Interests
	}
	return nil
}

// ChaincodeInterest defines an interest about an endorsement
// for a specific single chaincode invocation.
// Multiple chaincodes indicate chaincode to chaincode invocations.
type ChaincodeInterest struct {
	Chaincodes []*ChaincodeCall
}

func (m *ChaincodeInterest) GetChaincodes() []*ChaincodeCall {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeCall defines a call to a chaincode.
// It may have collections that are related to the chaincode
type ChaincodeCall struct {
	Name            string
	CollectionNames []string
	NoPrivateReads  bool
	NoPublicWrites  bool
}

func (m *ChaincodeCall) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeCall) GetCollectionNames() []string {
	if m != nil {
		return m.CollectionNames
	}
	return nil
}

func (m *ChaincodeCall) GetNoPrivateReads() bool {
	if m != nil {
		return m.NoPrivateReads
	}
	return false
}

func (m *ChaincodeCall) GetNoPublicWrites() bool {
	if m != nil {
		return m.NoPublicWrites
	}
	return false
}

// ChaincodeQueryResult contains EndorsementDescriptors for
// chaincodes
type ChaincodeQueryResult struct {
	Content []*EndorsementDescriptor
}

func (m *ChaincodeQueryResult) GetContent() []*EndorsementDescriptor {
	if m != nil {
		return m.Content
	}
	return nil
}

// LocalPeerQuery queries for peers in a non channel context
type LocalPeerQuery struct {
}

// EndorsementDescriptor contains information about which peers can be used
// to request endorsement from, such that the endorsement policy would be fulfilled.
// Here is how to compute a set of peers to ask an endorsement from, given an EndorsementDescriptor:
// Let e: G --> P be the endorsers_by_groups field that maps a group to a set of peers.
// Note that applying e on a group g yields a set of peers.
// 1) Select a layout l: G --> N out of the layouts given.
//    l is the quantities_by_group field of a Layout, and it maps a group to an integer.
// 2) R = {}  (an empty set of peers)
// 3) For each group g in the layout l, compute n = l(g)
//    3.1) Denote P_g as a set of n random peers {p0, p1, ... p_n} selected from e(g)
//    3.2) R = R U P_g  (add P_g to R)
// 4) The set of peers R is the peers the client needs to request endorsements from
type EndorsementDescriptor struct {
	Chaincode string
	// Specifies the endorsers, separated to groups.
	EndorsersByGroups map[string]*Peers
	// Specifies options of fulfulling the endorsement policy.
	// Each option lists the group names, and the amount of signatures needed
	// from each group.
	Layouts []*Layout
}

func (m *EndorsementDescriptor) GetChaincode() string {
	if m != nil {
		return m.Chaincode
	}
	return ""
}

func (m *EndorsementDescriptor) GetEndorsersByGroups() map[string]*Peers {
	if m != nil {
		return m.EndorsersByGroups
	}
	return nil
}

func (m *EndorsementDescriptor) GetLayouts() []*Layout {
	if m != nil {
		return m.Layouts
	}
	return nil
}

// Layout contains a mapping from a group name to number of peers
// that are needed for fulfilling an endorsement policy
type Layout struct {
	// Specifies how many non repeated signatures of each group
	// are needed for endorsement
	QuantitiesByGroup map[string]uint32
}

func (m *Layout) GetQuantitiesByGroup() map[string]uint32 {
	if m != nil {
		return m.QuantitiesByGroup
	}
	return nil
}

// Peers contains a list of Peer(s)
type Peers struct {
	Peers []*Peer
}

func (m *Peers) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Peer contains information about the peer such as its channel specific
// state, and membership information.
type Peer struct {
	// This is an Envelope of a GossipMessage with a gossip.StateInfo message
	StateInfo *gossip.Envelope
	// This is an Envelope of a GossipMessage with a gossip.AliveMessage message
	MembershipInfo *gossip.Envelope
	// This is the msp.SerializedIdentity of the peer, represented in bytes.
	Identity []byte
}

func (m *Peer) GetStateInfo() *gossip.Envelope {
	if m != nil {
		return m.StateInfo
	}
	return nil
}

func (m *Peer) GetMembershipInfo() *gossip.Envelope {
	if m != nil {
		return m.MembershipInfo
	}
	return nil
}

func (m *Peer) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

// Error denotes that something went wrong and contains the error message
type Error struct {
	Content string
}

func (m *Error) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Endpoints is a list of Endpoint(s)
type Endpoints struct {
	Endpoint []*Endpoint
}

func (m *Endpoints) GetEndpoint() []*Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

// Endpoint is a combination of a host and a port
type Endpoint struct {
	Host string
	Port uint32
}

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}
