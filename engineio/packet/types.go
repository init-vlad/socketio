package packet

import "github.com/byinit/chat-service/libs/socketiolib/engineio/frame"

type Frame struct {
	FType frame.Type
	Data  []byte
}

type Packet struct {
	FType frame.Type
	PType Type
	Data  []byte
}
