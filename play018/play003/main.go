package main

import "log"

func main() {
	persons := make(map[string]int)
	persons["小明"] = 19

	mp := &persons

	log.Printf("原始map的内存地址:%p", mp)
	modify(persons)
	log.Println(persons["小明"])
}

func modify(p map[string]int) {
	log.Printf("函数里接收到map的内存地址:%p", &p)
	p["小明"] = 20
}
