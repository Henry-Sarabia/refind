package scry

import (
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

const (
	popTarget int = 40
	popMax    int = 50
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

func spotifySeeds(sdr Seeder) (spotify.Seeds, error) {
	sds := sdr.Seeds()
	var spot spotify.Seeds

	for _, sd := range sds {
		switch sd.Category {
		case trackSeed:
			spot.Tracks = append(spot.Tracks, spotify.ID(sd.ID))
		case artistSeed:
			spot.Artists = append(spot.Artists, spotify.ID(sd.ID))
		case genreSeed:
			spot.Genres = append(spot.Genres, sd.Name)
		default:
			return spotify.Seeds{}, errors.New("unexpected seed category")
		}
	}

	return spot, nil
}
