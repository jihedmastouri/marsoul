package pkg

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type ClientRequester interface {
	Save() string
	Retr() string
}

type DefaultClient struct {
	resolverAddr string
	cert         tls.Certificate
}

func NewClient(addr, certFile, keyFile string) *DefaultClient {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	return &DefaultClient{
		resolverAddr: addr,
		cert:         cert,
	}
}

func (c *DefaultClient) Save(payload SaveRqPayload) (string, error) {
	fileNodeAddr, err := appDial(c.resolverAddr, c.cert, payload)
	if err != nil {
		return "", err
	}

	return fileNodeAddr, nil
}

func (c *DefaultClient) Retr(payload RetrRqPayload) (string, error) {
	fileNodeAddr, err := appDial(c.resolverAddr, c.cert, payload)
	if err != nil {
		return "", err
	}

	return fileNodeAddr, nil
}

func appDial[Payload SaveRqPayload | RetrRqPayload](addr string, cert tls.Certificate, payload Payload) (string, error) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}

	conn, err := tls.Dial("tcp", addr, conf)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	n, err := conn.Write([]byte(b))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s on byte: %d", err, n))
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s on byte: %d", err, n))
	}

	return string(buf), nil
}
