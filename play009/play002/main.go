package main

import (
	"log"
	"reflect"
)

type Class struct {
	Name    string   `json:"name"`
	Student *Student `json:"student"`
	Grade   int      `json:"grade"`
	school  string   `json:"school"`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	t := reflect.TypeOf(Student{})
	val := reflect.New(t)
	log.Println(val.Type().String())
}
