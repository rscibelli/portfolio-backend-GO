package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/trump-quote", GetDumbQuote)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
