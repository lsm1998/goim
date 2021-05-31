/*
 * Copyright (c) 2019-03-16.
 * Author: konakona
 * 功能描述:
 */

package adapter

type Message struct {
	data []byte
}

func (m *Message) ProtoMessage() {}

func (m *Message) Reset() {
	*m = Message{}
}

func (m *Message) String() string {
	return string(m.data)
}

func (m *Message) Marshal() ([]byte, error) {
	return m.data, nil
}

func (m *Message) Unmarshal(data []byte) error {
	m.data = data
	return nil
}

func NewMessage(data []byte) *Message {
	return &Message{data}
}
