package wordfreqency

import (
    "regexp"
    "sort"
    "strings"
)

type KeyValue struct {
    key   string
    value int
}

type KeyValueList []KeyValue

func (kvlist KeyValueList) Len() int {
    return len(kvlist)
}

func (kvlist KeyValueList) Swap(i, j int) {
    kvlist[i], kvlist[j] = kvlist[j], kvlist[i]
}

func (kvlist KeyValueList) Less(i, j int) bool {
    if kvlist[i].value > kvlist[j].value {
        return true
    }

    if kvlist[i].value == kvlist[j].value {
        return sort.StringsAreSorted([]string{kvlist[i].key, kvlist[j].key})
    }
    return false
}

// GetTopWords takes a text and n as input. It returns a KeyValueList with n most frequent words.
func GetTopWords(text string, n int) KeyValueList {
    wordFrequencies := buildFreqDict(text)
    kvlist := make(KeyValueList, len(wordFrequencies))
    for key, value := range wordFrequencies {
        kvlist = append(kvlist, KeyValue{key, value})
    }
    sort.Sort(KeyValueList(kvlist))
    if len(kvlist) < n {
        return kvlist
    }
    return kvlist[:n]
}

// buildFreqDict builds a word frequency dictionary for a given text.
func buildFreqDict(text string) map[string]int {
    wordFrequencies := make(map[string]int)
    for _, word := range strings.Fields(normaliseText(text)) {
        wordFrequencies[word] += 1
    }
    return wordFrequencies
}

// normaliseWords replaces any newline and punctuation symbols in a given text with spaces
// and makes all symbols lower case.
func normaliseText(text string) string {
    re := regexp.MustCompile("[[:punct:]\n\t]")
    text = re.ReplaceAllString(text, " ")
    return strings.ToLower(text)
}
