END {
    for(i = 1; i <= length; ++i) {
        char = substr($0, i, 1); 
        if (char ~ /[[:alpha:]]/) {
            ++count[toupper(char)]
        }
    }
    print length(count) == 26 ? "true" : "false"

}
