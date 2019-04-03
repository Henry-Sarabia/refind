package spotify

import (
	"github.com/zmb3/spotify"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		c clienter
		wantServ *service
		wantErr error
	}{
		{
			name: "Nil client",
			c: nil,
			wantServ: nil,
			wantErr: errNilClient,
		},
		{
			name: "Valid client",
			c: &spotify.Client{},
			wantServ: &service{c: &spotify.Client{}},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := New(test.c)
			if err != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantServ) {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantServ)
			}
		})
	}
}

func TestService_TopArtists(t *testing.T) {
	tests := []struct {
		name string
	}{

		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
