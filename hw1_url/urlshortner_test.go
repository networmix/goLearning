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

func TestMD5Shortener_Shorten(t *testing.T) {
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
			s := &MD5Shortener{
				urls: tt.fields.urls,
			}
			if got := s.Shorten(tt.args.url); got != tt.want {
				t.Errorf("Shortener.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5Shortener_Resolve(t *testing.T) {
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
		{"CachedURL", fields{map[string]string{"703a8b3ff83e3ea1": "https://gobyexample.com"}}, args{"703a8b3ff83e3ea1"}, "https://gobyexample.com"},
		{"NotCachedURL", fields{make(map[string]string)}, args{"NotCachedURL"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MD5Shortener{
				urls: tt.fields.urls,
			}
			if got := s.Resolve(tt.args.url); got != tt.want {
				t.Errorf("URLShortener.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBASE64Shortener_Shorten(t *testing.T) {
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
		{"ValidURL", fields{make(map[string]string)}, args{"https://gobyexample.com"}, "AA"},
		{"InvalidURL", fields{make(map[string]string)}, args{"gobyexample.com"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BASE64Shortener{
				urls: tt.fields.urls,
			}
			if got := s.Shorten(tt.args.url); got != tt.want {
				t.Errorf("Shortener.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBASE64Shortener_Resolve(t *testing.T) {
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
		{"CachedURL", fields{map[string]string{"AA": "https://gobyexample.com"}}, args{"AA"}, "https://gobyexample.com"},
		{"NotCachedURL", fields{make(map[string]string)}, args{"NotCachedURL"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BASE64Shortener{
				urls: tt.fields.urls,
			}
			if got := s.Resolve(tt.args.url); got != tt.want {
				t.Errorf("URLShortener.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortener_interface(t *testing.T) {
	md5shortener := &MD5Shortener{
		urls: make(map[string]string),
	}
	base64shortener := &BASE64Shortener{
		urls: make(map[string]string),
	}
	shorteners := []Shortener{md5shortener, base64shortener}
	testURL := "https://gobyexample.com"
	for _, s := range shorteners {
		shortURL := s.Shorten(testURL)
		if got := s.Resolve(shortURL); got != testURL {
			t.Errorf("URLShortener.Resolve() = %v, want %v", got, testURL)
		}
	}
}
