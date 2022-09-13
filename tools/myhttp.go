// Package tools Package utils /**
package tools

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("get errorhandler")
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("ioutil read errorhandler")
	}
	return string(body), err
}

func HttpPost(url string, jsonStr string) (string, error) {
	var json = []byte(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("post errorhandler")
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println("ioutil read errorhandler")
	}
	return string(body), err
}
