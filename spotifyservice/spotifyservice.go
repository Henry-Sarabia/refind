package spotifyservice

import (
	"github.com/Henry-Sarabia/scry"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

const (
	popTarget      int  = 40
	popMax         int  = 50
	publicPlaylist bool = true
)

type SpotifyService struct {
	c *spotify.Client
}

func New(c *spotify.Client) (*SpotifyService, error) {
	if c == nil {
		return nil, errors.New("client pointer is nil")
	}
	s := &SpotifyService{c: c}

	return s, nil
}

func (sp *SpotifyService) CurrentUser() (scry.User, error) {
	u, err := sp.c.CurrentUser()
	if err != nil {
		return scry.User{}, errors.Wrap(err, "cannot fetch user")
	}

	return ParseUser(*u), nil
}

func (sp *SpotifyService) TopArtists() ([]scry.Artist, error) {
	top, err := sp.c.CurrentUsersTopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	return ParseArtists(top.Artists...), nil
}

func (sp *SpotifyService) TopTracks() ([]scry.Track, error) {
	top, err := sp.c.CurrentUsersTopTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top tracks")
	}

	return ParseFullTracks(top.Tracks...), nil
}

func (sp *SpotifyService) RecentTracks() ([]scry.Track, error) {
	rec, err := sp.c.PlayerRecentlyPlayed()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recently played tracks")
	}

	var t []scry.Track
	for _, r := range rec {
		t = append(t, ParseTrack(r.Track))
	}

	return t, nil
}

func (sp *SpotifyService) Recommendations(seeds []scry.Seed) ([]scry.Track, error) {
	sds, err := ParseSeeds(seeds)
	if err != nil {
		return nil, err
	}

	var tracks []scry.Track
	attr := spotify.NewTrackAttributes().TargetPopularity(popTarget).MaxPopularity(popMax)

	for _, sd := range sds {
		recs, err := sp.c.GetRecommendations(sd, attr, nil)
		if err != nil {
			return nil, errors.Wrap(err, "cannot fetch recommendations")
		}

		t := ParseSimpleTracks(recs.Tracks...)
		tracks = append(tracks, t...)
	}

	return tracks, nil
}

func (sp *SpotifyService) Playlist(name string, tracks []scry.Track) (scry.Playlist, error) {
	u, err := sp.CurrentUser()
	if err != nil {
		return scry.Playlist{}, err
	}

	pl, err := sp.c.CreatePlaylistForUser(u.ID, name, "description", publicPlaylist)
	if err != nil {
		return scry.Playlist{}, errors.Wrap(err, "cannot create playlist")
	}

	var IDs []spotify.ID
	for _, t := range tracks {
		IDs = append(IDs, spotify.ID(t.ID))
	}

	_, err = sp.c.AddTracksToPlaylist(pl.ID, IDs...)
	if err != nil {
		return scry.Playlist{}, errors.Wrap(err, "cannot add tracks to playlist")
	}

	return ParsePlaylist(*pl), nil
}
