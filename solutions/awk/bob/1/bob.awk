END {
    if ($0 ~ /^[^[:lower:]]+$/ && $0 ~ /[[:upper:]]/) {
        if ($0 ~ /?[[:space:]]*$/) {
            print "Calm down, I know what I'm doing!"
        } else {
            print "Whoa, chill out!" 
        }
    } else if ($0 ~ /?[[:space:]]*$/) { 
        print "Sure." 
    } else if ($0 ~/^[[:space:]]*$/) {
        print "Fine. Be that way!"
    } else {
        print "Whatever."
    }
}