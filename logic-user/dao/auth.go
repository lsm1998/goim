package dao

import (
	"logic-user/config"
	"logic-user/model"
)

func SaveAuth(auth *model.Auth) error {
	return config.DB.Save(auth).Error
}
