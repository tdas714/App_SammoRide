package lifecycle

// ChaincodeEndorsementInfo is (most) everything the peer needs to know in order
// to execute a chaincode
type ChaincodeEndorsementInfo struct {
	Version           string
	InitRequired      bool
	EndorsementPlugin string
}

func (m *ChaincodeEndorsementInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeEndorsementInfo) GetInitRequired() bool {
	if m != nil {
		return m.InitRequired
	}
	return false
}

func (m *ChaincodeEndorsementInfo) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}

// ValidationInfo is (most) everything the peer needs to know in order
// to validate a transaction
type ChaincodeValidationInfo struct {
	ValidationPlugin    string
	ValidationParameter []byte
}

func (m *ChaincodeValidationInfo) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}

func (m *ChaincodeValidationInfo) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}
