package tests

import (
	"testing"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func TestBytes(t *testing.T) {
	//req := &message.MessageRequest{Message: &message.Message{Cmd: message.RequestType_PrivateMessage}}
	//req.Message.CreateTime = time.Now().Unix()
	//req.Message.Length = 1
	//req.Message.Body = []byte("hello")
	//
	//Len := unsafe.Sizeof(*req)
	//testBytes := &SliceMock{
	//	addr: uintptr(unsafe.Pointer(req)),
	//	cap:  int(Len),
	//	len:  int(Len),
	//}
	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	//fmt.Println("[]byte is : ", data)
	//
	//var temp = *(**message.MessageRequest)(unsafe.Pointer(&data))
	//fmt.Println(temp.Message.Cmd)
	//fmt.Println(temp.Message.Length)
	//fmt.Println(string(temp.Message.Body))
}
