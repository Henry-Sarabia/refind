package scry

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

type MusicService interface {
	CurrentUser() (User, error)
	TopArtists() ([]Artist, error)
	TopTracks() ([]Track, error)
	RecentTracks() ([]Track, error)
	Recommendation(Seeder) ([]Track, error)
	Playlist(string, []Track) (Playlist, error)
}

func Authenticator(URI string) (*spotify.Authenticator, error) {
	if blank.Is(URI) {
		return nil, errors.New("URI is blank")
	}

	auth := spotify.NewAuthenticator(
		URI,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopeUserReadPrivate,
		spotify.ScopeUserTopRead,
		spotify.ScopeUserReadRecentlyPlayed,
	)

	return &auth, nil
}

type Scryer struct {
	MusicService
}

func New(ms MusicService) (*Scryer, error) {
	return &Scryer{MusicService: ms}, nil
}
