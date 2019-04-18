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
	limitMax = 50
	timeShort = "short"
	timeMed = "medium"
	timeLong = "long"
)

var (
	errClientNil    = errors.New("client pointer is nil")
	errDataInvalid  = errors.New("invalid or empty data returned")
	errSeedsMissing = errors.New("missing seed input")
	errTracksMissing = errors.New("playlist track list is missing")
)
type clienter interface {
	artister
	recenter
	recommender
	playlister
}

type artister interface {
	CurrentUsersTopArtistsOpt(*spotify.Options) (*spotify.FullArtistPage, error)
}

type recenter interface {
	PlayerRecentlyPlayedOpt(*spotify.RecentlyPlayedOptions) ([]spotify.RecentlyPlayedItem, error)
}

type recommender interface {
	GetRecommendations(spotify.Seeds, *spotify.TrackAttributes, *spotify.Options) (*spotify.Recommendations, error)
}

type playlister interface {
	AddTracksToPlaylist(spotify.ID, ...spotify.ID) (string, error)
	CreatePlaylistForUser(string, string, string, bool) (*spotify.FullPlaylist, error)
	CurrentUser() (*spotify.PrivateUser, error)
}

type service struct {
	art artister
	rec recenter
	recom recommender
	play playlister
}

func New(c clienter) (*service, error) {
	if c == nil {
		return nil, errClientNil
	}
	s := &service{
		art: c,
		rec: c,
		recom: c,
		play: c,
	}

	return s, nil
}

func (s *service) TopArtists() ([]refind.Artist, error) {
	var top []refind.Artist

	short, err := s.topArtists(limitMax, timeShort)
	if err != nil {
		return nil, err
	}
	top = append(top, short...)

	med, err := s.topArtists(limitMax, timeMed)
	if err != nil {
		return nil, err
	}
	top = append(top, med...)

	long, err := s.topArtists(limitMax, timeLong)
	if err != nil {
		return nil, err
	}
	top = append(top, long...)

	return top, nil
}

func (s *service) topArtists(limit int, time string) ([]refind.Artist, error) {
	opt := &spotify.Options{
		Limit: &limit,
		Timerange: &time,
	}

	top, err := s.art.CurrentUsersTopArtistsOpt(opt)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	if top == nil {
		return nil, errDataInvalid
	}

	return parseArtists(top.Artists...), nil
}

func (s *service) RecentTracks() ([]refind.Track, error) {
	opt := &spotify.RecentlyPlayedOptions{
		Limit: limitMax,
	}

	rec, err := s.rec.PlayerRecentlyPlayedOpt(opt)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recently played tracks")
	}

	if len(rec) <= 0 {
		return nil, errDataInvalid
	}

	var t []refind.Track
	for _, r := range rec {
		t = append(t, parseTrack(r.Track))
	}

	return t, nil
}

func (s *service) Recommendations(seeds []refind.Seed) ([]refind.Track, error) {
	if len(seeds) <= 0 {
		return nil, errSeedsMissing
	}

	sds, err := parseSeeds(seeds)
	if err != nil {
		return nil, err
	}

	var tracks []refind.Track
	attr := spotify.NewTrackAttributes().TargetPopularity(popTarget).MaxPopularity(popMax)

	for _, sd := range sds {
		recs, err := s.recom.GetRecommendations(sd, attr, nil)
		if err != nil {
			return nil, errors.Wrap(err, "cannot fetch recommendations")
		}

		if recs == nil {
			return nil, errDataInvalid
		}

		t := parseSimpleTracks(recs.Tracks...)
		tracks = append(tracks, t...)
	}

	return tracks, nil
}

func (s *service) Playlist(name string, list []refind.Track) (*spotify.FullPlaylist, error) {
	if len(list) <= 0 {
		return nil, errTracksMissing
	}

	u, err := s.play.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch user")
	}

	if u == nil {
		return nil, errDataInvalid
	}

	pl, err := s.play.CreatePlaylistForUser(u.ID, name, "description", publicPlaylist)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create playlist")
	}

	if pl == nil {
		return nil, errDataInvalid
	}

	var IDs []spotify.ID
	for _, t := range list {
		IDs = append(IDs, spotify.ID(t.ID))
	}

	_, err = s.play.AddTracksToPlaylist(pl.ID, IDs...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot add tracks to playlist")
	}

	return pl, nil
}
