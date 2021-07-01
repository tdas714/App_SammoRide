package common

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	Height            uint64
	CurrentBlockHash  []byte
	PreviousBlockHash []byte
	// Specifies bootstrapping snapshot info if the channel is bootstrapped from a snapshot.
	// It is nil if the channel is not bootstrapped from a snapshot.
	BootstrappingSnapshotInfo *BootstrappingSnapshotInfo
}

func (m *BlockchainInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockchainInfo) GetCurrentBlockHash() []byte {
	if m != nil {
		return m.CurrentBlockHash
	}
	return nil
}

func (m *BlockchainInfo) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func (m *BlockchainInfo) GetBootstrappingSnapshotInfo() *BootstrappingSnapshotInfo {
	if m != nil {
		return m.BootstrappingSnapshotInfo
	}
	return nil
}

// Contains information for the bootstrapping snapshot.
type BootstrappingSnapshotInfo struct {
	LastBlockInSnapshot uint64
}

func (m *BootstrappingSnapshotInfo) GetLastBlockInSnapshot() uint64 {
	if m != nil {
		return m.LastBlockInSnapshot
	}
	return 0
}
