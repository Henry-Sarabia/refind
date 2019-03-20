package scry

import "github.com/pkg/errors"

type MusicService interface {
	CurrentUser() (User, error)
	TopArtists() ([]Artist, error)
	TopTracks() ([]Track, error)
	RecentTracks() ([]Track, error)
	Recommendations([]Seed) ([]Track, error)
	Playlist(string, []Track) (Playlist, error)
}

type Scryer struct {
	serv MusicService
}

func New(ms MusicService) (*Scryer, error) {
	return &Scryer{serv: ms}, nil
}

func (s *Scryer) FromTracks(name string) (*Playlist, error) {
	rec, err := s.serv.RecentTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recent tracks")
	}

	var sds []Seed
	for _, r := range rec {
		sd, err := r.Seed()
		if err != nil {
			return nil, errors.Wrap(err, "one or more tracks are invalid seeds")
		}
		sds = append(sds, sd)
	}

	recs, err := s.serv.Recommendations(sds)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recommendations")
	}

	pl, err := s.serv.Playlist(name, recs)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot create playlist with name: %s", name)
	}

	return &pl, nil
}

func (s *Scryer) FromArtists(name string) (*Playlist, error) {
	top, err := s.serv.TopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	var sds []Seed
	for _, t := range top {
		sd, err := t.Seed()
		if err != nil {
			return nil, errors.Wrap(err, "one or more artists are invalid seeds")
		}
		sds = append(sds, sd)
	}

	recs, err := s.serv.Recommendations(sds)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recommendations")
	}

	pl, err := s.serv.Playlist(name, recs)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot create playlist with name: %s", name)
	}

	return &pl, nil
}
