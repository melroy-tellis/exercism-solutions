BEGIN {
    leap = "false"
}
{
    if ($0 % 4 == 0 && ($0 % 100 != 0 || $0 % 400 == 0)) leap = "true"
}
END {
    print leap
}
