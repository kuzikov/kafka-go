package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/kafka-starter-go/kafka"
)

func main() {
	fmt.Println("Start listening on :9092")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(2)
	}

	buf := make([]byte, 64)
	msgLen := 0
	if msgLen, err = conn.Read(buf); err != nil {
		log.Printf("Error reading connection:%v\n", err)
		return
	}

	fmt.Printf("read conn msg: `%v` with length: %d\n", buf[:msgLen], msgLen)

	resp := kafka.NewResponse(7)
	conn.Write(resp.Bytes())
	conn.Close()
}
