package main

import (
	"errors"
	"log"
	"runtime/debug"
)

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic recover: %v", p)
			debug.PrintStack()
		}
	}()
	return funcB()
}

func funcB() error {
	panic("foo")
	log.Println("funcB error")
	return errors.New("success")
}

func main() {
	err := funcA()

	if err != nil {
		log.Println("err is nil")
	} else {
		log.Printf("err is %v\n", err)
	}
}
