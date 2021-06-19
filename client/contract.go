package client

import (
	"bytes"
	"encoding/gob"
	"math/big"
)

type OrderContract struct {
	PickupLoc   string
	DestLoc     string
	Traveler    *ClientInfo
	TravelerSig []byte
	RideFair    float32
	ArrivalTime string
	Driver      *ClientInfo
	DriverSig   *Sig
}

type Sig struct {
	r *big.Int
	s *big.Int
}

func (ra *OrderContract) ContractSerialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(ra)

	CheckErr(err, "ContactSer/encode")

	return res.Bytes()
}

func ContractDeserialize(data []byte) *OrderContract {
	var order OrderContract

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&order)

	CheckErr(err, "RAD/decode")

	return &order
}
