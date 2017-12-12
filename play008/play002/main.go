package main

import "sync"

type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func main() {
	safeMap := newSafeMap()

	for i := 0; i < 1000; i++ {
		go safeMap.setMap(i, i)
		go safeMap.readMap(i)
	}
}

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm
}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()

	return value
}

func (sm *SafeMap) setMap(key, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}
