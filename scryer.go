package scry

import (
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

const (
	popTarget      int  = 40
	popMax         int  = 50
	publicPlaylist bool = true
)

type scryer struct {
	c *spotify.Client
}

func (sc *scryer) CurrentUser() (user, error) {
	u, err := sc.c.CurrentUser()
	if err != nil {
		return user{}, errors.Wrap(err, "cannot fetch user")
	}

	return parseUser(*u), nil
}

func (sc *scryer) TopArtists() ([]artist, error) {
	top, err := sc.c.CurrentUsersTopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	return parseArtists(top.Artists...), nil
}

func (sc *scryer) TopTracks() ([]track, error) {
	top, err := sc.c.CurrentUsersTopTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top tracks")
	}

	return parseFullTracks(top.Tracks...), nil
}

func (sc *scryer) RecentTracks() ([]track, error) {
	rec, err := sc.c.PlayerRecentlyPlayed()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recently played tracks")
	}

	var t []track
	for _, r := range rec {
		t = append(t, parseTrack(r.Track))
	}

	return t, nil
}

func (sc *scryer) Recommendation(sdr Seeder) ([]track, error) {
	sds, err := spotifySeeds(sdr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create seeds from seeder")
	}

	attr := spotify.NewTrackAttributes().TargetPopularity(popTarget).MaxPopularity(popMax)
	recs, err := sc.c.GetRecommendations(sds, attr, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recommendations")
	}

	return parseSimpleTracks(recs.Tracks...), nil
}

func (sc *scryer) Playlist(name string, tracks []track) (playlist, error) {
	u, err := sc.CurrentUser()
	if err != nil {
		return playlist{}, err
	}

	pl, err := sc.c.CreatePlaylistForUser(u.id, name, "description", publicPlaylist)
	if err != nil {
		return playlist{}, errors.Wrap(err, "cannot create playlist")
	}

	var IDs []spotify.ID
	for _, t := range tracks {
		IDs = append(IDs, spotify.ID(t.id))
	}

	_, err = sc.c.AddTracksToPlaylist(pl.ID, IDs...)
	if err != nil {
		return playlist{}, errors.Wrap(err, "cannot add tracks to playlist")
	}

	return parsePlaylist(*pl), nil
}
