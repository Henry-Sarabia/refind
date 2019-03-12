package scry

import "github.com/zmb3/spotify"

type track struct {
	id     string
	name   string
	artist artist
}

func parseTrack(old spotify.SimpleTrack) track {
	return track{
		id:     string(old.ID),
		name:   old.Name,
		artist: parseArtist(old.Artists[0]),
	}
}

func parseSimpleTracks(old ...spotify.SimpleTrack) []track {
	var t []track

	for _, o := range old {
		t = append(t, parseTrack(o))
	}

	return t
}

func parseFullTracks(old ...spotify.FullTrack) []track {
	var t []track

	for _, o := range old {
		t = append(t, parseTrack(o.SimpleTrack))
	}

	return t
}
