package scry

import (
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	tests := []struct {
		name string
		prev []Artist
		want map[string]Artist
	}{
		{
			"Nil artist slice",
			nil,
			nil,
		},
		{
			"Empty artist slice",
			[]Artist{},
			nil,
		},
		{
			"Single artist slice",
			[]Artist{
				{ID: "1", Name: "one"},
			},
			map[string]Artist{
				"one": {ID: "1", Name: "one"},
			},
		},
		{
			"Multiple artists slice",
			[]Artist{
				{ID: "1", Name: "one"},
				{ID: "2", Name: "two"},
				{ID: "3", Name: "three"},
				{ID: "4", Name: "four"},
				{ID: "5", Name: "five"},
				{ID: "6", Name: "six"},
			},
			map[string]Artist{
				"one": {ID: "1", Name: "one"},
				"two": {ID: "2", Name: "two"},
				"three": {ID: "3", Name: "three"},
				"four": {ID: "4", Name: "four"},
				"five": {ID: "5", Name: "five"},
				"six": {ID: "6", Name: "six"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := toMap(test.prev)
			if !reflect.DeepEqual(got, test.want){
				t.Errorf("got: <%v>, want: <%v>", got, test.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		prev []Track
		rmv map[string]Artist
		want []Track
	}{
		{
			"Nil tracks and nil artist map",
			nil,
			nil,
			nil,
		},
		{
			"Multiple tracks and nil artist map",
			[]Track{
				{ID: "1", Name: "foo", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "2", Name: "bar", Artist: Artist{ID: "12", Name: "garply"}},
				{ID: "3", Name: "baz", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "4", Name: "qux", Artist: Artist{ID: "14", Name: "fred"}},
				{ID: "5", Name: "quux", Artist: Artist{ID: "14", Name: "fred"}},
			},
			nil,
			[]Track{
				{ID: "1", Name: "foo", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "2", Name: "bar", Artist: Artist{ID: "12", Name: "garply"}},
				{ID: "3", Name: "baz", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "4", Name: "qux", Artist: Artist{ID: "14", Name: "fred"}},
				{ID: "5", Name: "quux", Artist: Artist{ID: "14", Name: "fred"}},
			},
		},
		{
			"Nil tracks and multiple artist map",
			nil,
			map[string]Artist{
				"grault": {ID: "11", Name: "grault"},
				"fred": {ID: "14", Name: "fred"},
			},
			nil,
		},
		{
			"Multiple tracks and single artist map",
			[]Track{
				{ID: "1", Name: "foo", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "2", Name: "bar", Artist: Artist{ID: "12", Name: "garply"}},
				{ID: "3", Name: "baz", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "4", Name: "qux", Artist: Artist{ID: "14", Name: "fred"}},
				{ID: "5", Name: "quux", Artist: Artist{ID: "14", Name: "fred"}},
			},
			map[string]Artist{
				"garply": {ID: "12", Name: "garply"},
			},
			[]Track{
				{ID: "1", Name: "foo", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "3", Name: "baz", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "4", Name: "qux", Artist: Artist{ID: "14", Name: "fred"}},
				{ID: "5", Name: "quux", Artist: Artist{ID: "14", Name: "fred"}},
			},
		},
		{
			"Multiple tracks with same artist and artist map with that same artist",
			[]Track{
				{ID: "1", Name: "foo", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "2", Name: "bar", Artist: Artist{ID: "12", Name: "garply"}},
				{ID: "3", Name: "baz", Artist: Artist{ID: "11", Name: "grault"}},
				{ID: "4", Name: "qux", Artist: Artist{ID: "14", Name: "fred"}},
				{ID: "5", Name: "quux", Artist: Artist{ID: "14", Name: "fred"}},
			},
			map[string]Artist{
				"grault": {ID: "11", Name: "grault"},
				"fred": {ID: "14", Name: "fred"},
			},
			[]Track{
				{ID: "2", Name: "bar", Artist: Artist{ID: "12", Name: "garply"}},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := filter(test.prev, test.rmv)

			if !reflect.DeepEqual(got, test.want){
				t.Errorf("got: <%v>, want: <%v>", got, test.want)
			}
		})
	}
}
