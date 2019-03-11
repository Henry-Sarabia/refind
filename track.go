package scry

import "github.com/zmb3/spotify"

type track struct {
	id     string
	name   string
	artist artist
}

func newTrack(t spotify.SimpleTrack) track {
	return track{
		id:     string(t.ID),
		name:   t.Name,
		artist: newArtist(t.Artists[0]),
	}
}
