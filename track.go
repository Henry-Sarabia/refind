package refind

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
)

var errTrackSeed = errors.New("cannot create track seed with missing id")
type Track struct {
	ID     string
	Name   string
	Artist Artist
}

func (t Track) Seed() (Seed, error) {
	if blank.Is(t.ID) {
		return Seed{}, errTrackSeed
	}
	return Seed{Category: TrackSeed, ID: t.ID}, nil
}
