package controller

import (
	"example/weather-rest-api/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// panicError is a helper function to panic on error
func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

// WeatherData returns the weather data for the given query and number of days
func WeatherData(query string, days int, dt string) (int, []byte) {
	weatherUrl := config.GetConfig().GetString("server.weatherUrl")
	if dt == "empty" {
		// get the current date in the format YYYY-MM-DD
		dt = time.Now().Local().Format("2006-01-02")
	}
	url := fmt.Sprintf("%s?q=%s&days=%d&dt=%s", weatherUrl, query, days, dt)

	req, err := http.NewRequest("GET", url, nil)
	panicError(err)

	// Set the headers
	rapidAPIKey := config.GetConfig().GetString("server.RpidAPIKey")
	rapidAPIHost := config.GetConfig().GetString("server.RapidAPIHost")

	// add the RapidAPI key to the request header
	req.Header.Add("x-RapidAPI-Key", rapidAPIKey)
	req.Header.Add("x-RapidAPI-Host", rapidAPIHost)

	// request to RapidAPI server
	resp, err := http.DefaultClient.Do(req)

	panicError(err)

	// read the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	panicError(err)

	return resp.StatusCode, body
}
