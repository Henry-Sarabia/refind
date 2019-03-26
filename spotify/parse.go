package spotify

import (
	"github.com/Henry-Sarabia/refind"
	"github.com/zmb3/spotify"
)

func ParseArtist(prev spotify.SimpleArtist) refind.Artist {
	return refind.Artist{
		ID:   string(prev.ID),
		Name: prev.Name,
	}
}

func ParseArtists(prev ...spotify.FullArtist) []refind.Artist {
	var curr []refind.Artist

	for _, p := range prev {
		curr = append(curr, ParseArtist(p.SimpleArtist))
	}

	return curr
}

func ParseTrack(prev spotify.SimpleTrack) refind.Track {
	return refind.Track{
		ID:     string(prev.ID),
		Name:   prev.Name,
		Artist: ParseArtist(prev.Artists[0]),
	}
}

func ParseSimpleTracks(prev ...spotify.SimpleTrack) []refind.Track {
	var curr []refind.Track

	for _, p := range prev {
		curr = append(curr, ParseTrack(p))
	}

	return curr
}

func ParseFullTracks(prev ...spotify.FullTrack) []refind.Track {
	var curr []refind.Track

	for _, p := range prev {
		curr = append(curr, ParseTrack(p.SimpleTrack))
	}

	return curr
}
