package spotify

import (
	"github.com/Henry-Sarabia/refind"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

const (
	popTarget      int  = 40
	popMax         int  = 50
	publicPlaylist bool = true
)

var errNilClient = errors.New("client pointer is nil")

type clienter interface {
	CurrentUser() (*spotify.PrivateUser, error)
	CurrentUsersTopArtists() (*spotify.FullArtistPage, error)
	CurrentUsersTopTracks() (*spotify.FullTrackPage, error)
	PlayerRecentlyPlayed() ([]spotify.RecentlyPlayedItem, error)
	GetRecommendations(spotify.Seeds, *spotify.TrackAttributes, *spotify.Options) (*spotify.Recommendations, error)
	CreatePlaylistForUser(string, string, string, bool) (*spotify.FullPlaylist, error)
	AddTracksToPlaylist(spotify.ID, ...spotify.ID) (string, error)
	GetArtists(...spotify.ID) ([]*spotify.FullArtist, error)
}

type service struct {
	c clienter
}

func New(c clienter) (*service, error) {
	if c == nil {
		return nil, errNilClient
	}
	s := &service{c: c}

	return s, nil
}

func (s *service) TopArtists() ([]refind.Artist, error) {
	top, err := s.c.CurrentUsersTopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	return parseArtists(top.Artists...), nil
}

func (s *service) topTracks() ([]refind.Track, error) {
	top, err := s.c.CurrentUsersTopTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top tracks")
	}

	return parseFullTracks(top.Tracks...), nil
}

func (s *service) RecentTracks() ([]refind.Track, error) {
	rec, err := s.c.PlayerRecentlyPlayed()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recently played tracks")
	}

	var t []refind.Track
	for _, r := range rec {
		t = append(t, parseTrack(r.Track))
	}

	return t, nil
}

func (s *service) Recommendations(seeds []refind.Seed) ([]refind.Track, error) {
	sds, err := parseSeeds(seeds)
	if err != nil {
		return nil, err
	}

	var tracks []refind.Track
	attr := spotify.NewTrackAttributes().TargetPopularity(popTarget).MaxPopularity(popMax)

	for _, sd := range sds {
		recs, err := s.c.GetRecommendations(sd, attr, nil)
		if err != nil {
			return nil, errors.Wrap(err, "cannot fetch recommendations")
		}

		t := parseSimpleTracks(recs.Tracks...)
		tracks = append(tracks, t...)
	}

	return tracks, nil
}

func (s *service) Playlist(name string, list []refind.Track) (*spotify.FullPlaylist, error) {
	u, err := s.c.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch user")
	}

	pl, err := s.c.CreatePlaylistForUser(u.ID, name, "description", publicPlaylist)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create playlist")
	}

	var IDs []spotify.ID
	for _, t := range list {
		IDs = append(IDs, spotify.ID(t.ID))
	}

	_, err = s.c.AddTracksToPlaylist(pl.ID, IDs...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot add tracks to playlist")
	}

	return pl, nil
}
