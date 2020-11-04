package dao

import (
	"logic-user/config"
	"logic-user/model"
)

type friendsDetails struct {
	model.User
	GroupId int64 `json:"group_id" gorm:"group_id"`
}

type FriendsDetails struct {
	GroupId int64 `json:"group_id" gorm:"group_id"`
	List    []*model.User
}

func SaveFriends(friends *model.Friends) error {
	return config.DB.Save(friends).Error
}

// QueryFriendsList 查询好友列表
func QueryFriendsList(find *model.Friends) ([]*FriendsDetails, error) {
	query := config.DB.Model((*model.User)(nil)).Select("t_friends.group_id,t_user.*").Joins("inner join t_friends on t_user.id=t_friends.friends_id")
	if find.Id > 0 {
		query = query.Where("t_friends.id=?", find.Id)
	}
	if find.UserId > 0 {
		query = query.Where("t_friends.user_id=?", find.UserId)
	}
	if find.GroupId > 0 {
		query = query.Where("t_friends.group_id=?", find.GroupId)
	}
	var list []*friendsDetails
	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}
	m := make(map[int64][]*model.User, len(list))
	result := make([]*FriendsDetails, 0, len(list))
	for _, v := range list {
		_, ok := m[v.GroupId]
		if ok {
		THE:
			for i := 0; i < len(result); i++ {
				if result[i].GroupId == v.GroupId {
					result[i].List = append(result[i].List, &v.User)
					break THE
				}
			}
		} else {
			m[v.GroupId] = []*model.User{&v.User}
			result = append(result, &FriendsDetails{
				GroupId: v.GroupId,
				List:    m[v.GroupId],
			})
		}
	}
	return result, nil
}
