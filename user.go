package scry

import "github.com/zmb3/spotify"

type user struct {
	id string
}

func parseUser(old spotify.PrivateUser) user {
	return user{id: old.ID}
}
