package main

import (
	//why did we do it in this format
	"deedee/groupie-tracker/dal"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var artists []dal.ArtistDTM

type PageModel struct {
	Artists []dal.ArtistDTM
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var p PageModel = PageModel{Artists: dal.GetArtists()}
	log.Printf("Welcome to the Home Page! :D!")
	fmt.Println("Endpoint Hit:homePage")
	t, err := template.ParseFiles("./wwwroot/MainLayout.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, p)
	// dal.GetArtistsData()
}

// func events(w http.ResponseWriter, r *http.Request) {
// 	var p PageModel = PageModel{Artists: artists}
// 	log.Printf("Welcome to the Home Page! :D!")
// 	fmt.Println("Endpoint Hit:homePage")
// 	t, err := template.ParseFiles("./wwwroot/MainLayout.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	t.Execute(w, p)
// 	// dal.GetArtistsData()
// }

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

func init() {
	artists = dal.GetArtists()
}

// Right now, when you run the program, we are just
// getting a list of artists with our DAL (data access layer)
// functions, and printing out the name and members of each
func main() {
	for _, a := range artists {
		fmt.Println(a.Name)
		fmt.Printf("\nMembers:\n")
		for _, m := range a.Members {
			fmt.Println(m)
		}
		fmt.Println()
	}
	setupServer()
}
