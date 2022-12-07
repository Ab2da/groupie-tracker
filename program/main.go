package main

func setupServer() {
	// http.HandleFunc("/", homePage)
}

type Artist struct {
	name       string `json:"name"`
	image      string `json:"image"`
	year       int    `json:"year"`
	firstAlbum int    `json:"firstAlbum"`
	members    string `json:"members"`
}

func main() {
}
