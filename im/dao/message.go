package dao

import (
	"im/config"
	"im/model"
	"utils"
)

type MessageDetails struct {
	model.Message
	// 阅读状态
	Status int32
}

type MessageList struct {
	List  []*MessageDetails
	Total uint32
}

type MessageQuery struct {
	model.Message
	// 1.未读消息 2.已读消息 3.全部消息
	Type int32
}

// QueryMessageList 查询消息列表
func QueryMessageList(find *MessageQuery, page *utils.PageInfo) (*MessageList, error) {
	var total uint32
	var list []*MessageDetails
	query := config.DB.Model((*model.Message)(nil)).Select("t_message.*,t_message_read.status").
		Joins("left join t_message_read on t_message.id=t_message_read.message_id")
	if find.Cmd != 0 {
		query = query.Where("t_message.cmd=?", find.Cmd)
	}
	if find.ToId != 0 {
		query = query.Where("t_message.to_id=?", find.ToId)
	}
	if find.FormId != 0 {
		query = query.Where("t_message.form_id=?", find.FormId)
	}
	if find.Id != 0 {
		query = query.Where("t_message.id=?", find.Id)
	}
	switch {
	case find.Type == 1:
		query = query.Where("t_message_read.status is null")
	case find.Type == 2:
		query = query.Where("t_message_read.status=?", model.MsgUnread)
	case find.Type == 3:
		query = query.Where("t_message_read.status =? or t_message_read.status is null", model.MsgRead)
	}
	if err := query.Offset(page.Offset()).Limit(page.Size).Count(&total).Error; err != nil {
		return nil, err
	}
	if err := query.Offset(page.Offset()).Limit(page.Size).Find(&list).Error; err != nil {
		return nil, err
	}
	return &MessageList{List: list, Total: total}, nil
}
