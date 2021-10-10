package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

const (
	DEEPL_AUTH_KEY = "d601ad14-1297-2ba8-ee31-9df3be7b5153:fx"
	DEEPL_API      = "http://api-free.deepl.com/v2/?text="
)

func Translator(s string, results chan<- string) {
	defer close(results)
	// results <- "searching definitions"
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		queryDeepL(s, results)
	}()
	go func() {
		defer wg.Done()
		queryYoudao(s, results)
	}()
	wg.Wait()
}

func queryDeepL(text string, results chan<- string) {
	client := &http.Client{}
	url := DEEPL_API + text
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println("deepl is down", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("deepl is down", err)
		return
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// if body != nil {

	// }
	// results <- string(body)
}

func queryYoudao(text string, results chan<- string) {
	url := fmt.Sprintf("http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=%s", text)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("youdao is down", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	results <- jsoniter.Get(body, "translateResult", 0, 0, "tgt").ToString()
}
