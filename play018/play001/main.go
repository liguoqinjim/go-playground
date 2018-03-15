package main

import "log"

func main() {
	i := 10
	ip := &i
	log.Println("ip的内存地址:", ip)
	log.Printf("原始指针的内存地址是:%p\n", &ip)
	modify(ip)
	log.Println("i=", i)
}

func modify(ip *int) {
	log.Println("ip的内存地址:", ip)
	log.Printf("传递得到的指针的内存地址是:%p\n", &ip)
	*ip = 9
}
