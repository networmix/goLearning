//Unpack strings packed with naive Run-length encoding (RLE)
//  Examples:
//* "a4bc2d5e" => "aaaabccddddde"
//* "abcd" => "abcd"
//* "45" => "" <- Incorrect string yields empty string
//* `qwe\4\5` => `qwe45` (*) <- escape handling
//* `qwe\45` => `qwe44444` (*) <- escape handling
//* `qwe\\5` => `qwe\\\\\` (*) <- escape handling
package unpackrle

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type StringRLE struct {
	encodedStr string
	decodedStr string
}

func (s *StringRLE) UnpackString(encodedStr string) string {
	s.encodedStr = encodedStr
	for substr := s.getNextSubstr(); len(substr[0]) > 0; substr = s.getNextSubstr() {
		if err := s.unpackSubstring(substr[0]); err != nil {
			return ""
		}
	}
	return s.decodedStr
}

func (s *StringRLE) getNextSubstr() []string {
	if len(s.encodedStr) == 0 {
		return []string{""}
	}
	re := regexp.MustCompile(`^[a-zA-Z\\]+[0-9]*`)

	substring := re.FindStringSubmatch(s.encodedStr)
	if len(substring) == 0 {
		return []string{""}
	}
	s.encodedStr = s.encodedStr[len(substring[0]):]
	return substring
}

func (s *StringRLE) unpackSubstring(str string) error {
	re := regexp.MustCompile(`^(?P<letters>(?:[a-zA-Z]|\\\\)*(?:\\[0-9])*)(?P<number>[0-9]*$)`)

	match := re.FindStringSubmatch(str)
	switch {
	// Matched "letters" and the "number" part is empty
	case len(match[1]) > 0 && len(match[2]) == 0:
		// Adding the normalised "letters" part to the decodedStr
		s.decodedStr += s.normaliseStr(match[1])

	// Matched "letters" and the "number" part is not empty
	case len(match[1]) > 0 && len(match[2]) > 0:
		num, _ := strconv.Atoi(match[2]) // converting the "number" part into int
		s.decodedStr += s.normaliseStr(match[1][:len(match[1])-1]) + strings.Repeat(match[1][len(match[1])-1:], num)

	// No match or the "letters" part is empty
	default:
		return errors.New("no match or the 'letters' part is empty")
	}
	return nil
}

// normaliseStr is a private method handling escape symbols in a substring
func (s *StringRLE) normaliseStr(str string) string {
	// If it is a single escape symbol, removing it
	re1 := regexp.MustCompile(`^([a-zA-Z0-9])?\\((?:\\\\)*[a-zA-Z0-9])|([a-zA-Z0-9](?:\\\\)*)\\([a-zA-Z0-9])?$`)
	str = re1.ReplaceAllString(str, "$1$2$3$4")

	// Replacing all occurrences of double escape symbols with a single escape symbol
	re2 := regexp.MustCompile(`\\{2}`)
	return re2.ReplaceAllString(str, `\`)
}
