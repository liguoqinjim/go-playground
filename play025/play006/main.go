package main

//channel的三种状态,nil open close
//func main() {
//	var ch chan string
//
//	if ch == nil {
//		log.Println("ch is nil")
//	}
//
//	//open channel
//	ch = make(chan string)
//
//	//close channel
//	close(ch)
//}

//deadlock
//在nil的channel上收发，会block
//func main() {
//	var ch chan int
//
//	go func() {
//		ch <- 1
//	}()
//
//	log.Println(<-ch)
//}

//在nil的channel上收发，会block
//func main() {
//	sigs := make(chan os.Signal, 1)
//	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
//
//	var ch chan int
//	go func() {
//		log.Println("send start")
//		ch <- 1
//		log.Println("send end")
//	}()
//
//	go func() {
//		log.Println(<-ch)
//	}()
//
//	<-sigs
//	log.Println("program end")
//}

//先在为nil的channel上收发，然后再另外的线程里面open channel是无效的
//func main() {
//	sigs := make(chan os.Signal, 1)
//	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
//
//	var ch chan int
//	go func() {
//		log.Println("send start")
//		ch <- 1
//		log.Println("send end")
//	}()
//
//	go func() {
//		log.Println(<-ch)
//	}()
//
//	time.Sleep(time.Second * 5)
//	ch = make(chan int)
//	log.Println("ch open")
//
//	<-sigs
//	log.Println("program end")
//}

//关闭为nil的channel
//panic: close of nil channel
//func main() {
//	var ch chan int
//
//	close(ch)
//}
