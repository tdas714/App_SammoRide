package peer

import "github.com/App-SammoRide/structure"

// ApplicationPolicy captures the diffenrent policy types that
// are set and evaluted at the application level.
type ApplicationPolicy struct {
	// Types that are valid to be assigned to Type:
	//	*ApplicationPolicy_SignaturePolicy
	//	*ApplicationPolicy_ChannelConfigPolicyReference
	Type isApplicationPolicy_Type
}

type isApplicationPolicy_Type interface {
	isApplicationPolicy_Type()
}

type ApplicationPolicy_SignaturePolicy struct {
	SignaturePolicy *structure.SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

type ApplicationPolicy_ChannelConfigPolicyReference struct {
	ChannelConfigPolicyReference string `protobuf:"bytes,2,opt,name=channel_config_policy_reference,json=channelConfigPolicyReference,proto3,oneof"`
}

func (*ApplicationPolicy_SignaturePolicy) isApplicationPolicy_Type() {}

func (*ApplicationPolicy_ChannelConfigPolicyReference) isApplicationPolicy_Type() {}

func (m *ApplicationPolicy) GetType() isApplicationPolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ApplicationPolicy) GetSignaturePolicy() *structure.SignaturePolicyEnvelope {
	if x, ok := m.GetType().(*ApplicationPolicy_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

func (m *ApplicationPolicy) GetChannelConfigPolicyReference() string {
	if x, ok := m.GetType().(*ApplicationPolicy_ChannelConfigPolicyReference); ok {
		return x.ChannelConfigPolicyReference
	}
	return ""
}
