package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET: /")
		d := rand.Intn(3) + 1
		time.Sleep(time.Second * time.Duration(d))
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8081", nil)
}
