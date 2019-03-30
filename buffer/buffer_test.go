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

var testTracks = []refind.Track{
	{ID: "10", Name: "corge", Artist: refind.Artist{ID: "0", Name: "foo"}},
	{ID: "11", Name: "grault", Artist: refind.Artist{ID: "1", Name: "bar"}},
	{ID: "12" , Name: "garply", Artist: refind.Artist{ID: "1", Name: "bar"}},
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

func TestBuffer_RecentTracks(t *testing.T) {
	tests := []struct {
		name string
		buf buffer
		wantTracks []refind.Track
		wantErr error
	}{
		{
			name: "Valid response, nil buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: testTracks,
					trackErr: nil,
				},
				tracks: nil,
			},
			wantTracks: testTracks,
			wantErr: nil,
		},
		{
			name: "Valid response, empty buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: testTracks,
					trackErr: nil,
				},
				tracks: []refind.Track{},
			},
			wantTracks: testTracks,
			wantErr: nil,
		},
		{
			name: "Valid response, single track buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: testTracks,
					trackErr: nil,
				},
				tracks: []refind.Track{{ID: "11", Name: "grault", Artist: refind.Artist{ID: "1", Name: "bar"}}},
			},
			wantTracks: []refind.Track{{ID: "11", Name: "grault", Artist: refind.Artist{ID: "1", Name: "bar"}}},
			wantErr: nil,
		},
		{
			name: "Valid response, multiple track buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: testTracks,
					trackErr: nil,
				},
				tracks: testTracks,
			},
			wantTracks: testTracks,
			wantErr: nil,
		},
		{
			name: "Invalid response, nil buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: nil,
					trackErr: testErrTracks,
				},
				tracks: nil,
			},
			wantTracks: nil,
			wantErr: testErrTracks,
		},
		{
			name: "Invalid response, empty buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: nil,
					trackErr: testErrTracks,
				},
				tracks: []refind.Track{},
			},
			wantTracks: nil,
			wantErr: testErrTracks,
		},
		{
			name: "Invalid response, single track buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: nil,
					trackErr: testErrTracks,
				},
				tracks: []refind.Track{{ID: "11", Name: "grault", Artist: refind.Artist{ID: "1", Name: "bar"}}},
			},
			wantTracks: []refind.Track{{ID: "11", Name: "grault", Artist: refind.Artist{ID: "1", Name: "bar"}}},
			wantErr: nil,
		},
		{
			name: "Invalid response, multiple track buffer",
			buf: buffer{
				serv: fakeMusicService{
					tracks: nil,
					trackErr: testErrTracks,
				},
				tracks: testTracks,
			},
			wantTracks: testTracks,
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.buf.RecentTracks()
			if err != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantTracks) {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantTracks)
			}
		})
	}
}
