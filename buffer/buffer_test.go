package buffer

import (
	"github.com/Henry-Sarabia/refind"
	"reflect"
	"testing"
)

type fakeMusicService struct {
	artists []refind.Artist
	artistErr error
	tracks []refind.Track
	trackErr error
}

func (f fakeMusicService) TopArtists() ([]refind.Artist, error) {
	return f.artists, f.artistErr
}

func (f fakeMusicService) RecentTracks() ([]refind.Track, error) {
	return f.tracks, f.trackErr
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		serv refind.MusicService
		wantBuf *buffer
		wantErr error
	}{
		{
			"Nil interface",
			nil,
			nil,
			errNilBuf,
		},
		{
			"Valid interface",
			fakeMusicService{},
			&buffer{serv: fakeMusicService{}},
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := New(test.serv)
			if err != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantBuf) {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantBuf)
			}
		})
	}
}
