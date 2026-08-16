package house

func Verse(v int) string {
    fragments := []string{
        "house that Jack built",
        "malt\nthat lay in",
        "rat\nthat ate",
        "cat\nthat killed",
        "dog\nthat worried",
        "cow with the crumpled horn\nthat tossed",
        "maiden all forlorn\nthat milked",
        "man all tattered and torn\nthat kissed",
        "priest all shaven and shorn\nthat married",
        "rooster that crowed in the morn\nthat woke",
        "farmer sowing his corn\nthat kept",
        "horse and the hound and the horn\nthat belonged to",
    }
    if (v == 1) {
        return "This is the " + fragments[0] + ".";
    }
    return "This is the " + fragments[v - 1] + " the " + Verse(v - 1)[12:]
}

func Song() string {
	song := ""
    for v := 1; v <= 12; v++ {
        song += Verse(v)
        if v < 12 {
            song += "\n\n"
        }
    }
    return song
}
