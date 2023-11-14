package transport

import (
	"encoding/json"
	"io"
	"log"
	"net"
)

func Decode[Payload SaveRqPayload | RetrRqPayload](conn *net.Conn, payload *Payload) error {
	buf, err := io.ReadAll(*conn)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(buf, payload); err != nil {
		return err
	}

	return nil
}
