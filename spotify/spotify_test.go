package spotify

import (
	"encoding/json"
	"github.com/Henry-Sarabia/refind"
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
	"io/ioutil"
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
			wantServ: &service{
				art: &spotify.Client{},
				track: &spotify.Client{},
				rec: &spotify.Client{},
				recom: &spotify.Client{},
				play: &spotify.Client{},
			},
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

type fakeArtister struct {
	file string
	err error
}

func (f fakeArtister) CurrentUsersTopArtists() (*spotify.FullArtistPage, error) {
	b, err := ioutil.ReadFile(f.file)
	if err != nil {
		return nil, err
	}

	var artist *spotify.FullArtistPage
	if err := json.Unmarshal(b, &artist); err != nil {
		return nil, f.err
	}

	return artist, f.err
}

const testTopArtistFile string = "test_data/current_users_top_artists.json"
const testEmptyFile string = "test_data/empty.json"
var testErrNoData = errors.New("no data")

func TestService_TopArtists(t *testing.T) {
	tests := []struct {
		name string
		art artister
		wantArts []refind.Artist
		wantErr error
	}{
		{
			name: "Valid data, nil error",
			art: fakeArtister{
				file: testTopArtistFile,
				err: nil,
				},
			wantArts: []refind.Artist{
				{ID: "4Z8W4fKeB5YxbusRsdQVPb", Name: "Radiohead"},
				{ID: "3yY2gUcIsjMr8hjo51PoJ8", Name: "The Smiths"},
				{ID: "3iTsJGG39nMg9YiolUgLMQ", Name: "Morrissey"},
				{ID: "4BO8wK4OAaFsi6PSzs366S", Name: "Ricky Eat Acid"},
				{ID: "4uSftVc3FPWe6RJuMZNEe9", Name: "Andrew Bird"},
				{ID: "19I4tYiChJoxEO5EuviXpz", Name: "AFI"},
				{ID: "0Y6dVaC9DZtPNH4591M42W", Name: "TV Girl"},
				{ID: "7bu3H8JO7d0UbMoVzbo70s", Name: "The Cure"},
				{ID: "0Q2Tc5yZFJpumLMc7Yz4e4", Name: "Tomppabeats"},
				{ID: "19zqV9DV3txjMUjHvltl2D", Name: "Motion City Soundtrack"},
			},
			wantErr: nil,
		},
		{
			name: "Valid data, error",
			art: fakeArtister{
				file: testTopArtistFile,
				err: testErrNoData,
			},
			wantArts: nil,
			wantErr: testErrNoData,
		},
		{
			name: "No data, error",
			art: fakeArtister{
				file: testEmptyFile,
				err: testErrNoData,
			},
			wantArts: nil,
			wantErr: testErrNoData,
		},
		{
			name: "No data, no error",
			art: fakeArtister{
				file: testEmptyFile,
				err: nil,
			},
			wantArts: nil,
			wantErr: errInvalidData,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := service{art: test.art}

			got, err := serv.TopArtists()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantArts) {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantArts)
			}
		})
	}
}
