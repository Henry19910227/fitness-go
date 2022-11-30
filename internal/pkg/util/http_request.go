package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendRequest(method string, url string, header map[string]string, body map[string]interface{}, param map[string]interface{}) (map[string]interface{}, error) {
	reqs, err := NewRequest(method, url, body, param)
	if err != nil {
		return nil, err
	}
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

func NewRequest(method string, url string, body interface{}, param map[string]interface{}) (*http.Request, error) {
	if len(param) > 0 {
		var paramStr string
		for k, v := range param {
			if len(paramStr) == 0 {
				paramStr += "?"
			} else {
				paramStr += "&"
			}
			paramStr += fmt.Sprintf("%v=%v", k, v)
		}
		url = url + paramStr
	}
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
