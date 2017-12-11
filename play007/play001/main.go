package main

import (
	"log"
	"time"
)

var TestSign = make(chan int, 1)

func main() {
	//开启两个goroutine监听
	go func() {
		<-TestSign
		log.Println("get testsign1")
	}()

	go func() {
		<-TestSign
		log.Println("get testsign2")
	}()

	//发送
	time.Sleep(time.Second)
	TestSign <- 1

	time.Sleep(time.Hour)
}
