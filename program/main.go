package main

import (
	"deedee/groupie-tracker/dal"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Welcome to the Home Page! :D!")
	fmt.Println("Endpoint Hit:homePage")
	t, err := template.ParseFiles("./wwwroot/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, "hi")
	// dal.GetArtistsData()
}

func setupServer() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("wwwroot"))
	mux.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", fs))
	mux.HandleFunc("/", homePage)
	err := http.ListenAndServe(":9999", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Right now, when you run the program, we are just
// getting a list of artists with our DAL (data access layer)
// functions, and printing out the name and members of each
func main() {
	// 	setupServer()
	var artists []dal.ArtistDTM = dal.GetArtists()
	for _, a := range artists {
		fmt.Println(a.Name)
		fmt.Printf("\nMembers:\n")
		for _, m := range a.Members {
			fmt.Println(m)
		}
		fmt.Println()
	}
}
