package dal

// Define all DTMs (Data Transfer Models)
// These are the Go representation of the various
// JSON objects we receive in the body of the response
// from the Groupie Tracker API

// Artist Data Transfer Model
type ArtistDTM struct {
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

type RelationDTM struct {
	Id             string              `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
