package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// 秘钥
	secretKey = "sbl_lw_lw_lw"
	// 2个小时过期
	expireTime = 2 * time.Hour
)

type LoginClaims struct {
	Uid        int64
	Role       []int64
	ExpireTime int64
}

func (l LoginClaims) Valid() error {
	if time.Now().Unix() > l.ExpireTime {
		return errors.New("token过期")
	}
	return nil
}

func GenerateToken(uid int64, role []int64) (string, error) {
	expire := time.Now().Add(expireTime)
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		Uid:        uid,
		Role:       role,
		ExpireTime: expire.Unix(),
	})
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(secretKey))
}

func ValidToken(token string, uid int64) (*LoginClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(token, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenStr.Claims.(*LoginClaims); ok && tokenStr.Valid && claims.Uid != uid {
		return nil, errors.New("token校验失败")
	} else {
		return claims, nil
	}
}
