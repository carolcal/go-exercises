//Package weather provides weather information by implementing a client for weather.com API.
package weather

//CurrentCondition represents the weather condition at the current time at the CurrentLocation.
var CurrentCondition string

//CurrentLocation shows the city where we want to get weather conditions.
var CurrentLocation string

// Forecast receives the city and its condition as parameters, sets CurrentLocation and CurrentCondition to those values, and returns a string describing the location and its condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
