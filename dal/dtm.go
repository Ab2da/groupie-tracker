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

type DateDTM struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type LocationDTM struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Index struct {
	Index []RelationDTM
}

type RelationDTM struct {
	Index
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
