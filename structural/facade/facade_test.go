package facade

import (
	"testing"
)

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := openWeatherData{apiKey: ""}
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d", weather.ID)
	}

}
