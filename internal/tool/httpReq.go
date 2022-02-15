package tool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type httpRequest struct {}

func NewRequest() HttpRequest {
	return &httpRequest{}
}

func (h *httpRequest) SendPostRequestWithJsonBody(url string, param map[string]interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(body, &dict); err != nil {
		return nil, err
	}
	return dict, nil
}