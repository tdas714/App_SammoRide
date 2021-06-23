package client

import (
	"encoding/json"
	"io"
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
	js, err := json.Marshal(ra)
	CheckErr(err, "ContactSer/encode")

	return js
}

func ContractDeserialize(data io.Reader) *OrderContract {
	var order *OrderContract
	json.NewDecoder(data).Decode(&order)

	return order
}
