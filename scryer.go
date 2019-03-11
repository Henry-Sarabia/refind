package scry

import (
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
