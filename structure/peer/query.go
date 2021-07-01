package peer

// ChaincodeQueryResponse returns information about each chaincode that pertains
// to a query in lscc.go, such as GetChaincodes (returns all chaincodes
// instantiated on a channel), and GetInstalledChaincodes (returns all chaincodes
// installed on a peer)
type ChaincodeQueryResponse struct {
	Chaincodes []*ChaincodeInfo
}

func (m *ChaincodeQueryResponse) GetChaincodes() []*ChaincodeInfo {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeInfo contains general information about an installed/instantiated
// chaincode
type ChaincodeInfo struct {
	Name    string
	Version string
	// the path as specified by the install/instantiate transaction
	Path string
	// the chaincode function upon instantiation and its arguments. This will be
	// blank if the query is returning information about installed chaincodes.
	Input string
	// the name of the ESCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Escc string
	// the name of the VSCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Vscc string
	// the chaincode unique id.
	// computed as: H(
	//                H(name || version) ||
	//                H(CodePackage)
	//              )
	Id []byte
}

func (m *ChaincodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeInfo) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ChaincodeInfo) GetEscc() string {
	if m != nil {
		return m.Escc
	}
	return ""
}

func (m *ChaincodeInfo) GetVscc() string {
	if m != nil {
		return m.Vscc
	}
	return ""
}

func (m *ChaincodeInfo) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// ChannelQueryResponse returns information about each channel that pertains
// to a query in lscc.go, such as GetChannels (returns all channels for a
// given peer)
type ChannelQueryResponse struct {
	Channels []*ChannelInfo
}

func (m *ChannelQueryResponse) GetChannels() []*ChannelInfo {
	if m != nil {
		return m.Channels
	}
	return nil
}

// ChannelInfo contains general information about channels
type ChannelInfo struct {
	ChannelId string
}

func (m *ChannelInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// JoinBySnapshotStatus contains information about whether or a JoinBySnapshot operation
// is in progress and the related bootstrap dir if it is running.
type JoinBySnapshotStatus struct {
	InProgress               bool
	BootstrappingSnapshotDir string
}

func (m *JoinBySnapshotStatus) GetInProgress() bool {
	if m != nil {
		return m.InProgress
	}
	return false
}

func (m *JoinBySnapshotStatus) GetBootstrappingSnapshotDir() string {
	if m != nil {
		return m.BootstrappingSnapshotDir
	}
	return ""
}
