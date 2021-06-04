package network

import "github.com/panjf2000/gnet"

type NetCodec interface {
	gnet.ICodec
}

// ICodec interface {
//		// Encode encodes frames upon server responses into TCP stream.
//		Encode(c Conn, buf []byte) ([]byte, error)
//		// Decode decodes frames from TCP stream via specific implementation.
//		Decode(c Conn) ([]byte, error)
//	}
