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
	testFileEmpty           string = "test_data/empty.json"
	testFileTopArtists      string = "test_data/current_users_top_artists.json"
	testFileRecentTracks    string = "test_data/player_recently_played.json"
	testFileRecommendations string = "test_data/get_recommendations.json"
	testFileCurrentUser     string = "test_data/current_user.json"
	testFileCreatePlaylist  string = "test_data/create_playlist_for_user.json"
)

var testErrNoData = errors.New("no data")

func TestAuthenticator(t *testing.T) {
	tests := []struct {
		name string
		uri string
		wantErr error
	}{
		{
			name: "Valid URI",
			uri: "some_valid_uri",
			wantErr: nil,
		},
		{
			name:    "Empty URI",
			uri:     "",
			wantErr: errMissingURI,
		},
		{
			name:    "Blank URI",
			uri:     "     ",
			wantErr: errMissingURI,
		},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Authenticator(test.uri)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}
		})
	}
}

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
			wantErr:  errClientNil,
		},
		{
			name: "Valid client",
			c:    &spotify.Client{},
			wantServ: &service{
				art:   &spotify.Client{},
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

func (f fakeArtister) CurrentUsersTopArtistsOpt(opt *spotify.Options) (*spotify.FullArtistPage, error) {
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
			wantErr:  errDataInvalid,
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

func (f fakeRecenter) PlayerRecentlyPlayedOpt(opt *spotify.RecentlyPlayedOptions) ([]spotify.RecentlyPlayedItem, error) {
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
			wantErr:    testErrNoData,
		},
		{
			name: "No data, nil error",
			rec: fakeRecenter{
				file: testFileEmpty,
				err:  nil,
			},
			wantTracks: nil,
			wantErr:    errDataInvalid,
		},
		{
			name: "No data, error",
			rec: fakeRecenter{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			wantTracks: nil,
			wantErr:    testErrNoData,
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
	err  error
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

const testTotal int = 30

func TestService_Recommendations(t *testing.T) {
	tests := []struct {
		name       string
		recom      recommender
		total int
		sds        []refind.Seed
		wantTracks []refind.Track
		wantErr    error
	}{
		{
			name: "Valid data, valid total, valid seeds, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  nil,
			},
			total: testTotal,
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
			name: "Valid data, valid total, valid seeds, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  testErrNoData,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr:    testErrNoData,
		},
		{
			name: "Valid data, valid total, no seeds, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  nil,
			},
			total: testTotal,
			sds:        nil,
			wantTracks: nil,
			wantErr:    errSeedsMissing,
		},
		{
			name: "Valid data, valid total, no seeds, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  testErrNoData,
			},
			total: testTotal,
			sds:        nil,
			wantTracks: nil,
			wantErr:    errSeedsMissing,
		},
		{
			name: "Valid data, valid total, invalid seed ID, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  nil,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr:    errSeedID,
		},
		{
			name: "Valid data, valid total, invalid seed ID, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  testErrNoData,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr:    errSeedID,
		},
		{
			name: "Valid data, valid total, invalid seed category, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  nil,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr:    errSeedCategory,
		},
		{
			name: "Valid data, valid total, invalid seed category, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  testErrNoData,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr:    errSeedCategory,
		},
		{
			name: "Valid data, invalid total, valid seeds, nil error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  nil,
			},
			total: 0,
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr:    errRangeInvalid,
		},
		{
			name: "Valid data, invalid total, valid seeds, error",
			recom: fakeRecommender{
				file: testFileRecommendations,
				err:  testErrNoData,
			},
			total: 0,
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr:    errRangeInvalid,
		},
		{
			name: "No data, valid total, valid seeds, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  nil,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr:    errDataInvalid,
		},
		{
			name: "No data, valid total, valid seeds, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.ArtistSeed, ID: "4NHQUGzhtTLFvgF5SZesLK"},
				{Category: refind.TrackSeed, ID: "0c6xIDDpzE81m2q797ordA"},
				{Category: refind.GenreSeed, ID: "classical"},
				{Category: refind.GenreSeed, ID: "country"},
			},
			wantTracks: nil,
			wantErr:    testErrNoData,
		},
		{
			name: "No data, valid total, no seeds, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  nil,
			},
			total: testTotal,
			sds:        nil,
			wantTracks: nil,
			wantErr:    errSeedsMissing,
		},
		{
			name: "No data, valid total, no seeds, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			total: testTotal,
			sds:        nil,
			wantTracks: nil,
			wantErr:    errSeedsMissing,
		},
		{
			name: "No data, valid total, invalid seed ID, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  nil,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr:    errSeedID,
		},
		{
			name: "No data, valid total, invalid seed ID, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: refind.TrackSeed, ID: ""},
			},
			wantTracks: nil,
			wantErr:    errSeedID,
		},
		{
			name: "No data, valid total, invalid seed category, nil error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  nil,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr:    errSeedCategory,
		},
		{
			name: "No data, valid total, invalid seed category, error",
			recom: fakeRecommender{
				file: testFileEmpty,
				err:  testErrNoData,
			},
			total: testTotal,
			sds: []refind.Seed{
				{Category: -99, ID: "some_id"},
			},
			wantTracks: nil,
			wantErr:    errSeedCategory,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := service{recom: test.recom}

			got, err := serv.Recommendations(test.total, test.sds)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantTracks) {
				t.Errorf("\ngot:  <%v>, \nwant: <%v>", got, test.wantTracks)
			}
		})
	}
}

type fakePlaylister struct {
	userFile     string
	userErr      error
	playlistFile string
	playlistErr  error
	addTracksErr error
}

func (f fakePlaylister) AddTracksToPlaylist(spotify.ID, ...spotify.ID) (string, error) {
	return "", f.addTracksErr
}

func (f fakePlaylister) CreatePlaylistForUser(string, string, string, bool) (*spotify.FullPlaylist, error) {
	b, err := ioutil.ReadFile(f.playlistFile)
	if err != nil {
		return nil, err
	}

	var play *spotify.FullPlaylist
	if err := json.Unmarshal(b, &play); err != nil {
		return nil, f.playlistErr
	}

	return play, f.playlistErr
}

func (f fakePlaylister) CurrentUser() (*spotify.PrivateUser, error) {
	b, err := ioutil.ReadFile(f.userFile)
	if err != nil {
		return nil, err
	}

	var u *spotify.PrivateUser
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, f.userErr
	}

	return u, f.userErr
}

var testFullPlaylist = &spotify.FullPlaylist{
	SimplePlaylist: spotify.SimplePlaylist{
		Collaborative: false,
		ExternalURLs: map[string]string{
			"spotify": "http://open.spotify.com/user/someone/playlist/7I6yjOAxMq4qsvzgqxw7aU",
		},
		Endpoint: "https://api.spotify.com/v1/users/someone/playlists/7I6yjOAxMq4qsvzgqxw7aU?fields=fields=href,name,owner(!href,external_urls),tracks.items(added_by.id,track(name,href,album(name,href)))",
		ID: "7I6yjOAxMq4qsvzgqxw7aU",
		Images: []spotify.Image{
			{Height: 640, Width: 640, URL: "https://i.scdn.co/image/449156e8f2458c247ea9a668e498c598b159f12f"},
		},
		Name: "O.A.R. — The Rockville LP",
		Owner: spotify.User{
			ExternalURLs: map[string]string{
				"spotify": "http://open.spotify.com/user/someone",
			},
			Endpoint: "https://api.spotify.com/v1/users/someone",
			ID: "someone",
			URI: "spotify:user:someone",
		},
		IsPublic: true,
		SnapshotID: "Yo+BthwRfySLE498r55BaKSJNw0/3ZDUzVYBcRxVtMReZ3joqyhIlBoMJKif2OWJ",
		URI: "spotify:user:someone:playlist:7I6yjOAxMq4qsvzgqxw7aU",
	},
	Followers: spotify.Followers{},
	Tracks: spotify.PlaylistTrackPage{
		Tracks: []spotify.PlaylistTrack{
			{
				AddedAt: "2014-06-10T11:23:19Z",
				AddedBy: spotify.User{
					ExternalURLs: map[string]string{
						"spotify": "http://open.spotify.com/user/someone",
					},
					Endpoint: "https://api.spotify.com/v1/users/someone",
					ID: "someone",
					URI: "spotify:user:someone",
				},
				Track: spotify.FullTrack{
					SimpleTrack: spotify.SimpleTrack{
						Artists: []spotify.SimpleArtist{
							{
								ExternalURLs: map[string]string{
									"spotify": "https://open.spotify.com/artist/1Cq0LAHFfvUTBEtMPXUidI",
								},
								Endpoint: "https://api.spotify.com/v1/artists/1Cq0LAHFfvUTBEtMPXUidI",
								ID: "1Cq0LAHFfvUTBEtMPXUidI",
								Name: "O.A.R.",
								URI: "spotify:artist:1Cq0LAHFfvUTBEtMPXUidI",
							},
						},
						Endpoint: "https://api.spotify.com/v1/tracks/1l7E6PxXL78DscO7YEDSBc",
						ID: "1l7E6PxXL78DscO7YEDSBc",
						Name: "We'll Pick Up Where We Left Off",
						PreviewURL: "https://p.scdn.co/mp3-preview/e107d8f17f036868d8389de406bae057bb609afa",
						TrackNumber: 2,
						URI: "spotify:track:1l7E6PxXL78DscO7YEDSBc",
					},
					Popularity: 43,
				},
			},
		},
	},
}

func TestService_Playlist(t *testing.T) {
	tests := []struct {
		name     string
		play     playlister
		tracks []refind.Track
		wantPlaylist *spotify.FullPlaylist
		wantErr  error
	}{
		{
			name: "Valid user with nil error, valid playlist with nil error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: testFullPlaylist,
			wantErr:      nil,
		},
		{
			name: "Valid user with error, valid playlist with nil error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: testErrNoData,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr: testErrNoData,
		},
		{
			name: "No user with nil error, valid playlist with nil error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileEmpty,
				userErr: nil,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr: errDataInvalid,
		},
		{
			name: "No user with error, valid playlist with nil error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileEmpty,
				userErr: testErrNoData,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr: testErrNoData,
		},
		{
			name: "Valid user with nil error, valid playlist with error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileCreatePlaylist,
				playlistErr: testErrNoData,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr:      testErrNoData,
		},
		{
			name: "Valid user with nil error, no playlist with nil error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileEmpty,
				playlistErr: nil,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr: errDataInvalid,
		},
		{
			name: "Valid user with nil error, no playlist with error, valid tracks with nil error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileEmpty,
				playlistErr: testErrNoData,
				addTracksErr: nil,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr: testErrNoData,
		},
		{
			name: "Valid user with nil error, valid playlist with nil error, valid tracks with error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: testErrNoData,
			},
			tracks: []refind.Track{
				{ID: "6qK7CuehGu2DVwL8UgaEhV", Name: "Days - Remastered", Artist: refind.Artist{ID: "0S7Zur2g8YhqlzqtlYStli", Name: "Television"}},
				{ID: "2kL584Ddb8dVjAbga456kZ", Name: "Bells Bleed & Bloom", Artist: refind.Artist{ID: "4K7elTMrmeEYTE9w1zGP5e", Name: "ef"}},
				{ID: "6XGLiFTNkatlSjGimT0tGU", Name: "Omens And Portents 1: The Driver", Artist: refind.Artist{ID: "4mTFQE6aiehScgvreB9llC", Name: "Earth"}},
			},
			wantPlaylist: nil,
			wantErr:      testErrNoData,
		},
		{
			name: "Valid user with nil error, valid playlist with nil error, no tracks with nil error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: nil,
			},
			tracks: nil,
			wantPlaylist: nil,
			wantErr:      errTracksMissing,
		},
		{
			name: "Valid user with nil error, valid playlist with nil error, no tracks with error",
			play: fakePlaylister{
				userFile: testFileCurrentUser,
				userErr: nil,
				playlistFile: testFileCreatePlaylist,
				playlistErr: nil,
				addTracksErr: testErrNoData,
			},
			tracks: nil,
			wantPlaylist: nil,
			wantErr:      errTracksMissing,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := service{play: test.play}

			got, err := serv.Playlist(test.name, "info", test.tracks)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(got, test.wantPlaylist) {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantPlaylist)
			}
		})
	}
}
