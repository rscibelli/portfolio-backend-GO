package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HelloWorld)

	// router.HandleFunc("/trump-quote", GetDumbQuote)

	router.HandleFunc("/song-request", createSongRequest).Methods("POST")

	router.HandleFunc("/song-request", getSongRequests).Methods("GET")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := "Hello there, World!"
	fmt.Fprintf(w, string(msg))
}

func createSongRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	createSongRequest(w, r)
}

func getSongRequests(w http.ResponseWriter, r *http.Request) {
	GetSongRequests(w, r)
}
