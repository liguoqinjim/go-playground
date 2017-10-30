package main

import (
	"log"
	"math/rand"
	"play999/play001/play001/lib"
	//"time"
)

func main() {
	//rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		log.Println(rand.Intn(1000))
	}

	log.Println("--------------------------------------")
	for i := 0; i < 5; i++ {
		lib.ShowNext()
	}
}
