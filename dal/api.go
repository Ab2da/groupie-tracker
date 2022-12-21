package dal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Define the Base URL and all endpoints as constants
// Improves readability and reusability in code!
// For example, if we had 1000 functions making a call to the dates endpoint
// of the API, and tomorrow, the owners moved the website to a new domain name
// and renamed the 'dates' endpoint:
// e.g. 	Domain - newgroupietracker.org		Endpoint - /concert-dates
// We could simply update the baseUrl and the relevant endpoints once, and all
// our functions would work without having to re-write the URLs 1000 times
const (
	baseUrl           string = "https://groupietrackers.herokuapp.com/api"
	artistsEndpoint   string = "/artists"
	locationsEndpoint string = "/locations"
	datesEndpoint     string = "/dates"
	relationsEndpoint string = "/relations"
)

// Encapsulate the API call (through http.Get),
// reading the response body,JSON,
// into one function!
// Now you can simply replace the url and the pointer arguments
// with the specific API endpoint and the correct Go DTM
func getData(url string, ptr any) any {
	jsonMessage, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonBody, err := ioutil.ReadAll(jsonMessage.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(jsonBody, ptr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ptr
}

// Here's an example of using 'getData' with the URL
// that corresponds to the /artists endpoint of the
// Groupie-Tracker API. It returns a slice of ArtistDTM
// objects, which can then be taken by the UI (front-end)
// for display purposes
func GetArtists() []ArtistDTM {
	var artists []ArtistDTM
	var url string = baseUrl + artistsEndpoint
	getData(url, &artists)
	return artists
}

func GetDates(artists []ArtistDTM) []DateDTM {
	var dates []DateDTM
	for _, v := range artists {
		var url string = v.ConcertDates
		var d DateDTM
		getData(url, &d)
		dates = append(dates, d)
	}
	return dates
}

func GetLocations(artists []ArtistDTM) []LocationDTM {
	var locations []LocationDTM
	for _, v := range artists {
		var url string = v.Locations
		var l LocationDTM
		getData(url, &l)
		locations = append(locations, l)
	}
}

// func GetRelations() map[string][]string {
// 	var r map[string][]RelationDTM
// 	var url = baseUrl + relationsEndpoint
// 	getData(url, &r)
// 	m := RelationEditor(r["index"][0])
// 	return m
// }

// func RelationEditor(r RelationDTM) map[string][]string {
// 	//make a map and for loops to print out the data
// 	newMap := make(map[string][]string)
// 	for k, v := range r.DatesLocations {
// 		newMap[k] = v
// 	}
// 	return newMap
// }
