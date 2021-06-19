package gossip

import (
	"crypto/sha256"
	"strings"

	"github.com/App-SammoRide/client"
)

type Gossip struct {
	Header    string
	data      []byte
	KnownList []string
}

func (g *Gossip) Exect(data []byte, subDomain, Certificatepath, KeyPath string,
	conn *client.Connections, input []string) {

	pList := conn.GetRandom(1)
	header := sha256.Sum256(data)

	if client.Contains(g.KnownList, string(header[:])) {
		return
	}

	conn.Add(input)

	for _, p := range pList {
		splited := strings.Split(p, ":")

		g.KnownList = append(g.KnownList, string(header[:]))

		client.SendData("CAs/rootCa.crt",
			Certificatepath, KeyPath, splited[0], splited[1],
			subDomain, data, 10)
	}
}
