package main

//deadlock
//func main() {
//	ch1 := make(chan int)
//	ch1<-1
//}

//deadlock
//func main() {
//	ch1 := make(chan int)
//	ch1 <- 1
//
//	log.Println(<-ch1)
//}

//deadlock
//func main() {
//	ch1 := make(chan int)
//	ch1 <- 1
//
//	go func() {
//		log.Println(<-ch1)
//	}()
//}

//正确
//func main() {
//	ch1 := make(chan int)
//
//	go func() {
//		ch1 <- 1
//	}()
//
//	log.Println(<-ch1)
//}

//正确
//func main() {
//	ch1 := make(chan int, 1)
//	ch1 <- 1
//	log.Println(<-ch1)
//}

//deadlock，有buffer的channel也不能发超过buffer size的数据，在没有receiver的时候要是超过size也是会deadlock的
//func main() {
//	ch1 := make(chan int, 1)
//	ch1 <- 1
//	ch1 <- 2 //deadlock
//	log.Println(<-ch1)
//	log.Println(<-ch1)
//}

//deadlock 最后一个receiver是收不到消息的
//func main() {
//	ch1 := make(chan int, 1)
//	go func() {
//		ch1 <- 1
//		ch1 <- 2
//	}()
//
//	log.Println(<-ch1)
//	log.Println(<-ch1)
//	log.Println(<-ch1) //deadlock
//}

//deadlock
//func main() {
//	ch1 := make(chan int, 1)
//	go func() {
//		ch1 <- 1
//
//		time.Sleep(time.Second * 10)
//		ch1 <- 2
//	}()
//
//	log.Println(<-ch1)
//	log.Println(<-ch1)
//	log.Println(<-ch1)
//}

//deadlock，但是会打印出return!!!
//func main() {
//	ch1 := make(chan int, 1)
//	go func() {
//		time.Sleep(time.Second * 5)
//		log.Println("return!!!")
//		return
//
//		ch1 <- 2
//	}()
//
//	log.Println(<-ch1)
//}
