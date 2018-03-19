package main

import (
	"log"
	"sync"
)

type UserInfo struct {
	Name string
	Age  int
}

var m sync.Map

func main() {
	vv, ok := m.LoadOrStore(1, "one")
	log.Println(vv, ok)

	vv, ok = m.Load(1)
	log.Println(vv, ok)

	vv, ok = m.LoadOrStore(1, "oneone")
	log.Println(vv, ok)

	vv, ok = m.Load(1)
	log.Println(vv, ok)

	m.Store(1, "oneone")
	vv, ok = m.Load(1)
	log.Println(vv, ok)

	m.Store(2, "twotwo")
	m.Range(func(k, v interface{}) bool {
		log.Println("range1:", k, v)
		return true
	})

	m.Delete(1)
	m.Range(func(k, v interface{}) bool {
		log.Println("range2:", k, v)
		return true
	})
}
