package handler

import (
	"utils"
	"utils/chain"
)

func Register() {
	for i := 0; i < utils.MediumChanSize; i++ {
		registerEventHandler(i)
	}
}

func registerEventHandler(i int) {
	if err := chain.EventRegister(i, func() chain.EventHandler {
		return NewNetWorkDataEventHandler()
	}); err != nil {
		panic(err)
	}
}
