package spotify

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/Henry-Sarabia/refind"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

var (
	errSeedID       = errors.New("seed has missing or blank ID")
	errSeedCategory = errors.New("unexpected seed category")
)
type Seed struct {
	spotify.Seeds
}

func parseSeeds(old []refind.Seed) ([]spotify.Seeds, error) {
	var sds []spotify.Seeds

	for len(old) > 0 {
		sd, err := parseMaxSeeds(old)
		if err != nil {
			return nil, err
		}
		sds = append(sds, sd)

		if len(old) > spotify.MaxNumberOfSeeds {
			old = old[spotify.MaxNumberOfSeeds:]
		} else {
			old = nil
		}
	}

	return sds, nil
}

func parseMaxSeeds(old []refind.Seed) (spotify.Seeds, error) {
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

func parseSeed(old refind.Seed, sd *spotify.Seeds) error {
	if blank.Is(old.ID) {
		return errSeedID
	}

	switch old.Category {
	case refind.TrackSeed:
		sd.Tracks = append(sd.Tracks, spotify.ID(old.ID))
	case refind.ArtistSeed:
		sd.Artists = append(sd.Artists, spotify.ID(old.ID))
	case refind.GenreSeed:
		sd.Genres = append(sd.Genres, old.ID)
	default:
		return errSeedCategory
	}

	return nil
}
