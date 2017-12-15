package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}

	return http.HandlerFunc(fn)
}

func main() {
	th := timeHandler(time.RFC1123)

	http.Handle("/time", th)

	log.Println("Listening...")

	//这里的nil本来我们是要传mux的，但是现在用nil,也就是用http包的DefaultServeMux
	http.ListenAndServe(":3000", nil)
}
