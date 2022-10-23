package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TrumpResponse struct {
	Appeared_at string   `json:"appeared_at"`
	Created_at  string   `json:"created_at"`
	Quote_id    string   `json:"quote_id"`
	Tags        []string `json:"tags"`
	Updated_at  string   `json:"updated_at"`
	Value       string   `json:"value"`
}

const rapidAPIAPIKey = "8ad99fb4efmshe7df3ad722e4d77p122f8djsn97a24eea665e"

func GetTrumpQuote() (string, error) {

	url := "https://matchilling-tronald-dump-v1.p.rapidapi.com/random/quote"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/hal+json")
	req.Header.Add("X-RapidAPI-Key", rapidAPIAPIKey)
	req.Header.Add("X-RapidAPI-Host", "matchilling-tronald-dump-v1.p.rapidapi.com")

	client := http.Client{Timeout: 5 * time.Second}

	res, _ := client.Do(req)

	// defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	trumpResponse := TrumpResponse{}
	err := json.Unmarshal(body, &trumpResponse)

	if err != nil {
		return "", fmt.Errorf("error unmarshalling stupid quote: " + err.Error())
	}

	fmt.Println(string(body))

	return trumpResponse.Value, nil
}
