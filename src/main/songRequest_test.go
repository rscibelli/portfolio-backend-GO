package main

import (
	"net/http"
	"os"
	"testing"
)

func TestDelete(t *testing.T) {
	song := "Hey"
	loadTest(t, song)
	r, _ := http.NewRequest("GET", "/", nil)
	q := r.URL.Query()
	q.Add("song", song)
	r.URL.RawQuery = q.Encode()

	DeleteSongRequest(nil, r)
}

func loadTest(t *testing.T, song string) {
	f, err := os.OpenFile(ws, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		t.Error("unable to add file")
	}

	defer f.Close()

	_, err = f.WriteString(song + "--artist--person")
	if err != nil {
		t.Error("unable to add to file")
	}
}
