package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HttpRequest(method string, url string, headers map[string]string, payload io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	statusOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOk {
		err := errors.New("Error! HTTP Response Code: " + strconv.Itoa(resp.StatusCode) + ", Message: " + string(respBody))
		return nil, err
	}

	return respBody, nil
}
