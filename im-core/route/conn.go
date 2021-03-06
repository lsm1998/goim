package route

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"time"
)

// Join 加入一个连接
func Join(id int64, c gnet.Conn, aesKey string) {
	userMap.Store(id, &Connect{Conn: c, PongTime: time.Now().Unix(), AesKey: aesKey})
	fmt.Println("用户加入连接，Id=", id, "，当前在线人数=", Size())
}

// Get 获取一个连接
func Get(id int64) (gnet.Conn, int64, string) {
	val, ok := userMap.Load(id)
	if !ok {
		return nil, 0, ""
	}
	return val.(*Connect).Conn, val.(*Connect).PongTime, val.(*Connect).AesKey
}

// Remove 删除一个连接
func Remove(id int64) {
	userMap.Delete(id)
	fmt.Println("用户断开连接，Id=", id, "，当前在线人数=", Size())
}

// KeyList 在线列表
func KeyList() []int64 {
	list := make([]int64, 0, 10)
	userMap.Range(func(key, value interface{}) bool {
		list = append(list, key.(int64))
		return true
	})
	return list
}

// ConnList 连接列表
func ConnList() []gnet.Conn {
	list := make([]gnet.Conn, 0, 10)
	userMap.Range(func(key, value interface{}) bool {
		list = append(list, value.(*Connect).Conn)
		return true
	})
	return list
}

// ForEach 遍历Map
func ForEach(fu func(id int64, c *Connect) bool) {
	userMap.Range(func(key, value interface{}) bool {
		return fu(key.(int64), value.(*Connect))
	})
}

// Size 在线人数
func Size() int32 {
	var size int32
	ForEach(func(id int64, c *Connect) bool {
		size++
		return true
	})
	return size
}

// SetPongTime 设置PongTime
func SetPongTime(id int64) {
	userMap.Range(func(key, value interface{}) bool {
		if key.(int64) == key {
			temp := value.(*Connect)
			temp.PongTime = time.Now().Unix()
			return false
		}
		return true
	})
}
