package scry

import (
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

type Seeder interface {
	Seeds() []seed
}

type seed struct {
	Category seedCategory
	ID       string
	Name     string
}

type seedCategory int

const (
	trackSeed seedCategory = iota
	artistSeed
	genreSeed
)

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
