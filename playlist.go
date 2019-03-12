package scry

import "github.com/zmb3/spotify"

type Playlist struct {
	ID  string
	URI string
}

func ParsePlaylist(old spotify.FullPlaylist) Playlist {
	return Playlist{
		ID:  string(old.ID),
		URI: string(old.URI),
	}
}
