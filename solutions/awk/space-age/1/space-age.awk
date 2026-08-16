BEGIN {
    period_multiplier["Mercury"] = 0.2408467
    period_multiplier["Venus"] = 0.61519726
    period_multiplier["Earth"] = 1.0
    period_multiplier["Mars"] = 1.8808158
    period_multiplier["Jupiter"] = 11.862615
    period_multiplier["Saturn"] = 29.447498
    period_multiplier["Uranus"] = 84.016846
    period_multiplier["Neptune"] = 164.79132
}
{

    if (!($1 in period_multiplier)) {
        print "not a planet"
        exit 1
    }
    earth_years = $2 / (86400 * 365.25)
    planet_years = earth_years / period_multiplier[$1]
    printf "%.2f", planet_years
}