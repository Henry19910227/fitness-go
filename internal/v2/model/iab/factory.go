package iab

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func NewIABSubscriptionAPIResponse(data map[string]interface{}) *IABSubscriptionAPIResponse {
	var response IABSubscriptionAPIResponse
	if err := mapstructure.Decode(data, &response); err != nil {
		return nil
	}
	return &response
}

func NewIABSubscribeNotificationResponse(base64String string) *IABSubscribeNotificationResponse {
	//將base64字串轉換為map
	payloadString, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(payloadString, &dict); err != nil {
		return nil
	}
	//將map轉換為物件
	var response IABSubscribeNotificationResponse
	if err := mapstructure.Decode(dict, &response); err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func NewIABProductAPIResponse(data map[string]interface{}) *IABProductAPIResponse {
	var response IABProductAPIResponse
	if err := mapstructure.Decode(data, &response); err != nil {
		return nil
	}
	return &response
}
