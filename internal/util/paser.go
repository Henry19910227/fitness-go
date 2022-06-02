package util

import (
	"encoding/json"
)

func Parser(input interface{}, output interface{}) error {
	marshal, err := json.Marshal(input)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(marshal, output); err != nil {
		return err
	}
	return nil
}
