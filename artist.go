package scry

import "github.com/zmb3/spotify"

type Artist struct {
	ID   string
	Name string
}

func ParseArtist(old spotify.SimpleArtist) Artist {
	return Artist{
		ID:   string(old.ID),
		Name: old.Name,
	}
}

func ParseArtists(old ...spotify.FullArtist) []Artist {
	var a []Artist

	for _, o := range old {
		a = append(a, ParseArtist(o.SimpleArtist))
	}

	return a
}
