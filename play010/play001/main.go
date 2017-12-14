package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//RedirectHandler
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
