package main

import (
	"log"
	"math/rand"
	"time"
)

//listing 5(Guarantee of Delivery)
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch

		//Employee performs work here

		//Employee is done and free to go
		log.Println("p=", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper"
}

//listing 6
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		ch <- "paper"

		//Employee is done
	}()

	p := <-ch
	log.Println("p=", p)
}

func main() {
	time.Sleep(time.Hour)
}
