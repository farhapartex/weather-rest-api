package controller

import (
	"example/weather-rest-api/config"
	"fmt"
	"io/ioutil"
	"net/http"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func WeatherData(query string, days int) (int, []byte) {
	weatherUrl := config.GetConfig().GetString("server.weatherUrl")
	url := fmt.Sprintf("%s?q=%s&days=%d", weatherUrl, query, days)

	req, err := http.NewRequest("GET", url, nil)
	panicError(err)

	rapidAPIKey := config.GetConfig().GetString("server.RpidAPIKey")
	rapidAPIHost := config.GetConfig().GetString("server.RapidAPIHost")

	req.Header.Add("x-RapidAPI-Key", rapidAPIKey)
	req.Header.Add("x-RapidAPI-Host", rapidAPIHost)

	resp, err := http.DefaultClient.Do(req)

	panicError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	panicError(err)

	return resp.StatusCode, body
}
