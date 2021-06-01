package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GossipHandler(w http.ResponseWriter, r *http.Request, name string) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	CheckErr(err, "gossipRequest")
	SavePeerList("database", name, []string{string(bodyBytes)})
	// Write "Hello, world!" to the response body
	fmt.Println("Gossip from: ", r.RemoteAddr)
	io.WriteString(w, "Lets gossip!\n")
}

func RiderAHandler(w http.ResponseWriter, r *http.Request, name string) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	CheckErr(err, "RiderAnnonreq")
	riderA := RADeserialize(bodyBytes)
	SavePeerList("database", name, []string{riderA.Info.IP + ":" + riderA.Info.Port})
	Gossip(bodyBytes, riderA.Info, "Announcement/rider")

}
