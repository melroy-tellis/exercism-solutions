package wordcount
import (
    "strings"
    "unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	frequency := make(Frequency)
    word := []rune{}
    for i, c := range strings.ToLower(phrase) {
        if unicode.IsLower(c) || unicode.IsDigit(c) || c == '\'' {
            word = append(word, c)
            if i < len(phrase)-1 {
                continue
            }
        }
        trimmedWord := strings.Trim(string(word), "'")
        if len(trimmedWord) > 0 {
            frequency[trimmedWord]++    
        }
        word = word[:0]
    }
    return frequency
}
