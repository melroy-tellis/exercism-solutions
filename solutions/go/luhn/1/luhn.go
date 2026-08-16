package luhn

import (
    "strings"
    "unicode"
)


func Valid(id string) bool {
    stripped := []rune(strings.ReplaceAll(id, " ", ""))
    if len(stripped) < 2 {
        return false
    }
    
    sum := 0
	for i := 1; i <= len(stripped); i++ {
        if !unicode.IsDigit(stripped[len(stripped)-i]) {
            return false
        }
    	digit := int(stripped[len(stripped)-i]-'0')
        if i%2 == 0 {
            digit += digit
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
    }
    return sum%10 == 0
}
