package client

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
	"math/big"
)

type OrderContract struct {
	PickupLoc   string
	DestLoc     string
	Traveler    *ClientInfo
	TravelerSig []big.Int
	RideFair    float32
	ArrivalTime string
	Driver      *ClientInfo
	DriverSig   []big.Int
}

type Sig struct {
	r *big.Int
	s *big.Int
}

func (ra *OrderContract) ContractSerialize() []byte {
	js, err := json.Marshal(ra)
	CheckErr(err, "ContactSer/encode")

	return js
}

func ContractDeserialize(data io.Reader) *OrderContract {
	var order *OrderContract
	json.NewDecoder(data).Decode(&order)

	return order
}

func ContractFromBytes(data []byte) *OrderContract {
	var gData OrderContract

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&gData)

	CheckErr(err, "RAD/decode")

	return &gData
}
