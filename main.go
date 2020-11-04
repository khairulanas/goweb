package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	log.Println("server running : 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello golang"))
}
