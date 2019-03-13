package spotifyservice

import (
	"github.com/Henry-Sarabia/scry"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

type Seed struct {
	spotify.Seeds
}

func ParseSeeds(old []scry.Seed) ([]spotify.Seeds, error) {
	var sds []spotify.Seeds

	for len(old) > 0 {
		sd, err := parseMaxSeeds(old)
		if err != nil {
			return nil, errors.Wrap(err, "one or more seeds cannot be parsed")
		}
		sds = append(sds, sd)

		if len(old) > spotify.MaxNumberOfSeeds {
			old = old[:spotify.MaxNumberOfSeeds]
		}
	}

	return sds, nil
}

func parseMaxSeeds(old []scry.Seed) (spotify.Seeds, error) {
	var sd spotify.Seeds

	for _, o := range old {
		if len(sd.Tracks)+len(sd.Artists)+len(sd.Genres) >= spotify.MaxNumberOfSeeds {
			break
		}

		err := parseSeed(o, &sd)
		if err != nil {
			return spotify.Seeds{}, errors.Wrap(err, "one or more seeds cannot be parsed")
		}
	}

	return sd, nil
}

func parseSeed(old scry.Seed, sd *spotify.Seeds) error {
	switch old.Category {
	case scry.TrackSeed:
		sd.Tracks = append(sd.Tracks, spotify.ID(old.ID))
	case scry.ArtistSeed:
		sd.Artists = append(sd.Artists, spotify.ID(old.ID))
	case scry.GenreSeed:
		sd.Genres = append(sd.Genres, old.ID)
	default:
		return errors.New("unexpected Seed Category")
	}

	return nil
}

//func Seeds(sdr scry.Seeder) (spotify.Seeds, error) {
//	sds := sdr.Seeds()
//	var spot spotify.Seeds
//
//	for _, sd := range sds {
//		switch sd.Category {
//		case scry.TrackSeed:
//			spot.Tracks = append(spot.Tracks, spotify.ID(sd.ID))
//		case scry.ArtistSeed:
//			spot.Artists = append(spot.Artists, spotify.ID(sd.ID))
//		case scry.GenreSeed:
//			spot.Genres = append(spot.Genres, sd.Name)
//		default:
//			return spotify.Seeds{}, errors.New("unexpected Seed category")
//		}
//	}
//
//	return spot, nil
//}
