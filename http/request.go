package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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
