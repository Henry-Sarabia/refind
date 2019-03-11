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

func (sc *scryer) TopTracks() ([]track, error) {
	top, err := sc.c.CurrentUsersTopTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top tracks")
	}

	var trk []track
	for _, t := range top.Tracks {
		trk = append(trk, newTrack(t.SimpleTrack))
	}

	return trk, nil
}

func (sc *scryer) RecentTracks() ([]track, error) {
	rec, err := sc.c.PlayerRecentlyPlayed()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recently played tracks")
	}

	var trk []track
	for _, r := range rec {
		trk = append(trk, newTrack(r.Track))
	}

	return trk, nil
}
