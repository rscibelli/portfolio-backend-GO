package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorld)

	http.HandleFunc("/trump-quote", GetDumbQuote)

	fmt.Println("Starting server at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!"
	fmt.Fprintf(w, string(msg))
}
