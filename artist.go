package refind

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
)

type Artist struct {
	ID   string
	Name string
}

func (a Artist) Seed() (Seed, error) {
	if blank.Is(a.ID) {
		return Seed{}, errors.New("cannot create artist seed with missing id")
	}

	return Seed{Category: ArtistSeed, ID: a.ID}, nil
}
