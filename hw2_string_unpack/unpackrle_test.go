package unpackrle

import (
	"testing"
)

func TestStringRLE_normaliseStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Normalise1Left", args{`\ae4`}, `ae4`},
		{"Normalise1Right", args{`ae4\`}, `ae4`},
		{"Normalise1Center", args{`ae\4`}, `ae4`},
		{"Normalise2", args{`ae\\`}, `ae\`},
		{"Normalise3", args{`ae\\\`}, `ae\`},
		{"Normalise4", args{`ae\\\\`}, `ae\\`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringRLE{}
			if got := s.normaliseStr(tt.args.str); got != tt.want {
				t.Errorf("StringRLE.normaliseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringRLE_UnpackString(t *testing.T) {
	type args struct {
		encodedStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TrivialString1", args{"a"}, "a"},
		{"TrivialString2", args{"a4"}, "aaaa"},
		{"String1", args{"a4e2"}, "aaaaee"},
		{"String2", args{"qwea4e2qwe"}, "qweaaaaeeqwe"},
		{"String3", args{"a4bc2d5e"}, "aaaabccddddde"},
		{"EscString1", args{`qwe\4\5`}, "qwe45"},
		{"EscString2", args{`qwe\45`}, "qwe44444"},
		{"EscString3", args{`qwe\\5`}, `qwe\\\\\`},
		{"Xfail1", args{"45"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringRLE{}
			if got := s.UnpackString(tt.args.encodedStr); got != tt.want {
				t.Errorf("StringRLE.UnpackString() = %v, want %v", got, tt.want)
			}
		})
	}
}
