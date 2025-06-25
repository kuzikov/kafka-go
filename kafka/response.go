package kafka

import (
	"bytes"
	"encoding/binary"
	"log"
)

const msgSize int32 = 4

type Response struct {
	// message_size
	messageSize int32
	// header v0
	correlationID int32
	// body
}

func NewResponse(correlationID int32) *Response {
	return &Response{
		messageSize:   msgSize,
		correlationID: correlationID,
	}
}

func (r *Response) Bytes() []byte {
	var b []byte
	buf := bytes.NewBuffer(b)

	if err := binary.Write(buf, binary.BigEndian, r.messageSize); err != nil {
		log.Printf("Write message_size: %s\n", err.Error())
	}

	if err := binary.Write(buf, binary.BigEndian, r.correlationID); err != nil {
		log.Printf("Write correlationID: %s\n", err.Error())
	}

	return buf.Bytes()
}
