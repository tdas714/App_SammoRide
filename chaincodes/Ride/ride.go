package ride

import "fmt"

func StartRide(pickupLoc, destiLoc string, ridefair float32) {
	fmt.Println("PickUp Location: ", pickupLoc)
	fmt.Println("Destination Location: ", destiLoc)
	fmt.Println("Ride Fair: ", ridefair)
}

func GetRide() map[string]interface{} {
	m := make(map[string]interface{})
	m["Start"] = StartRide
	return m
}
