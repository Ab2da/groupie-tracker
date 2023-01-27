package ui

import (
	"deedee/groupie-tracker/dal"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const port string = ":8080"

var ArtistDTMs []dal.ArtistDTM
var RelationDTMs []dal.RelationDTM
var ArtistRouteMap map[string]bool
var ArtistPathModelMap map[int]ArtistViewModel

func init() {
	ArtistDTMs = dal.GetArtistDTMs()
	RelationDTMs = dal.GetRelationDTMs(ArtistDTMs)
	ArtistRouteMap = make(map[string]bool)
	for _, v := range ArtistDTMs {
		//every other ID is set to false  i.d = 1000: false
		ArtistRouteMap[fmt.Sprintf("%d", v.ID)] = true
	}
	InitArtistPathModelMap(ArtistDTMs)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Verify HTTP method
	if r.Method != http.MethodGet {
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This site does not support non-GET HTTP requests.\n"))
		return
	}
	// Basic Routing
	if r.URL.Path == "/" {
		var p HomeViewModel = HomeViewModel{Artists: ArtistViewModels}
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusOK, http.StatusText(http.StatusOK))
		t, err := template.ParseFiles("./wwwroot/MainLayout.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
			return
		}
		t.Execute(w, p)
		return
	}
	// Routing - Artist ID
	var path string = strings.TrimPrefix(r.URL.Path, "/")
	var id int
	var err error
	id, err = strconv.Atoi(path)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
		return
	}

	if ArtistRouteMap[path] {
		var model ArtistViewModel = ArtistPathModelMap[id]
		t, err := template.ParseFiles("./wwwroot/artists.html")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
			return
		}
		err = t.Execute(w, model)
		if err != nil {
			log.Println(err.Error())
			// w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
		}
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusOK, http.StatusText(http.StatusOK))
		return
	}
	// Print a 404 not found if page does not exist
	log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	http.NotFound(w, r)
}

func RunServer() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("wwwroot"))
	mux.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", fs))
	mux.HandleFunc("/", defaultHandler)
	log.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
