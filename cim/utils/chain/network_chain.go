package chain

import (
	"context"
	"github.com/panjf2000/gnet"
	log "github.com/sirupsen/logrus"
)

var netWorks func() NetWork

type UDPManager struct {
	transform Transforms
	con       NetWork //连接
	ctx       context.Context
	cancel    context.CancelFunc
}

func (obj *UDPManager) Setup() error {
	obj.ctx, obj.cancel = context.WithCancel(context.Background())

	obj.con = netWorks()
	if err := obj.con.Setup(func(data *[]byte, conn gnet.Conn) error {
		if err := obj.transform.Line(data, conn); err != nil {
			log.Errorf("transform handler line fail,err=%v \n", err)
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	go obj.con.Run()
	return nil
}

func (obj *UDPManager) SetSink(transform Transforms) {
	obj.transform = transform
}

func (obj *UDPManager) Run() {
	// 处理信息接口 可以开n个协程处理n个连接
The:
	for {
		select {
		case <-obj.ctx.Done():
			break The
		}
	}
	//通知退出
	obj.transform.Stop()
}

func (obj *UDPManager) Stop() {
	obj.cancel()
}

func (obj *UDPManager) UnInit() {
	obj.con.Stop()
}
