package main

import (
	dal "dal/DAL"
	"fmt"

	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page! :D!")
	fmt.Println("Endpoint Hit:homePage")
	dal.GetArtistsData()
}

func setupServer() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	setupServer()

}
