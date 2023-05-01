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

	router.HandleFunc("/song-request", CreateSongRequest).Methods("POST")
	router.HandleFunc("/song-request", GetSongRequests).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
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
