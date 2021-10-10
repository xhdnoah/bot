package handler

import "io/ioutil"

func FetchWeather() string {
	resp := get(weatherAPI+weatherToken, nil)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
