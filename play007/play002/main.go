package main

import (
	"log"
	"time"
)

var TestSign = make(chan int, 1)

func main() {
	go func() {
		<-TestSign
		log.Println("get testsign1")
	}()

	go func() {
		<-TestSign
		log.Println("get testsign2")
	}()

	time.Sleep(time.Second)
	close(TestSign)
	time.Sleep(time.Hour)
}
