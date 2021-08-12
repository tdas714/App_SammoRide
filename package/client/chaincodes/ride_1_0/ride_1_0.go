package ride_1_0

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/App-SammoRide/package/client/struct/common"
	"github.com/App-SammoRide/package/client/struct/peer"
)

type RideType int32

const (
	RIDE_COMPLETE            RideType = 1
	RIDE_REFUND_DRIVER       RideType = 2
	RIDE_FULLREFUND_TRAVELER RideType = 3
	RIDE_TYPE_UNKNOWN        RideType = 0
)

type Ride struct {
	TxID                string //Hash
	TxTimeStamp         time.Time
	PickupLocation      string
	DestinationLocation string
	DriverPublicKey     string
	TravelerPublicKey   string
	ArrivalTime         time.Duration
	RideFair            float32
}

func (ride *Ride) Resolution(currentDloc, currentTLoc string, txcommitted bool) RideType {
	if currentDloc == ride.DestinationLocation && currentTLoc == ride.DestinationLocation {
		return RIDE_COMPLETE
	}
	if txcommitted && time.Now().After(ride.TxTimeStamp.Add(ride.ArrivalTime).Add(time.Minute*15)) && currentDloc != ride.PickupLocation {
		return RIDE_FULLREFUND_TRAVELER
	}
	if txcommitted && currentDloc == ride.PickupLocation && time.Now().After(ride.TxTimeStamp.Add(ride.ArrivalTime).Add(time.Minute*15)) &&
		currentTLoc != ride.PickupLocation {
		return RIDE_REFUND_DRIVER
	}
	return RIDE_TYPE_UNKNOWN
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

func Execute(pickupLoc, destiLoc string, ridefair float32, txid string, driverPub, travelerPub string, codeId *peer.ChaincodeID,
	arivT time.Duration, txStamp time.Time) ([]byte, []byte, int) {
	kvread := common.KVRead{Key: "", Version: nil}
	ride := Ride{TxID: txid, TxTimeStamp: txStamp, PickupLocation: pickupLoc, DestinationLocation: destiLoc, DriverPublicKey: driverPub, TravelerPublicKey: travelerPub}

	value := make(map[string]string)
	value["TxID"] = txid
	value["PickUpLocation"] = pickupLoc
	value["DestinationLocation"] = destiLoc
	value["DriverPublicKey"] = driverPub
	value["TravelerPublicKey"] = travelerPub
	value["ArrivalTime"] = arivT.String()
	value["RideFair"] = fmt.Sprintf("%f", ridefair)

	kvwrite := common.KVWrite{Key: txid, Value: value, IsDelete: false}

	kvrwSet := common.KVRWSet{Reads: []*common.KVRead{&kvread}, Writes: []*common.KVWrite{&kvwrite}}

	// Construct complete, DriveCancle, TravelerCancle
	// completeCode := "IF DRIVER LOC, TRAVELERLOC, DESTINATIONLOC SAME OR ISSUER TRAVELER"
	// // travelerCancle: INVOKE COMLETE->IF TRANSACTION COMMITTED;REFUND 70% TRAVELER
	// // driverCancle: IF DRIVERLOC, PICKUPLOC SAME->REFUND 100% TRAVELER
	// disputeCode := "IF TRANSACTION-COMMITED AND DRIVER-LOC SAME PICKUP-LOC->(REFUND 70%;6% COMMISION;24% DRIVER);" +
	// 	"ELSE IF- EXCEEDS ARRIVAL-TIME AND DRIVER-LOC NOT SAME PICKUP-LOC->(100% REFUND);" +
	// 	"ELSE IF- DRIVER-LOC SAME PICKUP LOC AND ARRIVAL-TIME EXCEEDS->(REFUND 70%;6% COMMISION;24% DRIVER);"
	codeEventRide := peer.ChaincodeEvent{ChaincodeId: codeId, TxId: txid, EventName: "Ride", Payload: ride.Serialize()}

	// chaincodeEvents := []*peer.ChaincodeEvent{&codeEventRide}
	// fmt.Println(chaincodeEvents)
	var events peer.EventsStruct
	events.ChaincodeEvents = make([]*peer.ChaincodeEvent, 1)
	events.ChaincodeEvents = []*peer.ChaincodeEvent{&codeEventRide}

	return kvrwSet.Serialize(), events.Serialize(), 200

}
