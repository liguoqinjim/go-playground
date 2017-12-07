package main

import (
	"log"
	"math/rand"
	"time"
)

//listing 7
func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

			ch <- "paper"
		}()
	}

	for emps > 0 {
		p := <-ch
		log.Println("p=", p)
		emps--
	}
}

//listing 8
func selectDrop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			log.Println("employee : received", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			log.Println("manager : send ack")
		default:
			log.Println("manager : drop")
		}
	}

	close(ch)
}

func main() {
	selectDrop()

	time.Sleep(time.Hour)
}
