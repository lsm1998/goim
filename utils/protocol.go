package utils

type ProtocolType int32

const (
	Text = ProtocolType(iota)
	ProtoBuf
	JSON
	Custom
)
