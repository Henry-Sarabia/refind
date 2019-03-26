package refind

type Seed struct {
	Category SeedCategory
	ID       string
}

type SeedCategory int

const (
	TrackSeed SeedCategory = iota
	ArtistSeed
	GenreSeed
)
