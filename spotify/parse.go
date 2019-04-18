package spotify

import (
	"github.com/Henry-Sarabia/refind"
	"github.com/zmb3/spotify"
)

func parseArtist(prev spotify.SimpleArtist) refind.Artist {
	return refind.Artist{
		ID:   string(prev.ID),
		Name: prev.Name,
	}
}

func parseArtists(prev ...spotify.FullArtist) []refind.Artist {
	var curr []refind.Artist

	for _, p := range prev {
		curr = append(curr, parseArtist(p.SimpleArtist))
	}

	return curr
}

func parseTrack(prev spotify.SimpleTrack) refind.Track {
	return refind.Track{
		ID:     string(prev.ID),
		Name:   prev.Name,
		Artist: parseArtist(prev.Artists[0]),
	}
}

func parseSimpleTracks(prev ...spotify.SimpleTrack) []refind.Track {
	var curr []refind.Track

	for _, p := range prev {
		curr = append(curr, parseTrack(p))
	}

	return curr
}