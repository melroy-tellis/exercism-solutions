BEGIN {
    FS=","
    split("eggs,peanuts,shellfish,strawberries,tomatoes,chocolate,pollen,cats", allergies, ",")
    allergic_to = "false"
    list = ""
}
END {
    for (i in allergies) {
        allergy = allergies[i]
        if (and(rshift($1, i - 1), 1)) {
            if ($2 == "allergic_to" && $3 == allergy) {
                allergic_to = "true"
                break
            }
            
            if (length(list) > 0) {
                list = list ","
            }
            list = list allergy
        }
    }
    if ($2 == "allergic_to") {
        print allergic_to
    }
    if ($2 == "list") {
        print list
    }
}