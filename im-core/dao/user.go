package dao

import (
	"im/config"
	"im/model"
	"utils"
)

type UserList struct {
	Total uint32
	List  []*model.User
}

func QueryUserList(idList []int64, page *utils.PageInfo) (*UserList, error) {
	result := &UserList{}
	var userList []*model.User
	if err := config.DB.Model((*model.User)(nil)).Where("id in(?)", idList).
		Offset(page.Page * page.Size).Limit(page.Size).Find(&userList).Error; err != nil {
		return nil, err
	}
	result.List = userList
	result.Total = uint32(len(idList))
	return result, nil
}
