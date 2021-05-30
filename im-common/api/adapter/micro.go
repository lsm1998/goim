/*
 * Copyright (c) 2019-03-16.
 * Author: konakona
 * 功能描述:
 */

package adapter

import (
	pb "github.com/golang/protobuf/proto"
)

const _ = pb.ProtoPackageIsVersion2

//
//// RPCInput has all necessary information when calling downstream service (Now it is Micro)
//type RPCInput struct {
//	ServerName string          `protobuf:"bytes,1,opt,name=obj" json:"obj,omitempty"`
//	Func       string          `protobuf:"bytes,2,opt,name=func" json:"func,omitempty"`
//	Req        json.RawMessage `protobuf:"bytes,3,opt,name=req,proto3" json:"req,omitempty"`
//	Opt        map[int32]string  `protobuf:"bytes,4,rep,name=opt" json:"opt,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
//}
//
//
//func (m *RPCInput) Reset()         { *m = RPCInput{} }
//func (m *RPCInput) String() string { return pb.CompactTextString(m) }
//func (*RPCInput) ProtoMessage()    {}
//
//// RPCOutput is what the downstream service returns
//type RPCOutput struct {
//	Ret        int32           `protobuf:"zigzag32,1,opt,name=ret" json:"ret,omitempty"`
//	Rsp        json.RawMessage `protobuf:"bytes,2,opt,name=rsp,proto3" json:"rsp,omitempty"`
//	Opt        map[int32]string  `protobuf:"bytes,3,rep,name=opt" json:"opt,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
//	Desc       string          `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
//	ServerName string          `protobuf:"bytes,5,opt,name=obj" json:"obj,omitempty"`
//	Func       string          `protobuf:"bytes,6,opt,name=func" json:"func,omitempty"`
//}
//
//func (m *RPCOutput) Reset()         { *m = RPCOutput{} }
//func (m *RPCOutput) String() string { return pb.CompactTextString(m) }
//func (*RPCOutput) ProtoMessage()    {}
