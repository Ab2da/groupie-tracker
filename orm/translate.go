package orm

import (
	"deedee/groupie-tracker/dal"
)

func BuildArtistCardVM(a dal.ArtistDTM) ArtistCardVM {

	CardVM := ArtistCardVM{Image: a.Image, Name: a.Name}
	return CardVM
}

func BuildAllArtistCardVMs(dtms []ArtistDTM) []ArtistCardVM {
	for _, a := range dtms {
		BuildArtistCardVM(a)
	}
}
