package tests

import (
	"logic-user/config"
	"logic-user/model"
	"testing"
)

func TestTable(t *testing.T) {
	config.DB.CreateTable(&model.Auth{})

	config.DB.CreateTable(&model.Role{})

	config.DB.CreateTable(&model.Permission{})
}
