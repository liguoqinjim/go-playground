package main

import "log"

func main() {
	//错误的情况
	//{
	//	defer func() {
	//		log.Println("block: defer runs")
	//	}()
	//
	//	log.Println("block: ends")
	//}
	//log.Println("main: ends")

	//解决方法
	func() {
		defer func() {
			log.Println("func: defer runs")
		}()
		log.Println("func ends")
	}()
	log.Println("main ends")
}
