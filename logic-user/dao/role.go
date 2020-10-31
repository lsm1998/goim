package dao

import (
	"logic-user/config"
	"logic-user/model"
)

func QueryRoles(uid int64) ([]*model.Role, error) {
	var list []*model.Role
	if err := config.DB.Model((*model.Role)(nil)).
		Joins("inner join t_auth on t_role.id=t_auth.role_id").
		Select("t_role.*").
		Where("t_auth.user_id=?", uid).
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func SaveRole(role *model.Role) error {
	return config.DB.Save(role).Error
}
