# These variables are initialized on the command line (using '-v'):
# - distance
@load "ordchr"
BEGIN {
    rotated = ""
}

{
    for (i = 1; i <= length; ++i) {
        char = substr($0, i, 1)
        if (char ~ /[[:alpha:]]/) {
            base = char ~ /[[:lower:]]/ ? "a" : "A"
            offset = (ord(char) - ord(base) + distance) % 26
            char = chr(ord(base) + offset)
        }
        rotated = rotated char
    }
    print rotated
}