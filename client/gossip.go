package client

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgraph-io/badger/v3"
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
		err = item.Value(func(val []byte) error {
			fmt.Println(val)
			return nil
		})
		return err
	})

	CheckErr(err, "Gossip/dbView")

}
