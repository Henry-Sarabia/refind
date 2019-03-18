package scry

type Artist struct {
	ID   string
	Name string
}

func (a Artist) Seed() Seed {
	return Seed{Category: ArtistSeed, ID: a.Name}
}
