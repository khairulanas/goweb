package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)

	log.Println("server running : 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("wellcome to home"))
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello golang"))
}
