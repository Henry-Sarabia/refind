package buffer

import (
	"github.com/Henry-Sarabia/refind"
	"github.com/pkg/errors"
)

var errNilBuf = errors.New("cannot initialize new buffer using nil interface")

type buffer struct {
	serv refind.MusicService
	artists []refind.Artist
	tracks []refind.Track
}

func New(serv refind.MusicService) (*buffer, error) {
	if serv == nil {
		return nil, errNilBuf
	}
	return &buffer{serv: serv}, nil
}

func (b buffer) TopArtists() ([]refind.Artist, error) {
	if len(b.artists) > 0 {
		return b.artists, nil
	}

	top, err := b.serv.TopArtists()
	if err != nil {
		return nil, err
	}

	return top, nil
}

func (b buffer) RecentTracks() ([]refind.Track, error) {
	if len(b.tracks) > 0 {
		return b.tracks, nil
	}

	rec, err := b.serv.RecentTracks()
	if err != nil {
		return nil, err
	}

	return rec, nil
}
