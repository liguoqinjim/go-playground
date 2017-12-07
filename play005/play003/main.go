package main

import "log"

func waitForTasks() {
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			log.Println("employee : working", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
	}

	close(ch)
}

func main() {
	waitForTasks()
}
