package main

import "log"

//全局变量
var a int = 20

func main() {
	log.Println("a=", a)

	//局部变量1
	var a = 10
	log.Println("a=", a)

	for i := 0; i < 1; i++ {
		log.Println("a=", a)
		a = 4
		a := 1 //这个大括号中的局部变量a(称局部变量2)
		log.Println("a=", a)
		a = 3 //只赋值给局部变量2
		log.Println("a=", a)
	}
	log.Println("a=", a)
}
