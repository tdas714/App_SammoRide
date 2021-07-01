package peer

// SignedChaincodeDeploymentSpec carries the CDS along with endorsements
type SignedChaincodeDeploymentSpec struct {
	// This is the bytes of the ChaincodeDeploymentSpec
	ChaincodeDeploymentSpec []byte
	// This is the instantiation policy which is identical in structure
	// to endorsement policy.  This policy is checked by the VSCC at commit
	// time on the instantiation (all peers will get the same policy as it
	// will be part of the LSCC instantation record and will be part of the
	// hash as well)
	InstantiationPolicy []byte
	// The endorsements of the above deployment spec, the owner's signature over
	// chaincode_deployment_spec and Endorsement.endorser.
	OwnerEndorsements []*Endorsement
}

func (m *SignedChaincodeDeploymentSpec) GetChaincodeDeploymentSpec() []byte {
	if m != nil {
		return m.ChaincodeDeploymentSpec
	}
	return nil
}

func (m *SignedChaincodeDeploymentSpec) GetInstantiationPolicy() []byte {
	if m != nil {
		return m.InstantiationPolicy
	}
	return nil
}

func (m *SignedChaincodeDeploymentSpec) GetOwnerEndorsements() []*Endorsement {
	if m != nil {
		return m.OwnerEndorsements
	}
	return nil
}
