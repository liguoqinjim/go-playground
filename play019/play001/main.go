package main

import "log"

func main() {
	//不加这个defer就会出错
	//defer func() {
	//	recover()
	//}()

	do()
}

func do() {
	log.Println(recover())
	panic("error")
}
