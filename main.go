package main

import (
	"deedee/groupie-tracker/dal"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const port string = ":8080"

var artists []dal.ArtistDTM
var relations []dal.RelationDTM
var artistRouteMap map[string]bool
var artistPathModelMap map[string]ArtistVM
var artistModels []ArtistVM

// var dates []dal.DateDTM
// structs for the display model
type Band struct {
	Image string
	Name  string
}

type Info struct {
	Members      []string
	CreationDate int
	FirstAlbum   string
}

type Event struct {
	Location string
	Dates    []string
}

type HomeModel struct {
	Artists []ArtistVM
}

type ArtistVM struct {
	Image          string
	Name           string
	FirstAlbum     string
	Members        []string
	DatesLocations map[string][]string
}

func BuildArtistVM(a dal.ArtistDTM) ArtistVM {
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
	var datesLocations map[string][]string = rel.DatesLocations
	cardVM := ArtistVM{Image: a.Image, Name: a.Name, FirstAlbum: a.FirstAlbum, Members: a.Members, DatesLocations: datesLocations}
	return cardVM
}

func InitArtistPathModelMap(dtms []dal.ArtistDTM) {
	artistPathModelMap = make(map[string]ArtistVM)
	for _, artist := range dtms {
		var model ArtistVM = BuildArtistVM(artist)
		artistModels = append(artistModels, model)
		artistPathModelMap[artist.Name] = model
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Verify HTTP method
	if r.Method != http.MethodGet {
		log.Printf("%s - %s - %d\n", r.Method, r.URL.Path, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This site does not support non-GET HTTP requests.\n"))
		return
	}
	// Basic Routing
	if r.URL.Path == "/" {
		var p HomeModel = HomeModel{Artists: artistModels}
		log.Printf("%s - %s - %d\n", r.Method, r.URL.Path, http.StatusOK)
		t, err := template.ParseFiles("./wwwroot/MainLayout.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, p)
		return
	}
	// Routing - Artist Name
	var name string = strings.TrimPrefix(r.URL.Path, "/")
	if artistRouteMap[name] {
		var model ArtistVM = artistPathModelMap[name]
		t, err := template.ParseFiles("./wwwroot/artists.html")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("%s - %s - %d\n", r.Method, r.URL.Path, http.StatusOK)
		return
	}
	// Print a 404 not found if page does not exist
	log.Printf("%s - %s - %d\n", r.Method, r.URL.Path, http.StatusNotFound)
	http.NotFound(w, r)
}

func setupServer() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("wwwroot"))
	mux.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", fs))
	mux.HandleFunc("/", homePage)
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
		artistRouteMap[v.Name] = true
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

func ArtistDTMsToInfo(artists []dal.ArtistDTM) []Info {
	var info []Info
	for _, v := range artists {
		var i Info = Info{Members: v.Members, CreationDate: v.CreationDate, FirstAlbum: v.FirstAlbum}
		info = append(info, i)
	}
	return info
}

func ArtistDTMToInfo(a dal.ArtistDTM) Info {
	return Info{Members: a.Members, CreationDate: a.CreationDate, FirstAlbum: a.FirstAlbum}
}
