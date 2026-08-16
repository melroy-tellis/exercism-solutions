package pangram

import (
    "strings"
)

func IsPangram(input string) bool {
    containedLetters := make(map[rune]bool)
    lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
    for _, r := range strings.ToLower(input) {
        if strings.ContainsRune(lowercaseLetters, r) {
            containedLetters[r] = true
        }
    }
    return len(containedLetters) == len(lowercaseLetters)
}
