package chain

import (
	"context"
	log "github.com/sirupsen/logrus"
)

var sinks []Sink

type SinkChain struct {
	line   chan EventLine
	ctx    context.Context
	cancel context.CancelFunc
}

func (obj *SinkChain) PushLine(line EventLine) error {
	obj.line <- line
	return nil
}

func (obj *SinkChain) Setup() (err error) {
	for _, sink := range sinks {
		if err = sink.Setup(); err != nil {
			return
		}
	}
	obj.line = make(chan EventLine, 32)
	obj.ctx, obj.cancel = context.WithCancel(context.Background())
	return
}

func (obj *SinkChain) Run() {
Loop:
	for {
		select {
		case <-obj.ctx.Done():
			break Loop
		case line := <-obj.line:
			for _, s := range sinks {
				if err := s.PushLine(line); err != nil {
					log.Errorf("sink PushLine fail,err=%v \n", err)
				}
			}
		}
	}
}
func (obj *SinkChain) Stop() {
	obj.cancel()
}

func (obj *SinkChain) UnInit() {

}
