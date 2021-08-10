package config

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var RedisClient *redis.Pool

func initRedis() {
	RedisClient = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", C.Redis.Adder)
			if err != nil {
				return nil, err
			}
			if C.Redis.Auth != "" {
				if _, err = c.Do("auth", C.Redis.Auth); err != nil {
					return nil, err
				}
			}
			_, err = c.Do("select", C.Redis.Db)
			return c, err
		},
	}
}

func GetRedis() redis.Conn {
	return RedisClient.Get()
}
