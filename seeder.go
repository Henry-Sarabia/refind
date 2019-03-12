package scry

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
