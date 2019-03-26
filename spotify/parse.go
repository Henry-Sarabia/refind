package spotify

import (
	"github.com/Henry-Sarabia/refind"
	"github.com/zmb3/spotify"
)

func ParseArtist(old spotify.SimpleArtist) refind.Artist {
	return refind.Artist{
		ID:   string(old.ID),
		Name: old.Name,
	}
}

func ParseArtists(old ...spotify.FullArtist) []refind.Artist {
	var a []refind.Artist

	for _, o := range old {
		a = append(a, ParseArtist(o.SimpleArtist))
	}

	return a
}

func ParseTrack(old spotify.SimpleTrack) refind.Track {
	return refind.Track{
		ID:     string(old.ID),
		Name:   old.Name,
		Artist: ParseArtist(old.Artists[0]),
	}
}

func ParseSimpleTracks(old ...spotify.SimpleTrack) []refind.Track {
	var t []refind.Track

	for _, o := range old {
		t = append(t, ParseTrack(o))
	}

	return t
}

func ParseFullTracks(old ...spotify.FullTrack) []refind.Track {
	var t []refind.Track

	for _, o := range old {
		t = append(t, ParseTrack(o.SimpleTrack))
	}

	return t
}
