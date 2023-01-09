package main

import (
	"deedee/groupie-tracker/dal"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var artists []dal.ArtistDTM
var dates []dal.DateDTM

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
	Artists []ArtistCardVM
}

type ArtistVM struct {
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
}

type ArtistCardVM struct {
	Image string
	Name  string
}

func BuildArtistCardVM(a dal.ArtistDTM) ArtistCardVM {
	cardVM := ArtistCardVM{Image: a.Image, Name: a.Name}
	return cardVM
}

func BuildAllArtistCardVMs(dtms []dal.ArtistDTM) []ArtistCardVM {
	var result []ArtistCardVM
	for _, a := range dtms {
		artist := BuildArtistCardVM(a)
		result = append(result, artist)
	}
	return result
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var p PageModel = PageModel{Artists: BuildAllArtistCardVMs(artists)}
	log.Printf("Welcome to the Home Page! :D!")
	fmt.Println("Endpoint Hit:homePage")
	t, err := template.ParseFiles("./wwwroot/MainLayout.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, p)
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

func init() {
	artists = dal.GetArtists()
	dates = dal.GetDates(artists)
}

// Right now, when you run the program, we are just
// getting a list of artists with our DAL (data access layer)
// functions, and printing out the name and members of each

func Double(x int) {
	x *= 2
}

func main() {
	var x int = 5
	Double(x)
	fmt.Println(x)
	// for _, a := range artists {
	// 	fmt.Println(a.Name)
	// 	fmt.Println(a.ConcertDates)
	// 	// fmt.Printf("\nMembers:\n")
	// 	// for _, m := range a.Members {
	// 	// 	fmt.Println(m)
	// 	// }
	// 	// fmt.Println()
	// }

	// for _, d := range dates {
	// 	fmt.Printf("ID: %d Dates:\n", d.ID)
	// 	for _, v := range d.Dates {
	// 		fmt.Println(v)
	// 	}
	// }

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

// func ArtistDTMsToEvent(artists []dal.ArtistDTM) []Event {
// 	var event []Event
// 	for _, v := range artists {
// 		for _, d := range dates {
// 			if d.ID == v.ID {
// 				var gig Event = Event{Location: v.Locations, Dates: d.Dates}
// 				event = append(event, gig)
// 				break
// 			}
// 		}
// 	}
// 	return event
// }
