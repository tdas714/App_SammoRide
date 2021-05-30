package client

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/dgraph-io/badger/v3"
)

func GetPeerList(lenList int, dir string) []string {

	db, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("../%s/PeerList", dir)))
	CheckErr(err, "getPeerList/db")
	defer db.Close()

	lenData, err := ioutil.ReadFile(fmt.Sprintf("../%s/PeerList.len", dir))
	CheckErr(err, "GetPeerList/lenData")

	var peerList []string

	err = db.View(func(txn *badger.Txn) error {
		for i := 0; i < lenList; i++ {

			if i+1 == int(ByteArrayToInt(lenData)) {
				break
			}

			rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
			selectedId := rand.Intn(int(ByteArrayToInt(lenData)))

			item, err := txn.Get([]byte(string(selectedId)))
			CheckErr(err, "GetPeerList/item")
			err = item.Value(func(val []byte) error {
				peerList = append(peerList, string(val))
				return nil
			})
		}

		return err
	})

	CheckErr(err, "GetPeerList/dbView")
	return peerList
}

func SavePeerList(db *badger.DB, peerList []string) {

	lenData, err := ioutil.ReadFile(("../database/PeerList.len"))
	CheckErr(err, "savePeerList/Lendata")

	curlen := ByteArrayToInt(lenData)

	err = db.Update(func(txn *badger.Txn) error {
		for _, p := range peerList {
			curlen = curlen + 1
			err = txn.Set(IntToByteArray(curlen), []byte(p))
			CheckErr(err, "savePeerList/Set")
		}
		err = ioutil.WriteFile("../database/PeerList.len",
			IntToByteArray(curlen), 0700)
		return err
	})
	CheckErr(err, "SavePeerList/Update")

}
