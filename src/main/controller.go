package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// http.HandleFunc("/", HelloWorld)

	// http.HandleFunc("/quote", GetDumbQuote)

	router.HandleFunc("/song-request", createSongRequest).Methods("POST")

	router.HandleFunc("/song-request", getSongRequests).Methods("GET")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := "Hello there, World!"
	fmt.Fprintf(w, string(msg))
}

func createSongRequest(w http.ResponseWriter, r *http.Request) {
	CreateSongRequest(w, r)
}

func getSongRequests(w http.ResponseWriter, r *http.Request) {
	GetSongRequests(w, r)
}
