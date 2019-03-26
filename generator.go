package refind

import "github.com/pkg/errors"

type Generator struct {
	serv MusicService
	rec  Recommender
}

type MusicService interface {
	TopArtists() ([]Artist, error)
	RecentTracks() ([]Track, error)
}

type Recommender interface {
	Recommendations([]Seed) ([]Track, error)
}

func (g Generator) Tracklist() ([]Track, error) {
	list, err := g.fromTracks()
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate tracklist from track data")
	}

	return list, nil
}

func (g Generator) fromTracks() ([]Track, error) {
	tracks, err := g.serv.RecentTracks()
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

	recs, err := g.rec.Recommendations(sds)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch recommendations")
	}

	top, err := g.serv.TopArtists()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch top artists")
	}

	f := filter(recs, toMap(top))

	return f, nil
}

func toMap(prev []Artist) map[string]Artist {
	if len(prev) == 0 {
		return nil
	}

	curr := make(map[string]Artist)
	for _, p := range prev {
		curr[p.Name] = p
	}

	return curr
}

func filter(prev []Track, rmv map[string]Artist) []Track {
	if len(prev) == 0 {
		return nil
	}

	if len(rmv) == 0 {
		return prev
	}

	var curr []Track
	for _, p := range prev {
		if _, ok := rmv[p.Artist.Name]; !ok {
			curr = append(curr, p)
		}
	}

	return curr
}