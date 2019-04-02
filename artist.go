package refind

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
)

var errArtistSeed = errors.New("cannot create artist seed with missing id")

type Artist struct {
	ID   string
	Name string
}

func (a Artist) Seed() (Seed, error) {
	if blank.Is(a.ID) {
		return Seed{}, errArtistSeed
	}

	return Seed{Category: ArtistSeed, ID: a.ID}, nil
}
