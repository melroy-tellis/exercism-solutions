# These variables are initialized on the command line (using '-v'):
# - type

BEGIN {
    is_of_type["isosceles"] = "false"
    is_of_type["equilateral"] = "false"
    is_of_type["scalene"] = "false"
}
{

    if (($1 == 0 || $2 == 0 || $3 == 0) || ($1 + $2 < $3) || ($2 + $3 < $1) || ($1 + $3 < $2)) {
        print "false"
        exit
    }
    if ($1 == $2 || $1 == $3 || $2 == $3) {
        is_of_type["isosceles"] = "true"
        if ($1 == $2 && $2 == $3) {
            is_of_type["equilateral"] = "true"
        }
    } else {
        is_of_type["scalene"] = "true"
    }
    print is_of_type[type]
}
