package sink

import (
	"fmt"
	"utils/chain"
)

type FileDataSink struct {
}

func newFileDataSink() *FileDataSink {
	instance := &FileDataSink{}
	return instance
}

func (f *FileDataSink) PushLine(line chain.EventLine) error {
	f.DealMsg(line)
	return nil
}

func (f *FileDataSink) Setup() error {
	return nil
}

func (f *FileDataSink) Stop() {
}

// DealMsg 保存上线事件消息
func (f *FileDataSink) DealMsg(line chain.EventLine) {
	fmt.Println("line=", line)
}
