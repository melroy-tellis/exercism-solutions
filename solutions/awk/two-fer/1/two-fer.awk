END {
    name = "you"
    { if (length($0) > 0) name = $0 }
    printf "One for %s, one for me.", name
}