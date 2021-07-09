package ride_1_0

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/App-SammoRide/struct/common"
	"github.com/App-SammoRide/struct/peer"
)

type Ride struct {
	TxID                string //Hash
	PickupLocation      string
	DestinationLocation string
	DriverPublicKey     string
	TravelerPublicKey   string
}

func (m *Ride) Serialize() []byte {
	js, err := json.Marshal(m)
	if err != nil {
		log.Panic(err.Error() + " - " + "Ride/Serialize")
	}
	return js
}

func DeSerializeSignedProposal(data io.Reader) *Ride {
	var m *Ride
	json.NewDecoder(data).Decode(&m)
	return m
}

func StartRide(pickupLoc, destiLoc string, ridefair float32, txid string, driverPub, travelerPub string, codeId *peer.ChaincodeID) ([]byte, []byte, int) {
	kvread := common.KVRead{Key: "", Version: nil}
	ride := Ride{TxID: txid, PickupLocation: pickupLoc, DestinationLocation: destiLoc, DriverPublicKey: driverPub, TravelerPublicKey: travelerPub}
	kvwrite := common.KVWrite{Key: txid, Value: ride.Serialize(), IsDelete: false}

	kvrwSet := common.KVRWSet{Reads: []*common.KVRead{&kvread}, Writes: []*common.KVWrite{&kvwrite}}

	// Construct complete, DriveCancle, TravelerCancle
	completeCode := "IF DRIVER LOC, TRAVELERLOC, DESTINATIONLOC SAME OR ISSUER TRAVELER"
	// travelerCancle: INVOKE COMLETE->IF TRANSACTION COMMITTED;REFUND 70% TRAVELER
	// driverCancle: IF DRIVERLOC, PICKUPLOC SAME->REFUND 100% TRAVELER
	disputeCode := "IF TRANSACTION-COMMITED AND DRIVER-LOC SAME PICKUP-LOC->(REFUND 70%;6% COMMISION;24% DRIVER);" +
		"ELSE IF- EXCEEDS ARRIVAL-TIME AND DRIVER-LOC NOT SAME PICKUP-LOC->(100% REFUND);" +
		"ELSE IF- DRIVER-LOC SAME PICKUP LOC AND ARRIVAL-TIME EXCEEDS->(REFUND 70%;6% COMMISION;24% DRIVER);"
	codeEventComplete := peer.ChaincodeEvent{ChaincodeId: codeId, TxId: txid, EventName: "Complete", Payload: []byte(completeCode)}
	codeEventDispute := peer.ChaincodeEvent{ChaincodeId: codeId, TxId: txid, EventName: "Dispute", Payload: []byte(disputeCode)}

	chaincodeEvents := []*peer.ChaincodeEvent{&codeEventDispute, &codeEventComplete}
	fmt.Println(chaincodeEvents)
	var events peer.EventsStruct
	events.ChaincodeEvents = make([]*peer.ChaincodeEvent, 2)
	events.ChaincodeEvents = []*peer.ChaincodeEvent{&codeEventDispute, &codeEventComplete}

	return kvrwSet.Serialize(), events.Serialize(), 200

}
