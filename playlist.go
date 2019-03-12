package scry

import "github.com/zmb3/spotify"

type playlist struct {
	id  string
	uri string
}

func parsePlaylist(old spotify.FullPlaylist) playlist {
	return playlist{
		id:  string(old.ID),
		uri: string(old.URI),
	}
}
