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
	s := &Student{"LiLei", 18}
	c := &Class{"Class A", s, 6, "Century Ave"}

	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	if val.Kind() == reflect.Ptr {
		log.Println("Is is a pointer.Address its value")
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		fv := val.Field(i)
		ft := typ.Field(i)
		switch fv.Kind() {
		case reflect.String:
			log.Printf("The %d th %s types %s valuing %s with tag env %s\n", i, ft.Name, "string", fv.String(), ft.Tag.Get("json"))
		case reflect.Int:
			log.Printf("The %d th %s types %s valuing %d with tag env %s\n", i, ft.Name, "int", fv.Int(), ft.Tag.Get("json"))
		case reflect.Ptr:
			log.Printf("The %d th %s types %s valuing %d with tag env %s\n", i, ft.Name, "pointer", fv.Pointer(), ft.Tag.Get("json"))
		}
	}
}
