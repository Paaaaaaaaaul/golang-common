package common

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 0,
			DualStack: true,
		}).DialContext,
		IdleConnTimeout:       time.Minute,
		ResponseHeaderTimeout: 5 * time.Second,
	},
}

func GET(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func POST(urlStr string, data map[string]string) ([]byte, error) {
	v := make(url.Values)
	for key, val := range data {
		v.Set(key, val)
	}
	contentReader := strings.NewReader(v.Encode())
	contentType := "application/x-www-form-urlencoded;charset=utf-8"
	request, err := http.NewRequest("POST", urlStr, contentReader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
