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

// Look at the JSON that we receive from a call
// to the /locations endpoint - does it contain
// all the fields below? What do you think each
// field in the response represents?
type LocationDTM struct {
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

// This struct maps perfectly to the JSON response body
// As you edit the LocationDTM above, can you see how
// the ID and dates fields work together with the Artist
// and Location DTMs?
type DateDTM struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Similarly with the /location endpoint, look at the JSON fields
// An instance of this object should store a relationship between a Location
// and one or more Dates (see if you can work out what the
// leading '*' in some of the dates is for!)
type RelationDTM struct {
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
