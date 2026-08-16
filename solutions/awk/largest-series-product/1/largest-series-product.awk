BEGIN {
    FS=","
}
function digits_product(n) {
    if (length(n) == 0) {
        return 0;
    }
    result = 1
    for (x = 1; x <= length(n); ++x) {
        result = result * substr(n, x, 1)
    }
    return result
}
END {
    if ($2 > length($1)) {
        print "span must be smaller than string length"
        exit 1
    }
    if ($2 < 0) {
        print "span must not be negative"
        exit 1
    }
    if ($1 !~ /^[[:digit:]]+$/ || $2 !~ /^[[:digit:]+]$/) {
        print "input must only contain digits"
        exit 1;
    }
    max_product = 0
    for (i=1; i <= length($1) - $2 + 1; ++i) {
        product = digits_product(substr($1, i, $2))
        if (product > max_product) {
            max_product = product
        }
    }
    print max_product
}
