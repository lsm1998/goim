package utils

import (
	"sync"
)

type SyncMap struct {
	sync.Map
}

func (s *SyncMap) Clear() {
	s.Range(func(key, value interface{}) bool {
		s.Delete(key)
		return true
	})
}

func (s *SyncMap) Size() int32 {
	var size int32
	s.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}
