package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage :D!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Artist struct {
	name       string `json:"name"`
	image      string `json:"image"`
	year       int    `json"year"`
	firstAlbum int    `json"firstAlbum`
	members    string `json:"members"`
}

func main() {
	handleRequests()
}
