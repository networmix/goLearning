package urlshortener

import (
	"crypto/md5"
	"fmt"
	"net/url"
)

// Shortener is a simple URL shortner.
type Shortener struct {
	urls map[string]string
}

func (s *Shortener) Shorten(url string) string {
	if checkIfURLisValid(url) {
		urlBytes := []byte(url)
		urlHash := md5.Sum(urlBytes)
		shortURL := fmt.Sprintf("%x", urlHash[:8])
		s.urls[url] = shortURL
		return shortURL
	}
	return ""
}

func checkIfURLisValid(rawURL string) bool {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}
	return true
}
