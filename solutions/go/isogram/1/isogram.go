package isogram

import "strings"

func IsIsogram(word string) bool {
    letterCount := make(map[rune]int)
    for _, r := range(strings.ToLower(word)) {
        if r != '-' && r != ' ' {
        	letterCount[r]++    
        }
    }
    for _, c := range(letterCount) {
        if c > 1 {
            return false
        }
    }
    return true
	
}
