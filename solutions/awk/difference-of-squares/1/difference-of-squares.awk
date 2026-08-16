BEGIN {
    FS=","
}
function square_of_sum(n) {
    sum = (n * (n + 1))/2
    return  sum * sum;
}
function sum_of_squares(n) {
    return (n * (n + 1) * (2 * n + 1))/6;
}
END {
    switch($1) {
        case "square_of_sum":
            print square_of_sum($2)
            break
        case "sum_of_squares":
            print sum_of_squares($2)
            break
        default:
            print square_of_sum($2) - sum_of_squares($2)
    }
}