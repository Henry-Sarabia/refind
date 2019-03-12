package scry

import "github.com/zmb3/spotify"

type Track struct {
	ID     string
	Name   string
	Artist Artist
}

func ParseTrack(old spotify.SimpleTrack) Track {
	return Track{
		ID:     string(old.ID),
		Name:   old.Name,
		Artist: ParseArtist(old.Artists[0]),
	}
}

func ParseSimpleTracks(old ...spotify.SimpleTrack) []Track {
	var t []Track

	for _, o := range old {
		t = append(t, ParseTrack(o))
	}

	return t
}

func ParseFullTracks(old ...spotify.FullTrack) []Track {
	var t []Track

	for _, o := range old {
		t = append(t, ParseTrack(o.SimpleTrack))
	}

	return t
}
