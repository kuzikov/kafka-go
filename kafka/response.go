package kafka

import (
	"bytes"
	"encoding/binary"
	"log"
)

const msgSize int32 = 4

// `message_size`
// `header v0`
// `body`
type Response struct {
	messageSize int32
	header      responseHeaderV0
}

type responseHeaderV0 struct {
	correlationID int32
}

func NewResponse(correlationID int32) *Response {
	return &Response{
		messageSize: msgSize,
		header: responseHeaderV0{
			correlationID: correlationID,
		},
	}
}

func (r *Response) Bytes() []byte {
	var b []byte
	buf := bytes.NewBuffer(b)

	if err := binary.Write(buf, binary.BigEndian, r.messageSize); err != nil {
		log.Printf("Write message_size: %s\n", err.Error())
	}

	if err := binary.Write(buf, binary.BigEndian, r.header.correlationID); err != nil {
		log.Printf("Write correlationID: %s\n", err.Error())
	}

	return buf.Bytes()
}
