/*
The Facade design pattern shields the code from unwanted access, orders some calls,
and hides the complexity scope from the user.
You use Facade when you want to hide the complexity of some tasks, especially when most
of them share utilities (such as authentication an API).
Use the Facade design pattern in the following scenarios:
 - When you want to decrease the complexity of some parts of your code. You hide that
complexity behind the facade by providing a more easy-to-use method.
 - When you want to group actions that are cross-related in a single place.
 - When you want to build a library so that others can use your products without worrying
about hot it all works.
*/

package facade

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (*Weather, error)
	GetByGeoCoordinates(lat, lon float32) (*Weather, error)
	GetByZipAndCountryCode(zip int, countryCode string) (*Weather, error)
}

func NewWeatherData(apiKey string) CurrentWeatherDataRetriever {
	return &openWeatherData{apiKey: apiKey}
}

type openWeatherData struct {
	apiKey string
}

func (*openWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)

	if err != nil {
		return nil, err
	}
	return w, nil
}

func (o *openWeatherData) doRequest(uri string) (*Weather, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = errors.New(string(byt))
		}
		err = fmt.Errorf("status code was %d, aborting. Error message was:\n%s\n", resp.StatusCode, errMsg)
		return nil, err
	}
	weather, err := o.responseParser(resp.Body)
	defer resp.Body.Close()
	return weather, err
}

func (o *openWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return o.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s",
			lat, lon, o.apiKey),
	)
}

func (o *openWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return o.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s",
			city, countryCode, o.apiKey),
	)
}

func (o *openWeatherData) GetByZipAndCountryCode(zip int, countryCode string) (*Weather, error) {
	return o.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%d,%s&APPID=%s",
			zip, countryCode, o.apiKey),
	)
}
