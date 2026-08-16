BEGIN {
    rna = ""
}
{ 
    for (i = 1; i <= length; ++i) {
        switch(substr($0,i, 1)) {
            case "C":
              comp_nucleotide = "G";
              break;
            case "G":
              comp_nucleotide = "C";
              break;
            case "T":
                comp_nucleotide = "A";
                break;
            case "A":
                comp_nucleotide = "U";
                break;
            default:
                rna = "Invalid nucleotide detected."
                exit 1;
        }
        rna = rna comp_nucleotide;
    }
        
}
END {
    print rna
}

