package scry

import (
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

type scryer struct {
	c *spotify.Client
}

func (sc *scryer) CurrentUser() (string, error) {
	u, err := sc.c.CurrentUser()
	if err != nil {
		return "", err
	}

	return u.ID, nil
}

func (sc *scryer) TopArtists() ([]artist, error) {
	top, err := sc.c.CurrentUsersTopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	var art []artist
	for _, a := range top.Artists {
		art = append(art, newArtist(a.SimpleArtist))
	}

	return art, nil
}
