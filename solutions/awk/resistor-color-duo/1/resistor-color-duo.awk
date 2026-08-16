BEGIN {
    RS="[[:space:]]"
    resistance = 0
}

NR > 2 { exit; }
{ 
    resistance = resistance * 10
    value = -1;
    switch($0) {
        case "black":
            value = 0;
            break;
        case "brown":
            value = 1;
            break;
        case "red":
            value = 2;
            break;
        case "orange":
            value = 3;
            break;
        case "yellow":
            value = 4;
            break;
        case "green":
            value = 5;
            break;
        case "blue":
            value = 6;
            break;
        case "violet":
            value = 7;
            break;
        case "grey":
            value = 8;
            break;
        case "white":
            value = 9;
            break;
        default:
            print "invalid color"
            exit 1;
    }
    resistance = resistance + value
}

NR == 2 { print resistance }