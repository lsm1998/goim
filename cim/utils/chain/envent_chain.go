package chain

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sort"
	"sync"
	"utils"
)

type createStruct struct {
	create func() EventHandler //创建函数
	sortId int                 //创建序号
}

var create []createStruct

type EventChain struct {
	sink     Sink
	LineChan []chan EventLine
	handler  []EventHandler
	group    sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
}

func (obj *EventChain) SetSink(sink Sink) {
	obj.sink = sink
}

func (obj *EventChain) Index(line EventLine) int {
	return 0
}

func (obj *EventChain) PushLine(line EventLine) error {
	idx := obj.Index(line)
	obj.LineChan[idx] <- line
	return nil
}

func (obj *EventChain) Setup() error {
	obj.ctx, obj.cancel = context.WithCancel(context.Background())

	obj.LineChan = make([]chan EventLine, utils.MediumChanSize)
	for i := 0; i < utils.MediumChanSize; i++ {
		obj.LineChan[i] = make(chan EventLine, 32)
	}

	//排序
	sort.Slice(create, func(i, j int) bool { return create[i].sortId < create[j].sortId })

	var tmp []EventHandler
	for _, f := range create {
		tmp = append(tmp, f.create())
	}
	if len(tmp) > 0 {
		obj.handler = append(obj.handler, tmp...)
	}

	if len(obj.LineChan) != len(obj.handler) {
		panic("EventHandler setup error")
	}
	return nil
}

func (obj *EventChain) Run() {
	goFunc := func(index int) {
		obj.group.Add(1)

	filter:
		for {
			select {
			case <-obj.ctx.Done():
				obj.group.Done()
				break filter
			case line := <-obj.LineChan[index]:
				if obj.handler[index].Handler(line) {
					//各种处理(上线、版本变更)
					if !obj.handler[index].Filter(line) {
						//当返回false时，即被视为数据被过滤
						continue filter
					}
				}
				//下一节点推送
				if err := obj.sink.PushLine(line); err != nil {
					log.Errorf("sink PushLine fail,err=%v \n", err)
				}
			}
		}
	}
	for idx := range obj.LineChan {
		//很多协程
		go goFunc(idx) //避免协程内使用idx
	}
}

func (obj *EventChain) Stop() {
	obj.cancel()
	obj.group.Wait()
	obj.sink.Stop()
}

func (obj *EventChain) UnInit() {

}
