package urlshortener

import (
	"testing"
)

func Test_checkIfURLisValid(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ValidURL", args{"https://gobyexample.com"}, true},
		{"InvalidURL", args{"gobyexample.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfURLisValid(tt.args.rawURL); got != tt.want {
				t.Errorf("checkIfURLisValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortener_Shorten(t *testing.T) {
	type fields struct {
		urls map[string]string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"ValidURL", fields{make(map[string]string)}, args{"https://gobyexample.com"}, "703a8b3ff83e3ea1"},
		{"InvalidURL", fields{make(map[string]string)}, args{"gobyexample.com"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shortener{
				urls: tt.fields.urls,
			}
			if got := s.Shorten(tt.args.url); got != tt.want {
				t.Errorf("Shortener.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
