package buffer

import (
	"github.com/Henry-Sarabia/refind"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

var (
	testErrArtists = errors.New("cannot fetch artists")
	testErrTracks = errors.New("cannot fetch tracks")
)
var testArtists = []refind.Artist{
	{ID: "0", Name: "foo"},
	{ID: "1", Name: "bar"},
	{ID: "2", Name: "baz"},
}

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

func TestBuffer_TopArtists(t *testing.T) {
	tests := []struct {
		name string
		buf buffer
		wantArt []refind.Artist
		wantErr error
	}{
		{
			name: "Valid response, nil buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: testArtists,
					artistErr: nil,
				},
				artists: nil,
			},
			wantArt: testArtists,
			wantErr: nil,
		},
		{
			name: "Valid response, empty buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: testArtists,
					artistErr: nil,
				},
				artists: []refind.Artist{},
			},
			wantArt: testArtists,
			wantErr: nil,
		},
		{
			name: "Valid response, single artist buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: testArtists,
					artistErr: nil,
				},
				artists: []refind.Artist{{ID: "1", Name: "bar"}},
			},
			wantArt: []refind.Artist{{ID: "1", Name: "bar"}},
			wantErr: nil,
		},
		{
			name: "Valid response, multiple artist buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: testArtists,
					artistErr: nil,
				},
				artists: testArtists,
			},
			wantArt: testArtists,
			wantErr: nil,
		},
		{
			name: "Invalid response, nil buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: nil,
					artistErr: testErrArtists,
				},
				artists: nil,
			},
			wantArt: nil,
			wantErr: testErrArtists,
		},
		{
			name: "Invalid response, empty buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: nil,
					artistErr: testErrArtists,
				},
				artists: []refind.Artist{},
			},
			wantArt: nil,
			wantErr: testErrArtists,
		},
		{
			name: "Invalid response, single artist buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: nil,
					artistErr: testErrArtists,
				},
				artists: []refind.Artist{{ID: "1", Name: "bar"}},
			},
			wantArt: []refind.Artist{{ID: "1", Name: "bar"}},
			wantErr: nil,
		},
		{
			name: "Invalid response, multiple artist buffer",
			buf: buffer{
				serv: fakeMusicService{
					artists: nil,
					artistErr: testErrArtists,
				},
				artists: testArtists,
			},
			wantArt: testArtists,
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.buf.TopArtists()
			if err != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantArt) {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantArt)
			}
		})
	}
}
