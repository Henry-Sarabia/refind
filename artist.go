package scry

import "github.com/zmb3/spotify"

type artist struct {
	id   string
	name string
}

func newArtist(a spotify.SimpleArtist) artist {
	return artist{
		id:   string(a.ID),
		name: a.Name,
	}
}
