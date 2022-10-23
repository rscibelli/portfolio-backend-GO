package main

import (
	"fmt"
	"net/http"
)

func GetDumbQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := GetTrumpQuote()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(quote))
}
