package scry

import "github.com/zmb3/spotify"

type artist struct {
	id   string
	name string
}

func parseArtist(old spotify.SimpleArtist) artist {
	return artist{
		id:   string(old.ID),
		name: old.Name,
	}
}

func parseArtists(old ...spotify.FullArtist) []artist {
	var a []artist

	for _, o := range old {
		a = append(a, parseArtist(o.SimpleArtist))
	}

	return a
}
