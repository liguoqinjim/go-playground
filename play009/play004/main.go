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
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class *Class `json:"class"`
}

func main() {
	s := &Student{}

	val := reflect.ValueOf(s).Elem()
	typ := reflect.TypeOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		fv := val.Field(i)
		ft := typ.Field(i)
		if !fv.CanSet() {
			log.Printf("The %d th %s is unaccessible.\n", i, ft.Name)
			continue
		}

		switch fv.Kind() {
		case reflect.String:
			fv.SetString("Lilei")
		case reflect.Int:
			fv.SetInt(18)
		case reflect.Ptr:
			continue
		}
	}

	log.Println("The student is:", s)
}
