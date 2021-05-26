package client

import (
	"math/rand"
	"time"

	"github.com/dgraph-io/badger"
)

type PeerGossipData struct {
	Header string
	TCP    string
}

func Gossip(db *badger.DB, lenData int, data []byte) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	selectedId := rand.Intn(lenData)

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(string(selectedId)))
		CheckErr(err, "Gossip/item")
		val, err := item.Value()
		return err
	})
	CheckErr(err, "Gossip/dbView")

}
