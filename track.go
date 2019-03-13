package scry

type Track struct {
	ID     string
	Name   string
	Artist Artist
}

func (t Track) Seeds() Seed {

	return Seed{Category: TrackSeed, ID: t.ID}
}
