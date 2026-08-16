BEGIN {
    score = 0;
    word = "";
}
{
    for (i = 1; i <= length; i++) {
        letter = toupper(substr($0, i, 1))
        switch(letter) {
            case "A":
            case "E":
            case "I":
            case "O":
            case "U":
            case "L":
            case "N":
            case "R":
            case "S":
            case "T":
                ++score
                break
            case "D":
            case "G":
             score = score + 2
             break
            case "B":
            case "C":
            case "M":
            case "P":
                score = score + 3
                break
            case "F":
            case "H":
            case "V":
            case "W":
            case "Y":
                score = score + 4
                break;
            case "K":
                score = score + 5
                break
            case "J":
            case "X":
                score = score + 8
                break
            case "Q":
            case "Z":
                score = score + 10
                break;
        }
        word = word letter
        
    }
}
END {
    printf "%s,%d", word, score
}