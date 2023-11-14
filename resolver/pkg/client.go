package pkg

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
)

type ClientRequester interface {
	Save(SaveRqPayload) (string, error)
	Retr(RetrRqPayload) (string, error)
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
	fileNodeAddr, err := appDial(c.resolverAddr, SaveRq, c.cert, payload)
	if err != nil {
		return "", err
	}

	return fileNodeAddr, nil
}

func (c *DefaultClient) Retr(payload RetrRqPayload) (string, error) {
	fileNodeAddr, err := appDial(c.resolverAddr, RetrRq, c.cert, payload)
	if err != nil {
		return "", err
	}

	return fileNodeAddr, nil
}

func appDial[Payload SaveRqPayload | RetrRqPayload](addr string, header MessageType, cert tls.Certificate, payload Payload) (string, error) {
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

	b = append([]byte{byte(header)}, b...)
	n, err := conn.Write([]byte(b))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s on byte: %d", err, n))
	}
	conn.CloseWrite()

	buf, err := io.ReadAll(conn)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s on byte: %d", err, n))
	}

	return string(buf), nil
}
