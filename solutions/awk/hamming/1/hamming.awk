BEGIN {
    hamming_distance = 0;
}
{
    if (NR == 1) first = $0;
    if (NR == 2) {
        second = $0;
        if (length(first) != length(second)) {
            hamming_distance = "strands must be of equal length"
            exit 1;
        }
        for (i = 1; i <= length; ++i) {
            if (substr(first, i, 1) != substr(second, i, 1)) {
                ++hamming_distance;
            }
        }
    }
    
}
END {
    print hamming_distance;
}
