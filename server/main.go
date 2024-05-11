package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("hello world"))
	})

	log.Print("Server listening on localhost:2345")

	err := http.ListenAndServe(":2345", nil)
	if err != nil {
		log.Fatal(err)
	}
}
