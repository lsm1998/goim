package tests

import (
	"fmt"
	"logic-user/dao"
	"logic-user/model"
	"testing"
)

func TestRoles(t *testing.T) {
	fmt.Println(dao.QueryRoles(1))

	fmt.Println(dao.QueryUser(&model.User{
		Username: "jqr1",
	}))
}
