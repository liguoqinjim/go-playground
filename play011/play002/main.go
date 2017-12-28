package main

import "log"

func main() {
	//错误的方式
	//for i := 0; i < 3; i++ {
	//	defer func() {
	//		log.Printf("i=%d end", i)
	//	}()
	//
	//	log.Printf("i=%d", i)
	//}

	//解决方法1 (不使用defer)
	//log.Println("解决方法1")
	//for i := 0; i < 3; i++ {
	//	log.Printf("i=%d", i)
	//
	//	log.Printf("i=%d end", i)
	//}

	//解决方法2 (用一个新的func包裹原来的代码)
	log.Println("解决方法2")
	for i := 0; i < 3; i++ {
		func() {
			defer log.Printf("i=%d end", i)
			log.Printf("i=%d", i)
		}()
	}
}
