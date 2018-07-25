package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//func main() {
//	sigs := make(chan os.Signal, 1)
//	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
//
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			log.Println("send to ch:", i)
//			ch <- i
//			log.Println("send complete:", i)
//		}
//	}()
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			c := <-ch
//			log.Println("get from ch:", c)
//			time.Sleep(time.Second)
//			log.Println("get complete:", c)
//		}
//	}()
//
//	<-sigs
//	log.Println("program end")
//}

//不会报deadlock
//func main() {
//	sigs := make(chan os.Signal, 1)
//	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
//
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			log.Println("send to ch:", i)
//			ch <- i
//			log.Println("send complete:", i)
//		}
//	}()
//
//	go func() {
//		for i := 0; i < 6; i++ {
//			c := <-ch
//			log.Println("get from ch:", c)
//			time.Sleep(time.Second)
//			log.Println("get complete:", c)
//		}
//	}()
//
//	<-sigs
//	log.Println("program end")
//}

//deadlock 和上面这个例子不一样，上面的这个例子是不会报错的，因为deadlock是要all goroutine dead
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			log.Println("send to ch:", i)
//			ch <- i
//			log.Println("send complete:", i)
//		}
//	}()
//
//	for i := 0; i < 6; i++ {
//		c := <-ch
//		log.Println("get from ch:", c)
//		time.Sleep(time.Second)
//		log.Println("get complete:", c)
//	}
//}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			log.Println("send to ch:", i)
			ch <- i
			log.Println("send complete:", i)
		}
		close(ch)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c := <-ch
			log.Println("get from ch:", c)
			time.Sleep(time.Second)
			log.Println("get complete:", c)
		}
	}()

	<-sigs
	log.Println("program end")
}
