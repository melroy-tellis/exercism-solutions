BEGIN {
    protein["AUG"] = "Methionine"
    protein["UUU"] = protein["UUC"] = "Phenylalanine"
    protein["UUA"] = protein["UUG"] = "Leucine"
    protein["UCU"] = protein["UCC"] = protein["UCA"] = protein["UCG"] = "Serine"
    protein["UAU"] = protein["UAC"] = "Tyrosine"
    protein["UGU"] = protein["UGC"] = "Cysteine"
    protein["UGG"] = "Tryptophan"
    protein["UAA"] = protein["UAG"] = protein["UGA"] = "STOP"
    peptides = ""
}
{
     for (i = 1 ; i <= length; i=i+3) {
         codon = substr($0, i, 3)
         if (i + 2 > length || !(codon in protein)) {
             print "Invalid codon"
             exit 1
         }
         if (protein[codon] == "STOP") {
             break
         }
         if (length(peptides) > 0) {
             peptides = peptides " "
         }
         peptides = peptides protein[codon]
     }
     print peptides
 }