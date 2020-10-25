package route

import (
	"fmt"
	"github.com/panjf2000/gnet"
)

// Join 加入一个连接
func Join(id int64, c gnet.Conn) {
	userMap.Store(id, c)
	fmt.Println("用户加入连接，Id=", id, "，当前在线人数=", Size())
}

// Get 获取一个连接
func Get(id int64) gnet.Conn {
	val, ok := userMap.Load(id)
	if !ok {
		return nil
	}
	return val.(gnet.Conn)
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
		list = append(list, value.(gnet.Conn))
		return true
	})
	return list
}

// ForEach 遍历Map
func ForEach(fu func(id int64, c gnet.Conn) bool) {
	userMap.Range(func(key, value interface{}) bool {
		return fu(key.(int64), value.(gnet.Conn))
	})
}

// Size 在线人数
func Size() int32 {
	var size int32
	ForEach(func(id int64, c gnet.Conn) bool {
		size++
		return true
	})
	return size
}
