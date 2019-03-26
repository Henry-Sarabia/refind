package spotify

import (
	"github.com/Henry-Sarabia/scry"
	"github.com/zmb3/spotify"
)

func ParseArtist(old spotify.SimpleArtist) scry.Artist {
	return scry.Artist{
		ID:   string(old.ID),
		Name: old.Name,
	}
}

func ParseArtists(old ...spotify.FullArtist) []scry.Artist {
	var a []scry.Artist

	for _, o := range old {
		a = append(a, ParseArtist(o.SimpleArtist))
	}

	return a
}

func ParseTrack(old spotify.SimpleTrack) scry.Track {
	return scry.Track{
		ID:     string(old.ID),
		Name:   old.Name,
		Artist: ParseArtist(old.Artists[0]),
	}
}

func ParseSimpleTracks(old ...spotify.SimpleTrack) []scry.Track {
	var t []scry.Track

	for _, o := range old {
		t = append(t, ParseTrack(o))
	}

	return t
}

func ParseFullTracks(old ...spotify.FullTrack) []scry.Track {
	var t []scry.Track

	for _, o := range old {
		t = append(t, ParseTrack(o.SimpleTrack))
	}

	return t
}

func ParsePlaylist(old spotify.FullPlaylist) scry.Playlist {
	return scry.Playlist{
		ID:  string(old.ID),
		URI: string(old.URI),
	}
}
