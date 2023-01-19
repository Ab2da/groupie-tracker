package main

import (
	"deedee/groupie-tracker/dal"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var artists []dal.ArtistDTM
var relations []dal.RelationDTM

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

type PageModel struct {
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

func BuildAllArtistVMs(dtms []dal.ArtistDTM) []ArtistVM {
	var result []ArtistVM
	for _, a := range dtms {
		artist := BuildArtistVM(a)
		result = append(result, artist)
	}
	return result
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		var p PageModel = PageModel{Artists: BuildAllArtistVMs(artists)}
		log.Printf("Welcome to the Home Page! :D!")
		fmt.Println("Endpoint Hit:homePage")
		t, err := template.ParseFiles("./wwwroot/MainLayout.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, p)
		return
	}

	for _, v := range artists {
		if r.URL.Path == "/"+v.Name {
			var info ArtistVM = BuildArtistVM(v)
			t, err := template.ParseFiles("./wwwroot/artists.html")
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			t.Execute(w, info)
			return
		}
	}
	http.NotFound(w, r)
}

// Add defer close
func setupServer() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("wwwroot"))
	mux.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", fs))
	mux.HandleFunc("/", homePage)
	err := http.ListenAndServe(":9876", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	artists = dal.GetArtists()
	relations = dal.GetRelations(artists)
}

func main() {
	r := dal.GetRelations(artists)
	for _, v := range r {
		fmt.Println(v)
	}

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
