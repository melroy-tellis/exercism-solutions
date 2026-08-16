package microblog

func Truncate(phrase string) string {
	truncated := make([]rune, 0, 5)
    runes := 0
    for _, r := range(phrase) {
        if runes == 5 {
            break
        }
        truncated = append(truncated, r)
        runes++
    }
    return string(truncated)
    
}
