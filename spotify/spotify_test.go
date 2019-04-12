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

const (
	testFileEmpty        string = "test_data/empty.json"
	testFileTopArtists   string = "test_data/current_users_top_artists.json"
	testFileRecentTracks string = "test_data/player_recently_played.json"
	testFileRecommendations string = "test_data/get_recommendations.json"
)

var testErrNoData = errors.New("no data")

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		c        clienter
		wantServ *service
		wantErr  error
	}{
		{
			name:     "Nil client",
			c:        nil,
			wantServ: nil,
			wantErr:  errNilClient,
		},
		{
			name: "Valid client",
			c:    &spotify.Client{},
			wantServ: &service{
				art:   &spotify.Client{},
				track: &spotify.Client{},
				rec:   &spotify.Client{},
				recom: &spotify.Client{},
				play:  &spotify.Client{},
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
	err  error
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

func TestService_TopArtists(t *testing.T) {
	tests := []struct {
		name     string
		art      artister
		wantArts []refind.Artist
		wantErr  error
	}{
		{
			name: "Valid data, nil error",
			art: fakeArtister{
				file: testFileTopArtists,
				err:  nil,
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
				file: testFileTopArtists,
				err:  testErrNoData,
			},
			wantArts: nil,
			wantErr:  testErrNoData,
		},
		{
			name: "No data, nil error",
			art: fakeArtister{
				file: testFileEmpty,
				err:  nil,
			},
			wantArts: nil,
			wantErr:  errInvalidData,
		},
		{
			name: "No data, error",
			art: fakeArtister{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			wantArts: nil,
			wantErr:  testErrNoData,
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

type fakeRecenter struct {
	file string
	err  error
}

func (f fakeRecenter) PlayerRecentlyPlayed() ([]spotify.RecentlyPlayedItem, error) {
	b, err := ioutil.ReadFile(f.file)
	if err != nil {
		return nil, err
	}

	var rec []spotify.RecentlyPlayedItem
	if err := json.Unmarshal(b, &rec); err != nil {
		return nil, err
	}

	return rec, f.err
}

func TestService_RecentTracks(t *testing.T) {
	tests := []struct {
		name       string
		rec        recenter
		wantTracks []refind.Track
		wantErr    error
	}{
		{
			name: "Valid data, nil error",
			rec: fakeRecenter{
				file: testFileRecentTracks,
				err:  nil,
			},
			wantTracks: []refind.Track{
				{ID: "5ETM3aBrDf45TWg9AgnWQD", Name: "Black Nostaljack AKA Come On", Artist: refind.Artist{ID: "4oLZx5FplbgfM8DEe9U8LB", Name: "Camp Lo"}},
				{ID: "1s92LwFTivD2f9o0s2hb78", Name: "Naturally Born", Artist: refind.Artist{ID: "099tLNCZZvtjC7myKD0mFp", Name: "Kool G Rap"}},
				{ID: "0FcAIIz4Ti87cFBwyD3iCE", Name: "Little Darlin Seize the Sun", Artist: refind.Artist{ID: "4CMC2nnStv4EENjKBSDpKR", Name: "Christina Vantzou"}},
				{ID: "0brnyKRZKnNngbH444p8cn", Name: "Prince of the Sea", Artist: refind.Artist{ID: "4G1ZsxfEEztbE1VcnNInPg", Name: "Chihei Hatakeyama"}},
				{ID: "5nP1e5QSwT07XR2zpTVJGc", Name: "Ninteen Seventy Something", Artist: refind.Artist{ID: "1wo9h8DP7M0M1orKuGZgWv", Name: "Masta Ace"}},
				{ID: "53aUYPTwJe6YrbSs8lQCEF", Name: "Buck Em Down", Artist: refind.Artist{ID: "2yN6bq26wynQcRuPkBYTDb", Name: "Black Moon"}},
				{ID: "1qKsRg2PvBzhWkMOpanQq3", Name: "Hiatus", Artist: refind.Artist{ID: "6AdRO941ZEDh4GHcCUdEs4", Name: "Rafael Anton Irisarri"}},
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantErr: nil,
		},
		{
			name: "Valid data, error",
			rec: fakeRecenter{
				file: testFileRecentTracks,
				err:  testErrNoData,
			},
			wantTracks: nil,
			wantErr: testErrNoData,
		},
		{
			name: "No data, nil error",
			rec: fakeRecenter{
				file: testFileEmpty,
				err:  nil,
			},
			wantTracks: nil,
			wantErr: errInvalidData,
		},
		{
			name: "No data, error",
			rec: fakeRecenter{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			wantTracks: nil,
			wantErr: testErrNoData,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := service{rec: test.rec}

			got, err := serv.RecentTracks()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantTracks) {
				t.Errorf("\ngot:  <%v>, \nwant: <%v>", got, test.wantTracks)
			}
		})
	}
}

type fakeRecommender struct {
	file string
	err error
}

func (f fakeRecommender) GetRecommendations(sds spotify.Seeds, attr *spotify.TrackAttributes, opt *spotify.Options) (*spotify.Recommendations, error) {
	b, err := ioutil.ReadFile(f.file)
	if err != nil {
		return nil, err
	}

	var rec *spotify.Recommendations
	if err := json.Unmarshal(b, &rec); err != nil {
		return nil, f.err
	}

	return rec, f.err
}

func TestService_Recommendations(t *testing.T) {
	tests := []struct {
		name string
		recom recommender
		sds []refind.Seed
		wantTracks []refind.Track
		wantErr error
	}{
		{
			name: "Valid data, valid seeds, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: nil,
			},
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: []refind.Track{
				{ID: "7cgi6lRggiLAAzsuJOBBeW", Name: "Innsbruck, ich muß dich lassen", Artist: refind.Artist{ID: "1G6jUCigH2z7oGk7jm6OhS", Name: "Heinrich Isaac"}},
				{ID: "0T02WlrUAK45ApAVVixmcc", Name: "La Bohème / Act 1: \"Che gelida manina\"", Artist: refind.Artist{ID: "0OzxPXyowUEQ532c9AmHUR", Name: "Giacomo Puccini"}},
				{ID: "0GaVkII433PqC4EkMSjWEV", Name: "Moments - Seeb Remix", Artist: refind.Artist{ID: "4NHQUGzhtTLFvgF5SZesLK", Name: "Tove Lo"}},
				{ID: "1H9rGpQ1Xqh45Y13mzfJvU", Name: "Clarinet Concerto in B-Flat Major (reconstructed R. Meylan): I. Andante sostenuto", Artist: refind.Artist{ID: "2jCGEMSZXMSOImpD8sqo56", Name: "Gaetano Donizetti"}},
				{ID: "4BNUJM7oEYNPtXDzvZjcRQ", Name: "The Scene", Artist: refind.Artist{ID: "4FJPplt1JOVw8Q7NiwFmLv", Name: "Friend Within"}},
				{ID: "3lO38SiB2WAQRqTAHN7WTC", Name: "Borderline - Vanic Remix", Artist: refind.Artist{ID: "2QSPrJfYeRXaltEEiriXN9", Name: "Tove Styrke"}},
				{ID: "6zzZPhrTwS84pOkuqCwI5B", Name: "I Didn’t Just Come Here To Dance", Artist: refind.Artist{ID: "6sFIWsNpZYqfjUpaCgueju", Name: "Carly Rae Jepsen"}},
				{ID: "4dGJf1SER1T6ooX46vwzRB", Name: "Chicken Fried", Artist: refind.Artist{ID: "6yJCxee7QumYr820xdIsjo", Name: "Zac Brown Band"}},
				{ID: "25I4pBnup7EeerWd61G61i", Name: "Timebomb", Artist: refind.Artist{ID: "4NHQUGzhtTLFvgF5SZesLK", Name: "Tove Lo"}},
				{ID: "2URjwQulkDiDmFdjSPrcSc", Name: "Appalachian Spring: Moderato - Coda", Artist: refind.Artist{ID: "0nJvyjVTb8sAULPYyA1bqU", Name: "Aaron Copland"}},
			},
			wantErr: nil,
		},
		{
			name: "Valid data, valid seeds, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: testErrNoData,
			},
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr: testErrNoData,
		},
		{
			name: "Valid data, no seeds, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: nil,
			},
			sds: nil,
			wantTracks: nil,
			wantErr: errMissingSeeds,
		},
		{
			name: "Valid data, no seeds, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: testErrNoData,
			},
			sds: nil,
			wantTracks: nil,
			wantErr: errMissingSeeds,
		},
		{
			name: "Valid data, invalid seed ID, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: nil,
			},
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr: errSeedBlankID,
		},
		{
			name: "Valid data, invalid seed ID, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: testErrNoData,
			},
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr: errSeedBlankID,
		},
		{
			name: "Valid data, invalid seed category, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: nil,
			},
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr: errSeedCategory,
		},
		{
			name: "Valid data, invalid seed category, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err: testErrNoData,
			},
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr: errSeedCategory,
		},
		{
			name: "No data, valid seeds, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: nil,
			},
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr: errInvalidData,
		},
		{
			name: "No data, valid seeds, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: testErrNoData,
			},
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr: testErrNoData,
		},
		{
			name: "No data, no seeds, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: nil,
			},
			sds: nil,
			wantTracks: nil,
			wantErr: errMissingSeeds,
		},
		{
			name: "No data, no seeds, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: testErrNoData,
			},
			sds: nil,
			wantTracks: nil,
			wantErr: errMissingSeeds,
		},
		{
			name: "No data, invalid seed ID, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: nil,
			},
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr: errSeedBlankID,
		},
		{
			name: "No data, invalid seed ID, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: testErrNoData,
			},
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr: errSeedBlankID,
		},
		{
			name: "No data, invalid seed category, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: nil,
			},
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr: errSeedCategory,
		},
		{
			name: "No data, invalid seed category, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err: testErrNoData,
			},
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr: errSeedCategory,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := service{recom: test.recom}

			got, err := serv.Recommendations(test.sds)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantTracks) {
				t.Errorf("\ngot:  <%v>, \nwant: <%v>", got, test.wantTracks)
			}
		})
	}
}
