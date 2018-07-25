package main

//不会报deadlock，会得到0
//func main() {
//	ch := make(chan int)
//
//	close(ch)
//
//	log.Println(<-ch)
//	log.Println(<-ch)
//}

//报错 panic: send on closed channel
//这个错，如果正常写代码是不会出现的。因为我们应该要在sender的一侧去关闭channel，且关闭channel是要显式调用channel的，那么我们不会不知道这个channel是否已经关闭
//https://stackoverflow.com/questions/39213230/in-golang-how-to-test-if-a-channel-is-close-and-only-send-to-it-when-its-not-c?noredirect=1&lq=1
//func main() {
//	ch := make(chan int)
//	close(ch)
//
//	ch <- 1
//}

//open这个值会得到false当channel被关闭的时候
//func main() {
//	ch := make(chan int)
//	close(ch)
//
//	c, open := <-ch
//	log.Printf("c:%d,open:%t", c, open)
//}

//在缓存型channel里面，关闭了之后不会里面得到false的，要等channel里面的数据都空了之后才会返回false
//func main() {
//	ch := make(chan int, 3)
//	ch <- 1
//	ch <- 2
//
//	close(ch)
//
//	c, open := <-ch
//	log.Printf("c:%d,open:%t", c, open)
//	c, open = <-ch
//	log.Printf("c:%d,open:%t", c, open)
//	c, open = <-ch
//	log.Printf("c:%d,open:%t", c, open)
//}

//range和close一起
func main() {

}
