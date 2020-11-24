package logic

import (
	"github.com/hashicorp/go-uuid"
	"im/config"
	"im/model"
	"im/route"
)

func Login(token string) {

}

// QueryAesKey 查找公钥
func QueryAesKey(uid int64) (string, error) {
	if _, _, key := route.Get(uid); key != "" {
		return key, nil
	}
	var user model.User
	if err := config.DB.Select("aes_key").Where("id=?", uid).First(&user).Error; err != nil {
		return "", err
	}
	return user.AesKey, nil
}

// 获取并保存
func GetAndSaveAesKey(uid int64) (string, error) {
	var aesKey string
	var err error
	if aesKey, err = uuid.GenerateUUID(); err != nil {
		return "", err
	}
	aesKey = aesKey[0:8]
	return aesKey, config.DB.Model(&model.User{}).Where("id=?", uid).Update("aes_key", aesKey).Error
}
