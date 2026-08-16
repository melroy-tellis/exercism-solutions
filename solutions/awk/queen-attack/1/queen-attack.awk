{
    for (i = 1 ; i <= 4; ++i) {
        if ($i < 0 || $i > 7) {
            print "invalid"
            exit 1
        }
    }
}
$1 == $3 && $2 == $4 { print "invalid"; exit 1 }

$1 == $3 || $2 == $4 { print "true"; exit }
{
    row_dist = $1 > $3 ? $1 - $3 : $3 - $1
    col_dist = $2 > $4 ? $2 - $4 : $4 - $2
    print row_dist == col_dist ? "true" : "false" 
}