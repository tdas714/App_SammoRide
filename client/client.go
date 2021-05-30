package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type ClientInfo struct {
	Country    string `yaml:"Country"`
	Name       string `yaml:"Name"`
	Province   string `yaml:"Province"`
	City       string `yaml:"City"`
	IP         string `yaml:"IP"`
	Postalcode string `yaml:"PostalCode"`
}

func (c *ClientInfo) GetConf(filename string) *ClientInfo {

	yamlFile, err := ioutil.ReadFile(filename)
	CheckErr(err, "YamlFile Get")
	err = yaml.Unmarshal(yamlFile, c)
	CheckErr(err, " Unmarshal Error")

	return c
}

func SendEnrollRequest(country, name, province, city, postC,
	ipAddr, serverIp, serverPort string) {
	enrollReq := &PeerEnrollDataRequest{Country: country, Name: name, Province: province, IpAddr: ipAddr,
		City: city, PostalCode: postC}
	json_data, err := json.Marshal(enrollReq)
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf(fmt.Sprintf("http://%s:%s/post",
		serverIp, serverPort))

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(json_data))
	CheckErr(err, "SendEnrollRequest/Post")

	var res *PeerEnrollDataResponse
	json.NewDecoder(resp.Body).Decode(&res)

	VerifyPeer(res.RootCert, res.SenderCert, res.PeerCert, res.PrivateKey)
	VerifyOrderer(res.RootCert, res.SenderCert)
	//
	_ = os.Mkdir("PeerCerts", 0700)
	_ = os.Mkdir("CAs", 0700)

	err = ioutil.WriteFile("CAs/interCa.crt", res.SenderCert, 0700)
	CheckErr(err, "SendErollReq/interca")
	err = ioutil.WriteFile("CAs/rootCa.crt", res.RootCert, 0700)
	CheckErr(err, "SendErollReq/rootca")

	blocks, _ := pem.Decode(res.PeerCert)
	if blocks == nil {
		log.Panic("Block is nil")
	}
	// Public key
	certOut, err := os.Create(fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.crt",
		country, name, province,
		city))

	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: blocks.Bytes})
	certOut.Close()
	log.Print("written cert.pem\n")

	// Private Key
	keyOut, err := os.OpenFile(fmt.Sprintf("PeerCerts/%s_%s_%s_%s_Cert.key",
		country, name, province,
		city),
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	CheckErr(err, "SendEnrollRequest/keyOut")
	err = pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: res.PrivateKey})
	CheckErr(err, "SendEnrollRequest/penEncode")
	if err := keyOut.Close(); err != nil {
		log.Fatalf("Error closing key.pem: %v", err)
	}
	log.Print("wrote key.pem\n")
}

// ====================================
func createClientConfig(ca, crt, key string) (*tls.Config, error) {
	caCertPEM, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(crt, key)
	CheckErr(err, "createClientConfig/cert")

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    roots,
		RootCAs:      roots, // This Needs ATTENTION
	}, nil
}

func gossipHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	CheckErr(err, "gossipRequest")
	fmt.Println("Poking from ", string(bodyBytes))
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Lets gossip!\n")
}

func StartPeerServer(ipAddr, port, caPath, crtPath, keyPath string) {
	// Set up a /hello resource handler
	http.HandleFunc("/gossip", gossipHandler) //This can be diffrent for diffrent data types

	tlsConfig, err := createClientConfig(caPath, crtPath, keyPath)
	CheckErr(err, "StartOrederServer/config")

	server := &http.Server{
		Addr:      ipAddr + ":" + port,
		TLSConfig: tlsConfig,
	}

	// Listen to HTTPS connections with the server certificate and wait
	log.Fatal(server.ListenAndServeTLS(crtPath, keyPath))

}

func SendData(addr, ca, crt, key, ipAddr, port,
	reqSubDomain string, data []byte) {
	// addr := *connect
	if !strings.Contains(addr, ":") {
		addr += ":8443"
	}

	// Read the key pair to create certificate
	cert, err := tls.LoadX509KeyPair(crt, key)
	CheckErr(err, "SendData/cert")

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile(ca)
	CheckErr(err, "SendData/cacert")

	caCertPool, err := x509.SystemCertPool()
	CheckErr(err, "Client/CertPool")
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}

	// Request /hello via the created HTTPS client over port 8443 via GET
	// r, err := client.Get(fmt.Sprintf("https://%s/hello", addr))
	// CheckErr(err, "SendData/r")

	// =======POST
	url := fmt.Sprintf("https://%s:%s/%s", ipAddr,
		port, reqSubDomain)
	r, err := client.Post(url, "application/json", bytes.NewBuffer(data))

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}
