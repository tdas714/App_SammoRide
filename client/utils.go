package client

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
	"unsafe"
)

const (
	ENROLL_REQ = "EnrollRequest"
	ENROLL_RES = "EnrollResponce"
)

type PeerEnrollDataRequest struct {
	Country     string
	Name        string
	Province    string
	IpAddr      string
	City        string
	PostalCode  string
	ListingPort string
}

type PeerEnrollDataResponse struct {
	Header     string
	IpAddr     string
	PeerCert   []byte
	PrivateKey []byte
	SenderCert []byte
	RootCert   []byte
	PeerList   []string
}

type RiderAnnouncement struct {
	Latitude    string
	Longitude   string
	Avalability string
	Info        *ClientInfo
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	l := make([]string, 0)
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				l = append(l, ipnet.IP.String())
			}
		}
	}
	return l[0]
}

func LoadPrivateKey(f []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode([]byte(string(f)))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	return privateKey
}

func LoadCertificate(f []byte) *x509.Certificate {
	block, _ := pem.Decode(f)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return cert
}

func CheckErr(err error, origin string) {
	if err != nil {
		log.Fatalf("%s - %s", origin, err)
	}
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func IntToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func ByteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (ra *RiderAnnouncement) RASerialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(ra)

	CheckErr(err, "RAS/encode")

	return res.Bytes()
}

func RADeserialize(data []byte) *RiderAnnouncement {
	var riderA RiderAnnouncement

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&riderA)

	CheckErr(err, "RAD/decode")

	return &riderA
}

type Connections struct {
	Path     string
	PeerList []string
}

func LoadConnections(path string) *Connections {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return &Connections{Path: path}
	}
	return ConnDeserialize(content)
}

func ConnDeserialize(data []byte) *Connections {
	var conn Connections

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&conn)

	CheckErr(err, "Conn/decode")

	return &conn
}

func (c *Connections) Add(pLIst []string) {
	c.PeerList = append(c.PeerList, pLIst...)
}

func (c *Connections) len() int {
	return len(c.PeerList)
}

func (c *Connections) GetRandom(num int) []string {
	var gList []string
	var selectedPeer string
	for i := 1; i >= c.len(); i++ {
		rand.Seed(time.Now().Unix())
		selectedPeer = c.PeerList[rand.Intn(c.len())]
		gList = append(gList, selectedPeer)
	}
	return gList
}

func (c *Connections) Close() {
	bytes, err := GetBytes(c)
	CheckErr(err, "GetBytes/C.Close")
	err = ioutil.WriteFile(c.Path, bytes, 0700)
	CheckErr(err, "Write/C.Close")
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
