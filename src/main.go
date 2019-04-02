package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var PORT string
	fs := http.FileServer(http.Dir("static/"))
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi from path: %s\n", r.URL.Path)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello! greeting server reached")
	})
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}
