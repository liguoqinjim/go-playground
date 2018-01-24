package main

import (
	"log"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		num := rand.Intn(1000)
		log.Println(num)
	}

	log.Println("-----------------------------------------------------")

	//设置随机数因子为887
	rand.Seed(59)
	for i := 0; i < 5; i++ {
		num := rand.Intn(1000)
		log.Println(num)
	}
}
