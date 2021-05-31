package sample

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/Terry-Mao/goim/pkg/bufio"
	"github.com/Terry-Mao/goim/pkg/bytes"
	"github.com/Terry-Mao/goim/pkg/encoding/binary"
	"github.com/Terry-Mao/goim/pkg/websocket"
	log "github.com/golang/glog"
)

const (
	OP_HANDSHARE        = int32(0)
	OP_HANDSHARE_REPLY  = int32(1)
	OP_HEARTBEAT        = int32(2)
	OP_HEARTBEAT_REPLY  = int32(3)
	OP_SEND_SMS         = int32(4)
	OP_SEND_SMS_REPLY   = int32(5)
	OP_DISCONNECT_REPLY = int32(6)
	OP_AUTH             = int32(7)
	OP_AUTH_REPLY       = int32(8)
	OP_TEST             = int32(254)
	OP_TEST_REPLY       = int32(255)
)

const (
	rawHeaderLen = int16(16)
)

const (
	ProtoTCP          = 0
	ProtoWebsocket    = 1
	ProtoWebsocketTLS = 2
)

type Proto struct {
	Ver       int32           `json:"ver"`  // protocol version
	Operation int32           `json:"op"`   // operation for request
	SeqId     int32           `json:"seq"`  // sequence number chosen by client
	Body      json.RawMessage `json:"body"` // binary body bytes(json.RawMessage is []byte)
}

//func (p *Proto) Print() {
//	log.Info("\n-------- proto --------\nver: %d\nop: %d\nseq: %d\nbody: %s\n", p.Ver, p.Operation, p.SeqId, string(p.Body))
//}

const (
	// MaxBodySize max proto body size
	MaxBodySize = int32(1 << 12)
)

const (
	// size
	//_packSize      = 4
	//_headerSize    = 2
	//_verSize       = 2
	//_opSize        = 4
	//_seqSize       = 4
	_heartSize = 4
	//_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	//_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	//// offset
	//_packOffset   = 0
	//_headerOffset = _packOffset + _packSize
	//_verOffset    = _headerOffset + _headerSize
	//_opOffset     = _verOffset + _verSize
	//_seqOffset    = _opOffset + _opSize
	//_heartOffset  = _seqOffset + _seqSize

	_mars_headerLen     = 4
	_mars_versionLen    = 4
	_mars_cmdidLen      = 4
	_mars_seqLen        = 4
	_mars_bodyLen       = 4
	_mars_rawHeaderSize = _mars_headerLen + _mars_versionLen + _mars_cmdidLen + _mars_seqLen + _mars_bodyLen
	_marsmaxPackSize    = MaxBodySize + int32(_mars_rawHeaderSize)
	//_heartSize     = 4

	// marsoffset
	_mars_packOffset   = 0
	_mars_headerOffset = _mars_packOffset + _mars_headerLen
	_mars_verOffset    = _mars_headerOffset + _mars_versionLen
	_mars_opCmidOffset = _mars_verOffset + _mars_cmdidLen
	_mars_seqOffset    = _mars_opCmidOffset + _mars_seqLen
	_mars_bodyOffset   = _mars_seqOffset + _mars_bodyLen
)

var (
	// ErrProtoPackLen proto packet len error
	ErrProtoPackLen = errors.New("default server codec pack length error")
	// ErrProtoHeaderLen proto header len error
	ErrProtoHeaderLen = errors.New("default server codec header length error")
)

// WriteTo write a proto to bytes writer.
func (p *Proto) WriteTo(b *bytes.Writer) {
	var (
		buf = b.Peek(_mars_rawHeaderSize)
	)
	binary.BigEndian.PutInt32(buf[_mars_packOffset:], _mars_rawHeaderSize)
	binary.BigEndian.PutInt32(buf[_mars_headerOffset:], p.Ver)
	binary.BigEndian.PutInt32(buf[_mars_verOffset:], p.Operation)
	binary.BigEndian.PutInt32(buf[_mars_opCmidOffset:], p.SeqId)
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], int32(len(p.Body)))

	if p.Body != nil {
		b.Write(p.Body)
	}
}

// ReadTCP read a proto from TCP reader.
func (p *Proto) ReadTCP(rr *bufio.Reader) (err error) {
	var (
		//bodyLen   int
		//headerLen int16
		//packLen   int32
		buf []byte
	)

	if buf, err = rr.Pop(_mars_rawHeaderSize); err != nil {
		log.Errorf("tcp request bin (%s) err(%v) ", hex.Dump(buf), err.Error())
		return
	}
	headerLen := binary.BigEndian.Int32(buf[_mars_packOffset:_mars_headerOffset])
	p.Ver = binary.BigEndian.Int32(buf[_mars_headerOffset:_mars_verOffset])
	p.Operation = binary.BigEndian.Int32(buf[_mars_verOffset:_mars_opCmidOffset])
	p.SeqId = binary.BigEndian.Int32(buf[_mars_opCmidOffset:_mars_seqOffset])
	bodyLen := binary.BigEndian.Int32(buf[_mars_seqOffset:_mars_bodyOffset])
	if headerLen != _mars_rawHeaderSize {
		log.Errorf("tcp request bin (%s) headerLen (%d), not equals to _mars_rawHeaderSize (%d)   ", hex.Dump(buf), headerLen, _mars_rawHeaderSize)
	}
	p.Body, err = rr.Pop(int(bodyLen))
	bodyLenRaw := len(p.Body)
	if bodyLenRaw != int(bodyLen) {
		log.Errorf("tcp request bin (%s) bodyLenRaw (%d), not equals to bodyLen (%d)   ", hex.Dump(buf), bodyLenRaw, bodyLen)

		return ErrProtoHeaderLen
	}
	if bodyLenRaw > int(MaxBodySize) {
		return ErrProtoPackLen
	}

	return
}

// WriteTCP write a proto to TCP writer.
func (p *Proto) WriteTCP(wr *bufio.Writer) (err error) {
	var (
		buf []byte
		//packLen int32
	)
	//9 means raw
	if p.Operation == 9 {
		// write without buffer, job concact proto into raw buffer
		_, err = wr.WriteRaw(p.Body)
		return
	}
	//packLen = _mars_rawHeaderSize + int32(len(p.Body))
	if buf, err = wr.Peek(_mars_rawHeaderSize); err != nil {
		return
	}
	binary.BigEndian.PutInt32(buf[_mars_packOffset:], _mars_rawHeaderSize)
	binary.BigEndian.PutInt32(buf[_mars_headerOffset:], p.Ver)
	binary.BigEndian.PutInt32(buf[_mars_verOffset:], p.Operation)
	binary.BigEndian.PutInt32(buf[_mars_opCmidOffset:], p.SeqId)
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], int32(len(p.Body)))
	if p.Body != nil {
		_, err = wr.Write(p.Body)
		wr.Flush()
	}
	return
}

// WriteTCPHeart write TCP heartbeat with room online.
func (p *Proto) WriteTCPHeart(wr *bufio.Writer, online int32) (err error) {
	var (
		buf     []byte
		packLen int
	)
	packLen = _mars_rawHeaderSize + _heartSize
	if buf, err = wr.Peek(packLen); err != nil {
		return
	}
	// header
	binary.BigEndian.PutInt32(buf[_mars_packOffset:], _mars_rawHeaderSize)
	binary.BigEndian.PutInt32(buf[_mars_headerOffset:], p.Ver)
	binary.BigEndian.PutInt32(buf[_mars_verOffset:], p.Operation)
	binary.BigEndian.PutInt32(buf[_mars_opCmidOffset:], p.SeqId)
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], _heartSize)
	// body
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], online)
	return
}

// ReadWebsocket read a proto from websocket connection.
func (p *Proto) ReadWebsocket(ws *websocket.Conn) (err error) {
	var (
		//bodyLen   int32
		//headerLen int32
		//packLen int32
		buf []byte
	)
	if _, buf, err = ws.ReadMessage(); err != nil {
		return
	}
	if len(buf) < _mars_rawHeaderSize {
		return ErrProtoPackLen
	}

	headerLen := binary.BigEndian.Int32(buf[_mars_packOffset:_mars_headerOffset])
	p.Ver = binary.BigEndian.Int32(buf[_mars_headerOffset:_mars_verOffset])
	p.Operation = binary.BigEndian.Int32(buf[_mars_verOffset:_mars_opCmidOffset])
	p.SeqId = binary.BigEndian.Int32(buf[_mars_opCmidOffset:_mars_seqOffset])
	bodyLen := binary.BigEndian.Int32(buf[_mars_seqOffset:_mars_bodyOffset])
	//packLen = headerLen + bodyLen
	if headerLen != _mars_rawHeaderSize {
		log.Errorf(" headerLen (%d), not equals to _mars_rawHeaderSize (%d) tcp request bin (%s)  ", headerLen, _mars_rawHeaderSize, hex.Dump(buf))
	}
	p.Body = buf[headerLen:(bodyLen + headerLen)]
	bodyLenRaw := len(p.Body)

	if bodyLenRaw != int(bodyLen) {
		log.Errorf("tcp request bin (%s) bodyLenRaw (%d), not equals to bodyLen (%d)   ", bodyLenRaw, bodyLen, hex.Dump(buf))

		return ErrProtoHeaderLen
	}
	if bodyLenRaw > int(MaxBodySize) {
		return ErrProtoPackLen
	}

	return
}

// WriteWebsocket write a proto to websocket connection.
func (p *Proto) WriteWebsocket(ws *websocket.Conn) (err error) {
	var (
		buf     []byte
		packLen int
	)
	packLen = _mars_rawHeaderSize + len(p.Body)
	if err = ws.WriteHeader(websocket.BinaryMessage, packLen); err != nil {
		return
	}
	if buf, err = ws.Peek(_mars_rawHeaderSize); err != nil {
		return
	}

	binary.BigEndian.PutInt32(buf[_mars_packOffset:], _mars_rawHeaderSize)
	binary.BigEndian.PutInt32(buf[_mars_headerOffset:], p.Ver)
	binary.BigEndian.PutInt32(buf[_mars_verOffset:], p.Operation)
	binary.BigEndian.PutInt32(buf[_mars_opCmidOffset:], p.SeqId)
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], int32(len(p.Body)))

	if p.Body != nil {
		err = ws.WriteBody(p.Body)
	}
	return
}

func (p *Proto) GetBuf() []byte {
	return p.Body
}

// WriteWebsocketHeart write websocket heartbeat with room online.
func (p *Proto) WriteWebsocketHeart(wr *websocket.Conn, online int32) (err error) {
	var (
		buf     []byte
		packLen int
	)
	packLen = _mars_rawHeaderSize + _heartSize
	// websocket header
	if err = wr.WriteHeader(websocket.BinaryMessage, packLen); err != nil {
		return
	}
	if buf, err = wr.Peek(packLen); err != nil {
		return
	}
	// proto header
	binary.BigEndian.PutInt32(buf[_mars_packOffset:], _mars_rawHeaderSize)
	binary.BigEndian.PutInt32(buf[_mars_headerOffset:], p.Ver)
	binary.BigEndian.PutInt32(buf[_mars_verOffset:], p.Operation)
	binary.BigEndian.PutInt32(buf[_mars_opCmidOffset:], p.SeqId)
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], _heartSize)
	// proto body
	binary.BigEndian.PutInt32(buf[_mars_seqOffset:], online)
	return
}
