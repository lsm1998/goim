package message

type Group struct {
}

func (*Group) TableName() string {
	return "t_group"
}

type GroupAdmin struct {
}

func (*GroupAdmin) TableName() string {
	return "t_group_admin"
}
