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
	var top []refind.Artist
	if b.artists != nil {
		top = b.artists
	}

	var err error
	top, err = b.serv.TopArtists()
	if err != nil {
		return nil, err
	}

	return top, nil
}

func (b buffer) RecentTracks() ([]refind.Track, error) {
	var rec []refind.Track
	if b.tracks != nil {
		rec = b.tracks
	}

	var err error
	rec, err = b.serv.RecentTracks()
	if err != nil {
		return nil, err
	}

	return rec, nil
}
