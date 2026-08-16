package atbashcipher

import (
    "strings"
    "unicode"
)

func Atbash(s string) string {
	encrypted := make([]rune, 0, len(s))
    letters := 0
    for _, c := range strings.ToLower(s) {
        if !unicode.IsDigit(c) && !unicode.IsLower(c) {
            continue
        }
        
        if letters == 5 {
        	encrypted = append(encrypted, ' ')
            letters = 0
        }
        
        if unicode.IsDigit(c) {
         	encrypted = append(encrypted, c)   
        } else if unicode.IsLower(c) {
            encrypted = append(encrypted, 'a'+('z'-c))
        }
        letters++
        
    }
    return string(encrypted)
}
