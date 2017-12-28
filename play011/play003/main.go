package main

import "log"

type database struct {
}

func (db *database) connect() (disconnect func()) {
	log.Println("connnect")

	return func() {
		log.Println("disconnect")
	}
}

func main() {
	//错误的情况
	//db := &database{}
	//defer db.connect()
	//log.Println("query db...")

	//解决方法1
	//db := &database{}
	//disconnect := db.connect()
	//defer disconnect()
	//log.Println("query db...")

	//解决方法2
	db := &database{}
	defer db.connect()()
	log.Println("query db...")
}
