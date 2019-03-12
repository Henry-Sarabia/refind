package spotifyservice

import (
	"github.com/Henry-Sarabia/scry"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

func SpotifySeeds(sdr scry.Seeder) (spotify.Seeds, error) {
	sds := sdr.Seeds()
	var spot spotify.Seeds

	for _, sd := range sds {
		switch sd.Category {
		case scry.TrackSeed:
			spot.Tracks = append(spot.Tracks, spotify.ID(sd.ID))
		case scry.ArtistSeed:
			spot.Artists = append(spot.Artists, spotify.ID(sd.ID))
		case scry.GenreSeed:
			spot.Genres = append(spot.Genres, sd.Name)
		default:
			return spotify.Seeds{}, errors.New("unexpected Seed category")
		}
	}

	return spot, nil
}
