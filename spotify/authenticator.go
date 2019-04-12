package spotify

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

var errMissingURI = errors.New("URI is blank")

func Authenticator(URI string) (*spotify.Authenticator, error) {
	if blank.Is(URI) {
		return nil, errMissingURI
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
