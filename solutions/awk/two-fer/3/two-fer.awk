BEGIN {
  name = "you"
}
length($0) > 0 { name = $0 }
END {
    printf "One for %s, one for me.", name
}