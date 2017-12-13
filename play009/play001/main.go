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
	s := Student{Name: "Lilei", Age: 20}

	//object->reflect
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	log.Println("The type is", typ.String())
	log.Println("Student's Name is", val.FieldByName("Name").String())

	//reflect->object
	if s, ok := val.Interface().(Student); ok {
		log.Println("The student's Name is", s.Name)
	} else {
		log.Println("Wrong")
	}
}
