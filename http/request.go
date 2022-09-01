package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Request(method, url string, header map[string]string, body interface{}, data interface{}) error {
	var client = &http.Client{}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	return err
}

func RequestByteFile(method, url string, header map[string]string, body interface{}, byteRes *[]byte) error {
	var client = &http.Client{}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	*byteRes, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return err
}

// timoeut.Duration example => 1*time.Second
// Can use time.Milliseconds, time.Nanoseconds, etc
func RequestWithoutResponse(method, url string, header map[string]string, body interface{}, timeout time.Duration) error {
	var client = &http.Client{
		Timeout: timeout,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))

	for k, v := range header {
		req.Header.Set(k, v)
	}

	_, _ = client.Do(req)

	return nil
}

func RequestFormUrlEncoded(method, urlPath string, header map[string]string, body map[string]string, data interface{}) error {
	var client = &http.Client{}

	urlBody := url.Values{}
	for k, v := range body {
		urlBody.Set(k, v)
	}

	req, _ := http.NewRequest(method, urlPath, strings.NewReader(urlBody.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	return err
}
