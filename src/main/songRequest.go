package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	ws = "weddingSongs"
	// ws = "songTest"
)

func CreateSongRequest(w http.ResponseWriter, r *http.Request) {
	var retMsg ReturnData

	reqBody, _ := ioutil.ReadAll(r.Body)
	var songRequest SongRequest
	json.Unmarshal(reqBody, &songRequest)

	songRequestString := songRequest.Song + "--" + songRequest.Artist + "--" + songRequest.Name + "\n"

	f, err := os.OpenFile(ws, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		retMsg = ReturnData{Message: "Unable to add song, error opening file"}
		json.NewEncoder(w).Encode(retMsg)
		return
	}

	defer f.Close()

	_, err = f.WriteString(songRequestString)
	if err != nil {
		retMsg = ReturnData{Message: "Unable to add song, error writing to file"}
		json.NewEncoder(w).Encode(retMsg)
		return
	}

	retMsg = ReturnData{Message: "Song has been requested!"}
	json.NewEncoder(w).Encode(retMsg)
}

func GetSongRequests(w http.ResponseWriter, r *http.Request) {
	var s []SongRequest
	file, err := os.Open(ws)
	if err != nil {
		json.NewEncoder(w).Encode("error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		songArr := strings.Split(scanner.Text(), "--")
		song := SongRequest{Song: songArr[0], Artist: songArr[1], Name: songArr[2]}
		s = append(s, song)
	}

	songs := Songs{s}
	json.NewEncoder(w).Encode(songs)
}
