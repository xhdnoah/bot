package handler

import (
	"crypto/tls"
	"encoding/json"
	"go-chat/errno"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	weatherToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOiJlZWQ4ZmQ1ODBmYTRmNjkyIn0.d7qF_mjdXMC0R5M6f04Lnh6x61kaU4lqHT0Axt9xUOY"
	weatherAPI   = "https://www.douyacun.com/api/openapi/weather?weather_type=observe|air&token="
	stockQuery   = "https://query1.finance.yahoo.com/v7/finance/quote?symbols="
)

type BotMessage struct {
	Method string `json:"method"`
	Args   string `json:"args"`
}

type RetMessage struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func assembleRet(rm *RetMessage) string {
	json, err := json.Marshal(rm)
	if err != nil {
		return ""
	}
	return string(json)
}

func get(url string, headers map[string]string) *http.Response {
	// 解决x509: certificate signed by unknown authority
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return resp
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func HandleAi(ai chan<- string, msg string) {
	go func() {
		if ret := TlBot(msg); ret != "" {
			ai <- ret
		}
	}()
	go func() {
		if ret := TianBot(msg); ret != "" {
			ai <- ret
		}
	}()
}

func HandleWeather(w chan<- string) {
	ret := FetchWeather()
	w <- ret
}

func HandleStock(st chan<- string, symbol string) {
	ret := FetchStock(symbol)
	st <- ret
}

func HandleDefine(d chan<- string, word string) {
	go Translator(word, d)
}
