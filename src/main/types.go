package main

type SongRequest struct {
	Song   string `json:"song"`
	Artist string `json:"artist"`
	Name   string `json:"name"`
}

type ReturnData struct {
	Message string `json:"message"`
}

type Songs struct {
	Songs []SongRequest `json:"songs"`
}
