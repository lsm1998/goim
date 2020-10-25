package tests

import (
	"fmt"
	"im/config"
	_ "im/config"
	"im/dao"
	"im/model"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	//bean := &model.Message{}
	//
	//if err := config.DB.CreateTable(bean).Error; err != nil {
	//	println(err.Error())
	//}

	list, err := dao.QueryGroupList()
	fmt.Println(list, err)
}

func TestAddTable(t *testing.T) {
	bean := &model.MessageRead{}
	bean.Id = 1
	bean.CreateTime = time.Now()
	bean.UpdateTime = bean.CreateTime
	// bean.DeleteTime = nil
	if err:=config.DB.CreateTable(bean).Error;err!=nil{
		panic(err)
	}
}

func TestAddGroup(t *testing.T) {
	bean := &model.GroupItem{}
	bean.Id = 1
	bean.GroupId=1
	bean.UserId = 1
	bean.CreateTime = time.Now()
	bean.UpdateTime = bean.CreateTime
	// bean.DeleteTime = nil
	if err:=config.DB.Save(bean).Error;err!=nil{
		panic(err)
	}
}