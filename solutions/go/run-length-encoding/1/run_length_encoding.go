package encode

import (
    "strings"
    "fmt"
)

func RunLengthEncode(input string) string {
    if (len(input)) == 0 {
        return input
    }
    var sb strings.Builder
    runes := []rune(input)
    currRune, count := runes[0], 1
    for i := 1; i <= len(runes); i++ {
        if i == len(runes) || runes[i] != currRune {
            if count > 1 {
            	sb.WriteString(fmt.Sprint(count))    
            }
            sb.WriteRune(currRune)
            count = 1
        } else {
            count ++
        }
        if i < len(runes) {
            currRune = runes[i]
        }
    }
    return sb.String()
}

func RunLengthDecode(input string) string {
	count := 0
    var sb strings.Builder
    for _, c := range(input) {
        if c >= '0' && c <= '9' {
            count = count * 10 + int(c) - int('0')
        } else {
            if count == 0 {
                count++
            }
            for count > 0 {
                sb.WriteRune(c)
                count--
            }
        }
        
    }
    return sb.String()
}
