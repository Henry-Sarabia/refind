package refind

type Buffer struct {
	serv MusicService
	artists []Artist
	tracks []Track
}

func NewBuffer(serv MusicService) *Buffer {
	return &Buffer{serv: serv}
}

func (b Buffer) TopArtists() ([]Artist, error) {
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

func (b Buffer) RecentTracks() ([]Track, error) {
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
