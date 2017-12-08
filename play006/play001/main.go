package main

import (
	"context"
	"log"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Println("ctx.Done()")
					return
				case dst <- n:
					log.Println("n=", n)
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		log.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(time.Second * 10)
}
