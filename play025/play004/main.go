package main

import "log"

//非缓存类型的channel
//func main() {
//	ch := make(chan int)
//
//	log.Println(len(ch))
//}

//缓存类型的channel，len会返回channel里面现在有多少数据还没有被读取掉
//func main() {
//	ch := make(chan int, 10)
//	ch <- 1
//	ch <- 2
//	log.Println(len(ch))
//
//	<-ch
//	log.Println(len(ch))
//}

//非缓存型的channel len都为0
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 3; i++ {
//			ch <- i
//			log.Println("len1:", len(ch))
//		}
//	}()
//
//	for i := 0; i < 3; i++ {
//		<-ch
//		log.Println("len2:", len(ch))
//		time.Sleep(time.Second)
//	}
//}

//已经关闭的缓存型channel，len还是会返回现在channel里面有多少个数据
//func main() {
//	ch := make(chan int, 10)
//	ch <- 1
//	ch <- 2
//
//	log.Println("ch len:", len(ch))
//
//	close(ch)
//
//	log.Println("ch len:", len(ch))
//}

//已经关闭的非缓存型channel，len还是0
func main() {
	ch := make(chan int)

	log.Println(len(ch))

	close(ch)
	log.Println(len(ch))
}
