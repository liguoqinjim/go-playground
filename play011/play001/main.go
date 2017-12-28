package main

import "log"

func main() {
	var run func() = nil

	defer run()

	log.Println("before run")
}
