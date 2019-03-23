package scry

import "github.com/pkg/errors"

type Lister struct {
	serv MusicService
	rec  Recommender
}

type MusicService interface {
	TopArtists() ([]Artist, error)
	RecentTracks() ([]Track, error)
	Playlist(string, []Track) (*Playlist, error)
}

type Recommender interface {
	Recommendations([]Seed) ([]Track, error)
}

func (l Lister) FromTracks(name string) ([]Track, error) {
	tracks, err := l.serv.RecentTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recent tracks")
	}

	var sds []Seed
	for _, t := range tracks {
		sd, err := t.Seed()
		if err != nil {
			return nil, errors.Wrap(err, "one or more tracks are invalid seeds")
		}
		sds = append(sds, sd)
	}

	recs, err := l.rec.Recommendations(sds)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recommendations")
	}

	return recs, nil
}

func (l Lister) Filter(prev []Track, rmv map[string]Artist) []Track {
	var curr []Track
	for _, p := range prev {
		if _, ok := rmv[p.Artist.Name]; !ok {
			curr = append(curr, p)
		}
	}

	return curr
}

//func FromArtists(serv MusicService, rec Recommender, name string) (*Playlist, error) {
//	top, err := serv.TopArtists()
//	if err != nil {
//		return nil, errors.Wrap(err, "cannot fetch top artists")
//	}
//
//	var sds []Seed
//	for _, t := range top {
//		sd, err := t.Seed()
//		if err != nil {
//			return nil, errors.Wrap(err, "one or more artists are invalid seeds")
//		}
//		sds = append(sds, sd)
//	}
//
//	recs, err := rec.Recommendations(sds)
//	if err != nil {
//		return nil, errors.Wrap(err, "cannot fetch recommendations")
//	}
//
//	pl, err := serv.Playlist(name, recs)
//	if err != nil {
//		return nil, errors.Wrapf(err, "cannot create playlist with name: %s", name)
//	}
//
//	return pl, nil
//}
