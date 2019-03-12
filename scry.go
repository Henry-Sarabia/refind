package scry

type MusicService interface {
	CurrentUser() (User, error)
	TopArtists() ([]Artist, error)
	TopTracks() ([]Track, error)
	RecentTracks() ([]Track, error)
	Recommendation(Seeder) ([]Track, error)
	Playlist(string, []Track) (Playlist, error)
}

type Scryer struct {
	MusicService
}

func New(ms MusicService) (*Scryer, error) {
	return &Scryer{MusicService: ms}, nil
}
