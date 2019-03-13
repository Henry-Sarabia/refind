package scry

type Track struct {
	ID     string
	Name   string
	Artist Artist
}

func (t Track) Seed() Seed {
	return Seed{Category: TrackSeed, ID: t.ID}
}
