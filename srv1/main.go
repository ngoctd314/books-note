package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET: /")
		time.Sleep(time.Second)
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}
