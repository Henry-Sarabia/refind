package scry

type Seeder interface {
	Seeds() []Seed
}

type Seed struct {
	Category SeedCategory
	ID       string
	Name     string
}

type SeedCategory int

const (
	TrackSeed SeedCategory = iota
	ArtistSeed
	GenreSeed
)
