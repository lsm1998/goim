package handler

import (
	"utils/chain"
)

type NetWorkDataEventHandler struct {
}

func NewNetWorkDataEventHandler() chain.EventHandler {
	instance := new(NetWorkDataEventHandler)
	instance.Init()
	return instance
}

func (n *NetWorkDataEventHandler) Init() {

}

func (n *NetWorkDataEventHandler) Filter(line chain.EventLine) bool {

	return true
}

func (n *NetWorkDataEventHandler) Handler(line chain.EventLine) bool {
	return true
}

func (n *NetWorkDataEventHandler) Stop() {

}
