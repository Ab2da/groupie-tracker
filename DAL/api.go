package dal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUrl           string = "https://groupietrackers.herokuapp.com/api"
	artistsEndpoint   string = "/artists"
	locationsEndpoint string = "/locations"
	datesEndpoint     string = "/dates"
	relationsEndpoint string = "/relations"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var Artists []Artist

type Location struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var Locations []Location

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

var Dates []Date

type Relation struct {
	ID             int `json:"id"`
	DatesLocations struct {
		DunedinNewZealand []string `json:"dunedin-new_zealand"`
		GeorgiaUsa        []string `json:"georgia-usa"`
		LosAngelesUsa     []string `json:"los_angeles-usa"`
		NagoyaJapan       []string `json:"nagoya-japan"`
		NorthCarolinaUsa  []string `json:"north_carolina-usa"`
		OsakaJapan        []string `json:"osaka-japan"`
		PenroseNewZealand []string `json:"penrose-new_zealand"`
		SaitamaJapan      []string `json:"saitama-japan"`
	} `json:"datesLocations"`
}

var Relations []Relation

func GetArtistsData() {
	//Open our json file
	jsonMessage := getData(artistsEndpoint)
	//defer jsonFile.Close()
	//this converts the json file to []bytes
	jsonData, err := ioutil.ReadAll(jsonMessage.Body)
	if err != nil {
		log.Fatal(error(err))
	}

	//var result map[string]interface{}
	//var result []interface{}

	//unmarshall the []byte into Artists []struct
	if err := json.Unmarshal(jsonData, &Artists); err != nil {
		log.Fatal(error(err))
	}
	fmt.Println(Artists)
}

func getLocationsData() {
	//Open our json file
	jsonMessage := getData("/locations")
	//defer jsonMessage.Close()
	//this converts the json file to []bytes
	jsonData, err := ioutil.ReadAll(jsonMessage.Body)
	if err != nil {
		log.Fatal(error(err))
	}

	if err := json.Unmarshal(jsonData, &Locations); err != nil {
		log.Fatal(error(err))
	}
	fmt.Println(Locations)
}

func getDatesData() {
	//Open our json file
	jsonMessage := getData("/dates")

	//this converts the json file to []bytes
	jsonData, err := ioutil.ReadAll(jsonMessage.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(jsonData, &Dates)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(Dates)
}

func getRelationsData() {
	//Open our json file
	jsonMessage, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err.Error())
	}

	//this converts the json file to []bytes
	jsonData, err := ioutil.ReadAll(jsonMessage.Body)

	if err := json.Unmarshal(jsonData, &Relations); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(Relations)

}

func getData(endpoint string) *http.Response {
	jsonMessage, err := http.Get(baseUrl + endpoint)
	if err != nil {
		log.Fatal(err.Error())
	}
	return jsonMessage
}
