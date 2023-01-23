package main

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

var artists []dal.ArtistDTM
var relations []dal.RelationDTM
var artistRouteMap map[string]bool
var artistPathModelMap map[int]ArtistViewModel
var artistModels []ArtistViewModel

// var dates []dal.DateDTM
// structs for the display model
type Band struct {
	Image string
	Name  string
}

type Event struct {
	Location string
	Dates    []string
}

type HomeViewModel struct {
	Artists []ArtistViewModel
}

type ArtistViewModel struct {
	ID             int
	Image          string
	Name           string
	FirstAlbum     string
	Members        []string
	DatesLocations map[string][]string
}

func (a ArtistViewModel) GetNext() int {
	var result int = (a.ID + 1) % 52
	if result == 0 {
		result = 52
	}
	return result
}

func (a ArtistViewModel) GetPrev() int {
	var result int = (a.ID - 1) % 52
	if result == 0 {
		result = 52
	}
	return result
}

func BuildArtistViewModel(a dal.ArtistDTM) ArtistViewModel {
	var rel dal.RelationDTM
	var found bool = false
	for _, r := range relations {
		if r.ID == a.ID {
			rel = r
			found = true
			break
		}
	}
	if !found {
		log.Fatal("artist relation not found")
	}
	var datesLocations map[string][]string = make(map[string][]string)
	for key, value := range rel.DatesLocations {
		var loc string = key
		var runes []rune = []rune(loc)
		var titleRunes []rune
		for i, r := range runes {
			if i == 0 {
				titleRunes = append(titleRunes, r-32)
			} else if r == '-' && i < len(runes)-1 {
				titleRunes = append(titleRunes, ',')
				titleRunes = append(titleRunes, ' ')
				runes[i+1] -= 32
			} else if r == '_' && i < len(runes)-1 {
				titleRunes = append(titleRunes, ' ')
				runes[i+1] -= 32
			} else {
				titleRunes = append(titleRunes, r)
			}
		}
		datesLocations[string(titleRunes)] = value
	}

	cardVM := ArtistViewModel{ID: a.ID, Image: a.Image, Name: a.Name, FirstAlbum: a.FirstAlbum, Members: a.Members, DatesLocations: datesLocations}
	return cardVM
}

func InitArtistPathModelMap(dtms []dal.ArtistDTM) {
	artistPathModelMap = make(map[int]ArtistViewModel)
	for _, artist := range dtms {
		var model ArtistViewModel = BuildArtistViewModel(artist)
		artistModels = append(artistModels, model)
		artistPathModelMap[artist.ID] = model
	}
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
		var p HomeViewModel = HomeViewModel{Artists: artistModels}
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
	if artistRouteMap[path] {
		var model ArtistViewModel = artistPathModelMap[id]
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

func setupServer() {
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

func init() {
	artists = dal.GetArtists()
	relations = dal.GetRelations(artists)
	artistRouteMap = make(map[string]bool)
	for _, v := range artists {
		artistRouteMap[fmt.Sprintf("%d", v.ID)] = true
	}
	InitArtistPathModelMap(artists)
}

func main() {
	setupServer()
}

// bridge between the DTM and display model (the structs are a model for your eventual display)
func ArtistDTMsToBands(artists []dal.ArtistDTM) []Band {
	var bands []Band
	for _, artist := range artists {
		var b Band = Band{Name: artist.Name, Image: artist.Image}
		bands = append(bands, b)
	}
	return bands
}
