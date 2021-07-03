package common

import (
	"bytes"
	"encoding/json"
	"log"
	"time"
)

type ChannelID string

const (
	Ride ChannelID = "SammoRide"
)

type Header struct {
	ChannelHeader   *ChannelHeader
	SignatureHeader *SignatureHeader
}

func (m *Header) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Common/Serialize")
	}
	return js
}

func DeSerializeHeader(data []byte) *Header {
	var m *Header
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}

func (m *Header) GetChannelHeader() *ChannelHeader {
	if m != nil {
		return m.ChannelHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp time.Time

	ChannelId string
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64
}

func (m *ChannelHeader) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Common/Serialize")
	}
	return js
}

func DeSerializeChannelHeader(data []byte) *ChannelHeader {
	var m *ChannelHeader
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}

func (m *ChannelHeader) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Now().AddDate(25, 0, 0)
}

func (m *ChannelHeader) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChannelHeader) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChannelHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

type SignatureHeader struct {
	// Creator of the message, a marshaled msp.SerializedIdentity
	Driver   []byte
	Traveler []byte
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	DriverNonce   []byte
	TravelerNonce []byte
}

func (m *SignatureHeader) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Common/Serialize")
	}
	return js
}

func DeSerializeSignatureHeader(data []byte) *SignatureHeader {
	var m *SignatureHeader
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&m)
	return m
}
