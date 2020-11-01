package tests

import (
	"fmt"
	"logic-user/dao"
	"logic-user/model"
	"testing"
	"time"
)

func TestRoles(t *testing.T) {
	fmt.Println(dao.QueryRoles(1))

	fmt.Println(dao.QueryUser(&model.User{
		Username: "jqr1",
	}))
}

func TestF(t *testing.T) {
	find := model.Friends{}
	find.UserId = 1
	list, _ := dao.QueryFriendsList(&find)
	fmt.Println(list)
}

func TestS(t *testing.T) {
	friends := model.Friends{}
	friends.Id = 4
	friends.GroupId = 1
	friends.UserId = 2
	friends.FriendsId = 1
	friends.CreateTime = time.Now()
	friends.UpdateTime = time.Now()
	fmt.Println(dao.SaveFriends(&friends))
}
