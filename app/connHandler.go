package main

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/kafka"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 64)

	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading connection:%v\n", err)
		return
	}

	if n < 12 {
		panic("Wrong header length")
	}

	rawCorrelationID := buf[8:12]
	correlationId := binary.BigEndian.Uint32(rawCorrelationID)

	resp := kafka.NewResponse(int32(correlationId))

	conn.Write(resp.Bytes())
}
