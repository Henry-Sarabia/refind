package scry

import (
	"github.com/pkg/errors"
	"testing"
)

func TestAuthenticator(t *testing.T) {
	tests := []struct {
		name string
		uri string
		wantErr error
	}{
		{"Empty URI", "", errors.New("URI is blank")},
		{"Non-empty URI", "some-uri", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := authenticator(test.uri)
			if err != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", err, test.wantErr)
			}
		})
	}
}

