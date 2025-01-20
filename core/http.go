package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendGetRequest(url string, body any, param map[string]string, header map[string]string) ([]byte, error) {
	marshal, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodGet, url, bytes.NewReader(marshal))
	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}
	if param != nil {
		params := request.URL.Query()
		for k, v := range param {
			params.Add(k, v)
		}
		request.URL.RawQuery = params.Encode()
	}
	return sendHttpRequest(request)
}

func sendHttpRequest(request *http.Request) ([]byte, error) {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
