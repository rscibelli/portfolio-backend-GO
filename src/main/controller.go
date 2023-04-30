package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// headersOk := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin"})
	// originsOk := handlers.AllowedOrigins([]string{"http://10.0.0.251:3000"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	// http.HandleFunc("/", HelloWorld)

	// http.HandleFunc("/quote", GetDumbQuote)

	router.HandleFunc("/song-request", createSongRequest).Methods("POST")
	router.HandleFunc("/song-request", getSongRequests).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://10.0.0.251:3000", "http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
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
