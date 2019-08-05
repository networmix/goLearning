package wordfreqency

import (
    "reflect"
    "testing"
)

func Test_buildFreqDict(t *testing.T) {
    type args struct {
        text string
    }
    tests := []struct {
        name string
        args args
        want map[string]int
    }{
        {"Test1", args{"OneWord"}, map[string]int{"oneword": 1}},
        {"Test2", args{"Hello!\nWorld!"}, map[string]int{"hello": 1, "world": 1}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := buildFreqDict(tt.args.text); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetFreqDict() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_normaliseText(t *testing.T) {
    type args struct {
        text string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {"Test1", args{"Hello world!"}, "hello world "},
        {"Test2", args{"Hello!\nWorld!"}, "hello  world "},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := normaliseText(tt.args.text); got != tt.want {
                t.Errorf("normaliseText() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGetTopWords(t *testing.T) {
    type args struct {
        text string
        n    int
    }
    tests := []struct {
        name string
        args args
        want KeyValueList
    }{
        {"Test1", args{"", 10}, KeyValueList{}},
        {"Test2", args{"Hello!\n World, my bright bright world!", 3}, KeyValueList{
            KeyValue{"bright", 2},
            KeyValue{"world", 2},
            KeyValue{"hello", 1},
        }},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := GetTopWords(tt.args.text, tt.args.n); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetTopWords() = %v, want %v", got, tt.want)
            }
        })
    }
}
