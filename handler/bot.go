package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	USED_OUT_CODE = 40004
	TL_KEY        = "a5052a22b8232be1e387ff153e823975"
	TIAN_KEY      = "cf9d1148cc012ccec1a4a9ecc1d25cd1"
)

type tianNews struct {
	Reply    string `json:"reply"`
	DataType string `json:"datatype"`
}

type tianReply struct {
	code     int
	msg      string
	NewsList []tianNews `json:"newslist"`
}

func TianBot(question string) string {
	url := fmt.Sprintf("http://api.tianapi.com/txapi/robot/index?key=%s&question=%s", TIAN_KEY, question)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	reply := new(tianReply)
	err := json.NewDecoder(res.Body).Decode(&reply)
	if err != nil {
		return ""
	}
	var sb strings.Builder
	for _, v := range reply.NewsList {
		if v.DataType == "text" {
			sb.WriteString(v.Reply + "\r\n")
		}
	}
	return sb.String()
}

type tlReply struct {
	Code int
	URL  string `json:"url, omitempty"`
	Text string `json:"text"`
}

func TlBot(info string) string {
	url := fmt.Sprintf("http://www.tuling123.com/openapi/api?key=%s&info=%s", TL_KEY, info)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	reply := new(tlReply)
	decodeErr := json.NewDecoder(resp.Body).Decode(&reply)
	if decodeErr != nil || reply.Code == USED_OUT_CODE {
		return ""
	}
	return reply.Text
}
