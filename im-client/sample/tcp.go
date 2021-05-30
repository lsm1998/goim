package sample

import (
	"encoding/json"
	"fmt"
	"github.com/Terry-Mao/goim/pkg/bufio"

	"github.com/Terry-Mao/goim/api/adapter"
	pb "github.com/golang/protobuf/proto"

	log "github.com/golang/glog"
	"net"
	"time"
)

func EnterRoom(conn net.Conn) {
	//seqId := int32(0)
	//wr := bufio.NewWriter(conn)
	//rd := bufio.NewReader(conn)
	//// reader
	//proto := new(Proto)
	//proto.Ver = 1
	//// auth
	//// test handshake timeout
	//// time.Sleep(time.Second * 31)
	//proto.Operation = OP_AUTH
	//proto.SeqId = seqId
	//rpcInput := adapter.RPCInput{}
	//var token = "{\"mid\":123, \"room_id\":\"live://1000\", \"platform\":\"web\", \"accepts\":[1000,1001,1002]}"
	//fmt.Println(token)
	//rpcInput.Req = []byte(token)
	//bt, _ := pb.Marshal(&rpcInput)
	//proto.Body = bt
	//fmt.Println("1-------------")
	//if err = proto.WriteTCP(wr); err != nil {
	//	log.Error("WriteTCP() error(%v)", err)
	//	return
	//}
}
func InitTCP() {
	conn, err := net.Dial("tcp", Conf.TCPAddr)
	if err != nil {
		log.Error("net.Dial(\"%s\") error(%v)", Conf.TCPAddr, err)
		return
	}
	seqId := int32(2)
	wr := bufio.NewWriter(conn)
	rd := bufio.NewReader(conn)
	// reader
	proto := new(Proto)
	proto.Ver = 1
	// auth
	// test handshake timeout
	// time.Sleep(time.Second * 31)
	proto.Operation = OP_AUTH
	proto.SeqId = seqId
	rpcInput := adapter.RPCInput{}
	var token = "{\"mid\":123, \"room_id\":\"live://1000\", \"platform\":\"web\", \"accepts\":[1000,1001,1002]}"
	fmt.Println(token)
	rpcInput.Req = []byte(token)
	bt, _ := pb.Marshal(&rpcInput)
	proto.Body = bt
	fmt.Println("1-------------")
	if err = proto.WriteTCP(wr); err != nil {
		log.Error("WriteTCP() error(%v)", err)
		return
	}
	fmt.Println("2-------------")

	//if err = proto.ReadTCP(rd); err != nil {
	//	log.Error("tcpReadProto() error(%v)", err)
	//	return
	//}
	//fmt.Println("3-------------")

	log.Infof("auth ok, proto: %v", proto)
	seqId++
	// writer
	go func() {
		fmt.Println("heart beat")

		proto1 := new(Proto)
		for {
			seqId++

			// heartbeat
			proto1.Operation = OP_HEARTBEAT
			proto1.SeqId = seqId
			proto1.Body = nil
			if err = proto1.WriteTCP(wr); err != nil {
				log.Error("WriteTCP() hearbt error(%v)", err)
				return
			}
			// test heartbeat
			//time.Sleep(time.Second * 31)
			seqId++
			// op_test
			//proto1.Operation = OP_TEST
			//proto1.SeqId = seqId
			//if err=proto1.WriteTCP(wr); err != nil {
			//	log.Error("WriteTCP() hearbt test error(%v)", err)
			//	return
			//}
			//seqId++

			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		fmt.Println("recv--")

		if err = proto.ReadTCP(rd); err != nil {
			log.Error("proto.ReadTCP() error(%v)", err)
			return
		}
		log.Errorf("resulto peration %v,seq %v", proto.Operation, proto.SeqId)
		rpcOutput := adapter.RPCOutput{}
		json.Unmarshal(proto.Body, &rpcOutput)
		log.Errorf("result pt%v", proto)
		opert := string(proto.Body[:])
		log.Errorf("result %v", string(proto.Body[:]))
		if opert == "EnterRoom" {

		}

		if proto.Operation == OP_HEARTBEAT_REPLY {
			log.Infof("receive heartbeat")
			if err = conn.SetReadDeadline(time.Now().Add(25 * time.Second)); err != nil {
				log.Error("conn.SetReadDeadline() error(%v)", err)
				return
			}
		} else if proto.Operation == OP_TEST_REPLY {
			log.Infof("body: %s", string(proto.Body))
		} else if proto.Operation == OP_SEND_SMS_REPLY {
			log.Infof("body: %s", string(proto.Body))
		}
	}
}
