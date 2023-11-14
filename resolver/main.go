package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/jihedmastouri/marsoul/resolver/internal"
	"github.com/jihedmastouri/marsoul/resolver/transport"
)

func main() {
	config, err := internal.NewConfigs()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)

	cert, err := tls.LoadX509KeyPair(config.CertLoc, config.PrvKeyLoc)
	if err != nil {
		fmt.Println("Error loading server key pair:", err)
		return
	}

	laddr := fmt.Sprintf("%s:%d", config.Addr, config.Port)

	lis, err := tls.Listen("tcp", laddr, &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	header := make([]byte, 1)
	_, err := io.ReadFull(conn, header)
	if err != nil {
		log.Fatal(err)
	}

	msgType := transport.MessageType(header[0])
	log.Println(msgType)

	switch msgType {
	case transport.SaveRq:
		var payload transport.SaveRqPayload
		if err = transport.Decode(&conn, &payload); err != nil {
			log.Fatal(err)
		}
		fmt.Println(payload)

	case transport.RetrRq:
		var payload transport.RetrRqPayload
		if err = transport.Decode(&conn, &payload); err != nil {
			log.Fatal(err)
		}
		fmt.Println(payload)

	case transport.Ping:
		if _, err := conn.Write([]byte("Pong")); err != nil {
			log.Fatal(err)
		}

	}

	if _, err := conn.Write([]byte("bye")); err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
