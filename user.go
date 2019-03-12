package scry

import "github.com/zmb3/spotify"

type User struct {
	ID string
}

func ParseUser(old spotify.PrivateUser) User {
	return User{ID: old.ID}
}
