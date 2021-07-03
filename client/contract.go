package client

import (
	"encoding/json"
	"io"
)

var (
	StartRide = byte(0x01)
	StopRide  = byte(0x02)
	Dispute   = byte(0x03)
)

type TransactionProposalResponse struct {
	Msg string
}

func (ra *TransactionProposalResponse) TransResSerialize() []byte {
	js, err := json.Marshal(ra)
	CheckErr(err, "ContactSer/encode")

	return js
}

func TransResDeserialize(data io.Reader) *TransactionProposalResponse {
	var txPropRes *TransactionProposalResponse
	json.NewDecoder(data).Decode(&txPropRes)

	return txPropRes
}

// func ContractDeserialize(data io.Reader) *TransactionProposal {
// 	var txProp *TransactionProposal
// 	json.NewDecoder(data).Decode(&txProp)

// 	return txProp
// }

// func ContractFromBytes(data []byte) *TransactionProposal {
// 	var gData TransactionProposal

// 	decoder := gob.NewDecoder(bytes.NewReader(data))

// 	err := decoder.Decode(&gData)

// 	CheckErr(err, "ContractDS/decode")

// 	return &gData
// }
