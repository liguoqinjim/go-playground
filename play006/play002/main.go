package main

import (
	"context"
	"log"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		log.Println("overslept")
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
