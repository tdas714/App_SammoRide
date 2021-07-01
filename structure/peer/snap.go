package peer

import "github.com/App-SammoRide/structure"

// SnapshotRequest contains information for a generate/cancel snapshot request
type SnapshotRequest struct {
	// The signature header that contains creator identity and nonce
	SignatureHeader *structure.SignatureHeader
	// The channel ID
	ChannelId string
	// The block number to generate a snapshot
	BlockNumber uint64
}

func (m *SnapshotRequest) GetSignatureHeader() *structure.SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *SnapshotRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SnapshotRequest) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

// SnapshotQuery contains information for a query snapshot request
type SnapshotQuery struct {
	// The signature header that contains creator identity and nonce
	SignatureHeader *structure.SignatureHeader
	// The channel ID
	ChannelId string
}

func (m *SnapshotQuery) GetSignatureHeader() *structure.SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *SnapshotQuery) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// SignedSnapshotRequest contains marshalled request bytes and signature
type SignedSnapshotRequest struct {
	// The bytes of SnapshotRequest or SnapshotQuery
	Request []byte
	// Signaure over request bytes; this signature is to be verified against the client identity
	Signature []byte
}

func (m *SignedSnapshotRequest) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SignedSnapshotRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// QueryPendingSnapshotsResponse specifies the response payload of a query pending snapshots request
type QueryPendingSnapshotsResponse struct {
	BlockNumbers []uint64
}

func (m *QueryPendingSnapshotsResponse) GetBlockNumbers() []uint64 {
	if m != nil {
		return m.BlockNumbers
	}
	return nil
}
