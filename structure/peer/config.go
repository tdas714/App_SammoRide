package peer

// AnchorPeers simply represents list of anchor peers which is used in ConfigurationItem
type AnchorPeers struct {
	AnchorPeers []*AnchorPeer
}

func (m *AnchorPeers) GetAnchorPeers() []*AnchorPeer {
	if m != nil {
		return m.AnchorPeers
	}
	return nil
}

// AnchorPeer message structure which provides information about anchor peer, it includes host name,
// port number and peer certificate.
type AnchorPeer struct {
	Host string
	Port int32
}

func (m *AnchorPeer) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AnchorPeer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// APIResource represents an API resource in the peer whose ACL
// is determined by the policy_ref field
type APIResource struct {
	PolicyRef string
}

func (m *APIResource) GetPolicyRef() string {
	if m != nil {
		return m.PolicyRef
	}
	return ""
}

// ACLs provides mappings for resources in a channel. APIResource encapsulates
// reference to a policy used to determine ACL for the resource
type ACLs struct {
	Acls map[string]*APIResource
}

func (m *ACLs) GetAcls() map[string]*APIResource {
	if m != nil {
		return m.Acls
	}
	return nil
}
