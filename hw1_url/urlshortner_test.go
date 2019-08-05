package urlshortener

import (
    "encoding/hex"
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

// Based on a well-known collision example https://en.wikipedia.org/wiki/MD5#Collision_vulnerabilities
func TestMD5Shortener_AddWithCollision(t *testing.T) {
    s := &MD5Shortener{
        urls: make(map[string]string),
    }

    testURL1, _ := hex.DecodeString("d131dd02c5e6eec4693d9a0698aff95c2fcab58712467eab4004583eb8fb7f8955ad340609f4b30283e488832571415a085125e8f7cdc99fd91dbdf280373c5bd8823e3156348f5bae6dacd436c919c6dd53e2b487da03fd02396306d248cda0e99f33420f577ee8ce54b67080a80d1ec69821bcb6a8839396f9652b6ff72a70")
    testURL2, _ := hex.DecodeString("d131dd02c5e6eec4693d9a0698aff95c2fcab50712467eab4004583eb8fb7f8955ad340609f4b30283e4888325f1415a085125e8f7cdc99fd91dbd7280373c5bd8823e3156348f5bae6dacd436c919c6dd53e23487da03fd02396306d248cda0e99f33420f577ee8ce54b67080280d1ec69821bcb6a8839396f965ab6ff72a70")
    shortURL1 := s.addURL(string(testURL1))
    shortURL2 := s.addURL(string(testURL2))
    t.Logf("'%s'", shortURL1)
    t.Logf("'%s'", shortURL2)
    if got := s.Resolve(shortURL1); got != string(testURL1) {
        t.Errorf("URLShortener.Resolve() = %v, want %v", got, string(testURL1))
    }
    if got := s.Resolve(shortURL2); got != string(testURL2) {
        t.Errorf("URLShortener.Resolve() = %v, want %v", got, string(testURL2))
    }
}
