package dao

import (
	"logic-user/config"
	"logic-user/model"
)

func QueryUser(find *model.User) (*model.User, error) {
	query := config.DB.Model(find)
	if find.Id != 0 {
		query = query.Where("id=?", find.Id)
	}
	if find.Username != "" {
		query = query.Where("username=?", find.Username)
	}
	if find.Password != "" {
		query = query.Where("password=?", find.Password)
	}
	var result model.User
	if err := query.First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func SaveUser(role *model.User) error {
	return config.DB.Save(role).Error
}
