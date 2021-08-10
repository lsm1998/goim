package tests

import (
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	redisClient := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "119.91.113.111:6379")
			if err != nil {
				return nil, err
			}
			if _, err = c.Do("auth", "redisyyds123"); err != nil {
				return nil, err
			}
			_, err = c.Do("select", 0)
			return c, err
		},
	}

	conn := redisClient.Get()
	key := "demo-set"
	_, err := conn.Do("sadd", key, 1)
	assert.Nil(t, err)
}
