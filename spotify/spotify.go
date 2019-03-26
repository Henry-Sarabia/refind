package spotify

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

type Service struct {
	c *spotify.Client
}

func New(c *spotify.Client) (*Service, error) {
	if c == nil {
		return nil, errors.New("client pointer is nil")
	}
	s := &Service{c: c}

	return s, nil
}

func (s *Service) TopArtists() ([]scry.Artist, error) {
	top, err := s.c.CurrentUsersTopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	return ParseArtists(top.Artists...), nil
}

func (s *Service) TopTracks() ([]scry.Track, error) {
	top, err := s.c.CurrentUsersTopTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top tracks")
	}

	return ParseFullTracks(top.Tracks...), nil
}

func (s *Service) RecentTracks() ([]scry.Track, error) {
	rec, err := s.c.PlayerRecentlyPlayed()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recently played tracks")
	}

	var t []scry.Track
	for _, r := range rec {
		t = append(t, ParseTrack(r.Track))
	}

	return t, nil
}

func (s *Service) Recommendations(seeds []scry.Seed) ([]scry.Track, error) {
	sds, err := ParseSeeds(seeds)
	if err != nil {
		return nil, err
	}

	var tracks []scry.Track
	attr := spotify.NewTrackAttributes().TargetPopularity(popTarget).MaxPopularity(popMax)

	for _, sd := range sds {
		recs, err := s.c.GetRecommendations(sd, attr, nil)
		if err != nil {
			return nil, errors.Wrap(err, "cannot fetch recommendations")
		}

		t := ParseSimpleTracks(recs.Tracks...)
		tracks = append(tracks, t...)
	}

	return tracks, nil
}

func (s *Service) Playlist(name string, tracks []scry.Track) (*spotify.FullPlaylist, error) {
	u, err := s.c.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch user")
	}

	pl, err := s.c.CreatePlaylistForUser(u.ID, name, "description", publicPlaylist)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create playlist")
	}

	var IDs []spotify.ID
	for _, t := range tracks {
		IDs = append(IDs, spotify.ID(t.ID))
	}

	_, err = s.c.AddTracksToPlaylist(pl.ID, IDs...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot add tracks to playlist")
	}

	return pl, nil
}
