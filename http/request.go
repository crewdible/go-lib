package http

import (
	"bytes"
	"encoding/json"
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

func RequestByteFile(method, url string, header map[string]string, body interface{}, errData interface{}, byteRes *[]byte) error {
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

	err = json.NewDecoder(resp.Body).Decode(&errData)
	if err != nil {
		// *byteRes, err = io.ReadAll(resp.Body)
		err = json.NewDecoder(resp.Body).Decode(&byteRes)
		if err != nil {
			return err
		}
		// Use this to write file from response (If can)
		// src : https://stackoverflow.com/questions/16311232/how-to-pipe-an-http-response-to-a-file-in-go
		// f, e := os.Create("filename.pdf")
		// if err != nil {
		// 	return err
		// }
		// defer f.Close()
		// f.ReadFrom(resp.Body)
	}

	return err
}
