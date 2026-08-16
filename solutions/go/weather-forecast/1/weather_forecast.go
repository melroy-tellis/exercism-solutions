// Package weather provides tools to forecast the weather
// conditions in a city.
package weather

// CurrentCondition represents the weather condition in the current city.
var CurrentCondition string
// CurrentLocation represents the current city.
var CurrentLocation string

// Forecast returns a string value describing the weather in the
// specified city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
