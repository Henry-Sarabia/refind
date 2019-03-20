package scry

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
)

type Track struct {
	ID     string
	Name   string
	Artist Artist
}

func (t Track) Seed() (Seed, error) {
	if blank.Is(t.ID) {
		return Seed{}, errors.New("cannot create track seed with missing id")
	}
	return Seed{Category: TrackSeed, ID: t.ID}, nil
}
