BEGIN {
 FS="[[:space:]-]"
 acronym=""
}
{
    for (i = 1; i <= NF; ++i) {
        gsub(/[^a-zA-Z]/, "", $i)
        if (length($i) > 0) {
            acronym = acronym toupper(substr($i, 1, 1))
        }
    }
}
END {
    print acronym
}
