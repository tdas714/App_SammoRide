package state

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type WorldState struct {
	States map[string][]interface{} // Might have to change this
	Path   string
}

func (ws *WorldState) Close() {
	err := ioutil.WriteFile(ws.Path, ws.Serialize(), 0700)
	if err != nil {
		log.Panic(err, "WSClose")
	}

}

func (ws *WorldState) SetValue(key string, value interface{}, version int) {
	ws.States[key] = []interface{}{value, version}
}

func (ws *WorldState) Query(key string) []interface{} {
	re, ok := ws.States[key]
	if ok {
		return nil
	}
	// value, version := re
	return re
}

func (ws *WorldState) LoadState(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err, "LoadWorldState")
	}
	ws = StateDeserialize(data)
}

func (ws *WorldState) Serialize() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(ws)
	if err != nil {
		log.Panic(err, "Serialize/WorldState")
	}
	return buf.Bytes()
}

func StateDeserialize(data []byte) *WorldState {
	var ws *WorldState
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	err := dec.Decode(&ws)
	if err != nil {
		log.Panic(err, "Deserialize/WorldState")
	}
	return ws
}
