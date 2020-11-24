package route

import (
	"im/dao"
	"im/model"
	"log"
)

// init 加载群组信息
func init() {
	Flash()
}

// Flash 刷新群组信息
func Flash() {
	list, err := dao.QueryGroupList()
	if err != nil {
		log.Fatalf("QueryGroupList error,err=%+v \n", err)
		return
	}
	groupMap.Clear()
	for _, v := range list {
		groupMap.Store(v.Group.Id, v.GroupUsers)
	}
}

// JoinGroup 加入群组
func JoinGroup(groupId int64, user *model.User) {
	val, ok := groupMap.Load(groupId)
	if ok {
		list := val.([]*model.User)
		for _, v := range list {
			if v.Id == user.Id {
				return
			}
		}
		list = append(list, user)
	} else {
		groupMap.Store(groupId, []*model.User{user})
	}
}

// LeaveGroup 离开群组
func LeaveGroup(groupId int64, user *model.User) {
	val, ok := groupMap.Load(groupId)
	if ok {
		list := val.([]*model.User)
		for i, v := range list {
			if v.Id == user.Id {
				list = append(list[:i], list[i+1:]...)
				break
			}
		}
	}
}

// GetGroupUsers 获取群组成员
func GetGroupUsers(groupId int64) []*model.User {
	val, ok := groupMap.Load(groupId)
	if ok {
		return val.([]*model.User)
	}
	return nil
}
