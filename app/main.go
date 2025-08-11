package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start listening on :9092")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		return
	}

	handleConn(conn)

}
