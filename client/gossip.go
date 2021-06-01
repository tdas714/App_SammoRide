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

func SavePeerList(dir string, peerList []string) {
	var curlen int64

	db, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("../%s/PeerList", dir)))
	CheckErr(err, "SavePeerList/db")
	defer db.Close()

	filename := fmt.Sprintf("../%s/PeerList.len", dir)

	if fileExists(filename) {
		lenData, err := ioutil.ReadFile(filename)
		CheckErr(err, "SavePeerList/lenData")
		curlen = ByteArrayToInt(lenData)
	} else {
		curlen = 0
	}

	err = db.Update(func(txn *badger.Txn) error {
		for _, p := range peerList {
			curlen = curlen + 1
			err = txn.Set(IntToByteArray(curlen), []byte(p))
			CheckErr(err, "savePeerList/Set")
		}
		err = ioutil.WriteFile(filename,
			IntToByteArray(curlen), 0700)
		return err
	})
	CheckErr(err, "SavePeerList/Update")
	db.Close()
}
