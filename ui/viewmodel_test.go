package ui

import (
	"deedee/groupie-tracker/dal"
	"testing"
)

func TestExpectFailureBuildArtistViewModel(t *testing.T) {
	var artistDTM dal.ArtistDTM = dal.ArtistDTM{ID: -1}
	_, err := BuildArtistViewModel(artistDTM)
	if err == nil {
		t.Error("should have failed")
	}
	artistDTM = dal.ArtistDTM{ID: 53}
	_, err = BuildArtistViewModel(artistDTM)
	if err == nil {
		t.Error("should have failed")
	}
}

func TestExpectSuccessBuildArtistViewModel(t *testing.T) {
	var artistDTM dal.ArtistDTM = dal.ArtistDTM{ID: 1}
	_, err := BuildArtistViewModel(artistDTM)
	if err != nil {
		t.Error("should have been successful")
	}
	artistDTM = dal.ArtistDTM{ID: 52}
	_, err = BuildArtistViewModel(artistDTM)
	if err != nil {
		t.Error("should have been successful")
	}
}
