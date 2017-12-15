package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func timeHandler2(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC3339)
	w.Write([]byte("The time2 is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	//强制转换到HandlerFunc，再调用mux.Handle
	th := http.HandlerFunc(timeHandler)
	mux.Handle("/time", th)

	//直接用HandleFunc，跳过强制转换
	mux.HandleFunc("/time2", timeHandler2)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
