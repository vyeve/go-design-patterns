package main

import (
	"fmt"
	"log"

	"github.com/go-design-patterns/structural/facade"
)

const secretKey = "6b198464a3190fefba65cefbd8e710ec"

const (
	city        = "Lviv"
	countryCode = "UA"
)

func main() {
	weatherMap := facade.NewWeatherData(secretKey)
	weather, err := weatherMap.GetByCityAndCountryCode(city, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature in %s is %0.1f celsius\n", city, weather.Main.Temp-273.15)
	weather, err = weatherMap.GetByZipAndCountryCode(49102, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature in %s is %0.1f celsius\n", weather.Name, weather.Main.Temp-273.15)
}
