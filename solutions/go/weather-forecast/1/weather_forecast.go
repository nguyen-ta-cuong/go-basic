// Package Weather use for predict weather forecast
package weather

var (
	CurrentCondition string // weather condition
	CurrentLocation  string // weather of the location
)

/*
* This public function for forecast
- city location
- weather condition such as sunny, rain or foggy
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
