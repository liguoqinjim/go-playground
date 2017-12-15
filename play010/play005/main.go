package main

import (
	"log"
	"net/http"
	"time"
)

//第一种方法
func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

//第二种方法
func timeHandler2(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}

//第三种方法
func timeHandler3(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is:" + tm))
	}
}

func main() {
	mux := http.NewServeMux()

	//第一种方法
	th := timeHandler(time.RFC1123)
	mux.Handle("/time", th)

	//第二种方法
	th2 := timeHandler2(time.RFC3339)
	mux.Handle("/time2", th2)

	//第三种方法
	th3 := timeHandler3(time.RFC822)
	mux.HandleFunc("/time3", th3)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
