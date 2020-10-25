package dao

import (
	"im/config"
	"im/model"
	"log"
)

type GroupDetails struct {
	Group      *model.Group
	GroupUsers []*model.User
}

// QueryGroupList 群组列表
func QueryGroupList() ([]*GroupDetails, error) {
	var list []*model.Group
	tx := config.DB.Begin()
	if err := tx.Model((*model.Group)(nil)).Where("delete_time is null").Find(&list).Error; err != nil {
		log.Fatalf("QueryGroupList error,err=%+v \n", err)
		tx.Rollback()
		return nil, err
	}
	result := make([]*GroupDetails, 0, len(list))
	for _, v := range list {
		var item GroupDetails
		item.Group = v
		var itemList []*model.User
		if err := tx.Model((*model.User)(nil)).Select("t_user.*").Joins("inner join t_group_item on t_user.id=t_group_item.user_id").
			Where("t_group_item.group_id=?", v.Id).Find(&itemList).Error; err != nil {
			log.Fatalf("QueryGroupList error,err=%+v \n", err)
			tx.Rollback()
			return nil, err
		}
		item.GroupUsers = itemList
		result = append(result, &item)
	}
	return result, tx.Commit().Error
}
