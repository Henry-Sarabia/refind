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

func (s *Scryer) FromTracks() ([]Track, error) {
	rec, err := s.serv.RecentTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recent tracks")
	}

	var sds []Seed
	for _, r := range rec {
		sds = append(sds, r.Seed())
	}

	recs, err := s.serv.Recommendations(sds)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recommendations")
	}

	return recs, nil
}
