package main

import "log"

//deadlock 一定要显式的关闭channel，range才会退出
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 3; i++ {
//			ch <- i
//		}
//	}()
//
//	for c := range ch {
//		log.Printf("c=%d", c)
//	}
//}

//range能够不断的读取channel里面的数据，直到该channel被显式的关闭
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 3; i++ {
//			ch <- i
//		}
//		close(ch) //显式关闭
//	}()
//
//	for c := range ch {
//		log.Printf("c=%d", c)
//	}
//}

//缓存型channel
//deadlock
//func main() {
//	ch := make(chan int, 10)
//
//	go func() {
//		for i := 0; i < 3; i++ {
//			ch <- i
//		}
//	}()
//
//	for c := range ch {
//		log.Printf("c=%d", c)
//	}
//}

//缓存型的channel，也是一样，要显式的关闭channel
func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for c := range ch {
		log.Printf("c=%d", c)
	}
}
