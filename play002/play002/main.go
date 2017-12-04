package main

import (
	"errors"
	"log"
	"runtime/debug"
)

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic recover:%v", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			debug.PrintStack()
		}
	}()

	return funcB()
}

func funcB() error {
	panic("foo")
	return errors.New("success")
}

func main() {
	err := funcA()

	if err == nil {
		log.Println("err is nil")
	} else {
		log.Printf("err is %v", err)
	}
}
