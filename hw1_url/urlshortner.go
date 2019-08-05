package urlshortener

import (
    "crypto/md5"
    "encoding/base64"
    "fmt"
    "net/url"
)

// Shortener is the interface of a URL shortener.
// Shorten - accepts a string and, if that string is a valid URL, it returns a shortened URL.
// Resolve - accepts a string and, if that string corresponds to a cached URL, it returns the previously memorised URL.
type Shortener interface {
    Shorten(url string) string
    Resolve(url string) string
}

// MD5Shortener is a simple URL shortner which uses MD5 to shorten supplied URLs.
type MD5Shortener struct {
    urls map[string]string
}

// Shorten method accepts a string URL and, if the string is a valid URL, it returns a shortened URL.
// If URL is not valid, the method returns an empty string.
func (s *MD5Shortener) Shorten(url string) string {
    if checkIfURLisValid(url) {
        shortURL := s.addURL(url)
        return shortURL
    }
    return ""
}

// addURL adds a given URL into the associative array with the key
// calculated as an MD5 hash.
// Cryptographic attacks do not apply in our use-case, hence collisions are highly improbable.
// We need to accumulate about 2^32 16-hexadecimal digits long 'shortURLs' in the map to face approximately
// 50% chance of collision. Read more here https://en.wikipedia.org/wiki/Birthday_problem.
// Still, to resolve this quite unlikely collision, we will try increasing the length of the shortURL by 1 hex digit
// up to the full length of MD5 vector - 128 bits or 32 hex digits.
// For those interested in cryptographic attacks against MD5, please see the link below :)
// https://www.mscs.dal.ca/~selinger/md5collision/
func (s *MD5Shortener) addURL(url string) string {
    urlBytes := []byte(url)
    urlHash := fmt.Sprintf("%x", md5.Sum(urlBytes))

    for hashLen := 16; hashLen <= 32; hashLen++ {
        shortURL := urlHash[:hashLen]

        if val, inCache := s.urls[shortURL]; inCache {
            if val == url {
                return shortURL
            } else if val != url && hashLen < 32 {
                continue
            }
        } else {
            s.urls[shortURL] = url
            return shortURL
        }
    }
    return ""
}

// Resolve expects a string with a short url and returns corresponding cached long url (if any) or an empty string.
func (s *MD5Shortener) Resolve(url string) string {
    originalURL := s.urls[url]
    return originalURL
}

// BASE64Shortener is a simple URL shortner which uses BASE64 to shorten supplied URLs.
type BASE64Shortener struct {
    urls map[string]string
}

func (s *BASE64Shortener) getNextID() int {
    return len(s.urls)
}

// Shorten method accepts accepts a string URL and, if the string is a valid URL, it returns a shortened URL.
// If URL is not valid, the method returns an empty string.
func (s *BASE64Shortener) Shorten(url string) string {
    if checkIfURLisValid(url) {
        idBytes := []byte(string(s.getNextID()))
        // From RFC4648:
        // The encoding process represents 24-bit groups of input bits as output
        // strings of 4 encoded characters.  Proceeding from left to right, a
        // 24-bit input group is formed by concatenating 3 8-bit input groups.
        // These 24 bits are then treated as 4 concatenated 6-bit groups, each
        // of which is translated into a single character in the base 64
        // alphabet.
        // Please look at https://tools.ietf.org/html/rfc4648#page-7 for the full reference.
        urlEncodedIndex := base64.RawURLEncoding.EncodeToString(idBytes)
        s.urls[urlEncodedIndex] = url
        return urlEncodedIndex
    }
    return ""
}

// Resolve expects a string with a short url and returns corresponding cached long url (if any) or an empty string.
func (s *BASE64Shortener) Resolve(url string) string {
    originalURL := s.urls[url]
    return originalURL
}

// checkIfURLisValid checks the validity of the URL.
func checkIfURLisValid(rawURL string) bool {
    _, err := url.ParseRequestURI(rawURL)
    if err != nil {
        return false
    }
    return true
}
