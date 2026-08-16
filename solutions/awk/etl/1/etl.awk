BEGIN {
    FS=":"
    alphas = "abcdefghijklmnopqrstuvwxyz"
}
{
    gsub(/[[:space:]"]/, "", $2)
    split($2, letters, ",")
    for (i in letters) {
        value[tolower(letters[i])] = $1
    }
    
}
END {
    for (i=1; i <= length(alphas); ++i) {
        letter = substr(alphas, i, 1) 
        if (letter in value) {
            printf "%s,%s\n", letter, value[letter]
        
        }
    }
}
