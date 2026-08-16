BEGIN {
    reverse = ""
}
{ 
    for (i = length; i > 0; i--) reverse=reverse substr($0,i,1);
}
END {
    print reverse
}