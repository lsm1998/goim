package message

type Message struct {
}

func (*Message) TableName() string {
	return "t_message"
}
