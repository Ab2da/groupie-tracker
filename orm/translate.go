package orm

func BuildArtistCardVM(artistDtm ArtistDTM) ArtistCardVM {
	
}

func BuildAllArtistCardVMs(dtms []ArtistDTM) []ArtistCardVM {
	for _ a := range dtms {
		BuilBuildArtistCardVM(a)
	}
}