package client

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v3"
)

func GetPeerList(lenList int, dir, name string) []string {

	db, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("../%s/%s/PeerList", dir, name)))
	CheckErr(err, "getPeerList/db")
	defer db.Close()

	lenData, err := ioutil.ReadFile(fmt.Sprintf("../%s/%s/PeerList.len", dir, name))
	CheckErr(err, "GetPeerList/lenData")

	var peerList []string

	err = db.View(func(txn *badger.Txn) error {
		for i := 0; i < lenList; i++ {

			if i+1 == int(ByteArrayToInt(lenData)) {
				break
			}

			rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
			selectedId := int64(rand.Intn(int(ByteArrayToInt(lenData))))
			// fmt.Println("Selected ", selectedId)
			item, err := txn.Get(IntToByteArray(selectedId))
			CheckErr(err, "GetPeerList/item")
			err = item.Value(func(val []byte) error {
				peerList = append(peerList, string(val))
				return nil
			})
		}

		return err
	})
	db.Close()
	CheckErr(err, "GetPeerList/dbView")
	return peerList
}

func SavePeerList(dir, name string, peerList []string) {
	var curlen int64

	db, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("../%s/%s/PeerList", dir, name)))
	CheckErr(err, "SavePeerList/db")
	defer db.Close()

	filename := fmt.Sprintf("../%s/%s/PeerList.len", dir, name)

	if fileExists(filename) {
		lenData, err := ioutil.ReadFile(filename)
		CheckErr(err, "SavePeerList/lenData")
		curlen = ByteArrayToInt(lenData)
	} else {
		curlen = 0
	}

	err = db.Update(func(txn *badger.Txn) error {
		for _, p := range peerList {
			err = txn.Set(IntToByteArray(curlen), []byte(p))
			CheckErr(err, "savePeerList/Set")
			curlen = curlen + 1
		}
		err = ioutil.WriteFile(filename,
			IntToByteArray(curlen), 0700)
		return err
	})
	CheckErr(err, "SavePeerList/Update")
	db.Close()
}

func Gossip(data []byte, c *ClientInfo, subDomain string) {

	cPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		c.Country, c.Name, c.Province,
		c.City)
	kPath := fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		c.Country, c.Name, c.Province,
		c.City)

	pList := GetPeerList(1, "database", c.Name)
	for _, p := range pList {
		splited := strings.Split(p, ":")

		SendData("CAs/rootCa.crt",
			cPath, kPath, splited[0], splited[1],
			subDomain, data)
	}
}
