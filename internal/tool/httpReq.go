package tool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type httpRequest struct{}

func NewRequest() HttpRequest {
	return &httpRequest{}
}

func (h *httpRequest) SendRequest(method string, url string, header map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	reqs, err := h.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	reqs.Header.Add("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			reqs.Header.Add(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(reqs)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(resBody, &dict); err != nil {
		return nil, err
	}
	return dict, nil
}

func (h *httpRequest) NewRequest(method string, url string, body map[string]interface{}) (*http.Request, error) {
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqs, err := http.NewRequest(method, url, bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}
		return reqs, nil
	}
	reqs, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}
