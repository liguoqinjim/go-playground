package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func main() {

}

//错误的方法
func FindDigits(filename string) []byte {
	b, err := ioutil.ReadFile("1024")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	return digitRegexp.Find(b)
}

//正确的方法
func FindCopyDigits(filename string) []byte {
	b, err := ioutil.ReadFile("1024")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
