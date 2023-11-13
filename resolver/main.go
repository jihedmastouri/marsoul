package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/jihedmastouri/marsoul/resolver/internal"
	"github.com/jihedmastouri/marsoul/resolver/pkg"
)

func main() {
	config, err := internal.NewConfigs()
	if err != nil {
		log.Fatal(err)
	}

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

	msgType := pkg.MessageType(header[0])

	switch msgType {
	case pkg.SaveRq:
		buf, err := io.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}

		var payload pkg.SaveRqPayload
		json.Unmarshal(buf[1:], &payload)

		fmt.Println(payload)

	case pkg.RetrRq:
		buf, err := io.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}

		var payload pkg.RetrRqPayload
		json.Unmarshal(buf[1:], &payload)

		fmt.Println(payload)

	case pkg.Ping:
		if _, err := conn.Write([]byte("Pong")); err != nil {
			log.Fatal(err)
		}
	}

	if _, err := conn.Write([]byte("bye")); err != nil {
		log.Fatal(err)
	}
}
