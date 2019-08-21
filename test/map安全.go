package main

import "sync"

type SafeMap struct {
	sync.RWMutex
	Map map[int64]string
}

func NewSafeMap(size int) *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int64]string, size)
	return sm
}

func (sm *SafeMap) ReadMap(key int64) string {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) WriteMap(key int64, value string) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

// 用于for k，_ := range m.Keys(){v := m.ReadMap(k) ....}
func (sm *SafeMap) Keys() []int64 {
	sm.RLock()
	value := make([]int64, 0)
	for k, _ := range sm.Map {
		value = append(value, k)
	}
	sm.RUnlock()
	return value
}
