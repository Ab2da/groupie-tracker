package ui

import (
	"deedee/groupie-tracker/dal"
	"errors"
	"log"
)

var ArtistViewModels []ArtistViewModel

// Initialise a map that stores each entry with a
// key of the ID of an Artist, and a
// value of the corresponding ArtistViewModel
func InitIDToViewModelMap(dtms []dal.ArtistDTM) {
	// Make the map (in memory)
	ArtistIDToViewModelMap = make(map[int]ArtistViewModel)
	// Range over each ArtistDTM
	for _, artist := range dtms {
		// Build an ArtistViewModel, with relations data,
		// from an ArtistDTM
		viewModel, err := BuildArtistViewModel(artist)
		// If there was a problem with building an ArtistViewModel,
		// log the error and stop the program
		if err != nil {
			log.Fatal(err.Error())
		}
		// Add the individual view model to the slice: ArtistViewModels
		ArtistViewModels = append(ArtistViewModels, viewModel)
		//initalizing our map, setting key:ID, value:ArtistViewModel
		ArtistIDToViewModelMap[artist.ID] = viewModel
	}
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
	CreationDate   int
	DatesLocations map[string][]string
}

// GetNext function is a method of the ArtistViewModel struct
// Therefore, only an ArtistViewModel variable can call this function
// -> GetNext is now a method of the ArtistViewModel type
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

func BuildArtistViewModel(a dal.ArtistDTM) (ArtistViewModel, error) {
	var rel dal.RelationDTM
	var found bool = false
	for _, r := range RelationDTMs {
		if r.ID == a.ID {
			rel = r
			found = true
			break
		}
	}
	if !found {
		return ArtistViewModel{}, errors.New("relation does not exist")
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

	var viewModel ArtistViewModel = ArtistViewModel{ID: a.ID, Image: a.Image, Name: a.Name, FirstAlbum: a.FirstAlbum, Members: a.Members, CreationDate: a.CreationDate, DatesLocations: datesLocations}
	return viewModel, nil
}
