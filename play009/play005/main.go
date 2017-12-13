package main

import (
	"reflect"
	"log"
)

func main() {
	m := map[string]string{
		"a": "A",
		"b": "B",
	}

	mv := reflect.ValueOf(m)
	for _,k := range mv.MapKeys(){
		v := mv.MapIndex(k)
		log.Printf("%s - %s\n",k,v)
	}
}
