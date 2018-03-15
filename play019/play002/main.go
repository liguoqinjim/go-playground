package main

import (
	"log"
	"net/http"
)

func main() {
	//res, err := http.Get("http://notexists")
	//defer res.Body.Close()
	//if err != nil {
	//	log.Println(err)
	//}

	res, err := http.Get("http://notexists")
	if err == nil {
		defer res.Body.Close()
	} else {
		log.Println("err=", err)
	}
}
