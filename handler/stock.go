package handler

import "io/ioutil"

func FetchStock(symbol string) string {
	url := stockQuery + symbol
	resp := get(url, nil)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
