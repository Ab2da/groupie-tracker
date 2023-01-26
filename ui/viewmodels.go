package ui

import (
	"deedee/groupie-tracker/dal"
	"log"
)

var ArtistPathModelMap map[int]ArtistViewModel
var ArtistViewModels []ArtistViewModel

func InitArtistPathModelMap(dtms []dal.ArtistDTM) {
	ArtistPathModelMap = make(map[int]ArtistViewModel)
	for _, artist := range dtms {
		var model ArtistViewModel = BuildArtistViewModel(artist)
		ArtistViewModels = append(ArtistViewModels, model)
		ArtistPathModelMap[artist.ID] = model
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
	for _, r := range RelationDTMs {
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

	var viewModel ArtistViewModel = ArtistViewModel{ID: a.ID, Image: a.Image, Name: a.Name, FirstAlbum: a.FirstAlbum, Members: a.Members, DatesLocations: datesLocations}
	return viewModel
}