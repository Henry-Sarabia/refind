package refind

type buffer struct {
	serv MusicService
	artists []Artist
	tracks []Track
}

func newBuffer(serv MusicService) *buffer {
	return &buffer{serv: serv}
}

func (b buffer) TopArtists() ([]Artist, error) {
	var top []Artist
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

func (b buffer) RecentTracks() ([]Track, error) {
	var rec []Track
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
