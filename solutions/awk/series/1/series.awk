# These variables are initialized on the command line (using '-v'):
# - len

BEGIN {
    if (len <= 0) {
        print "invalid length";
        exit 1;
    }
}

length == 0 { print "series cannot be empty"; exit 1; }
len > length { print "invalid length"; exit 1; }
{
    slices = substr($0, 1, len)
    for (i = 2; i <= length + 1 - len; ++i) {
        slices = slices " " substr($0, i, len)
    }
    print slices
}
