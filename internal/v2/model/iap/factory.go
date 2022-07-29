package iap

import (
	"encoding/base64"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"strings"
)

func NewIAPNotificationResponse(base64String string) *IAPNotificationResponse {
	//將base64字串轉換為map
	payloadString, err := base64.RawURLEncoding.DecodeString(base64String)
	if err != nil {
		return nil
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(payloadString, &dict); err != nil {
		return nil
	}
	//將map轉換為物件
	var response IAPNotificationResponse
	if err := mapstructure.Decode(dict, &response); err != nil {
		return nil
	}
	data, ok := dict["data"].(map[string]interface{})
	if ok {
		response.Data = NewIAPNotificationData(data)
	}
	return &response
}

func NewIAPNotificationData(input map[string]interface{}) *IAPNotificationData {
	var response IAPNotificationData
	if err := mapstructure.Decode(input, &response); err != nil {
		return nil
	}
	base64Transaction, ok := input["signedTransactionInfo"].(string)
	if ok {
		response.SignedTransactionInfo = NewSignedTransactionInfo(base64Transaction)
	}
	base64Renewal, ok := input["signedRenewalInfo"].(string)
	if ok {
		response.SignedRenewalInfo = NewSignedRenewalInfo(base64Renewal)
	}
	return &response
}

func NewIAPSubscribeAPIResponse(input map[string]interface{}) *IAPSubscribeAPIResponse {
	var response IAPSubscribeAPIResponse
	response.Data = make([]*SubscriptionGroupIdentifierItem, 0)
	if err := mapstructure.Decode(input, &response); err != nil {
		return nil
	}
	datas, ok := input["data"].([]interface{})
	if ok {
		for _, data := range datas {
			dict, ok := data.(map[string]interface{})
			if ok {
				response.Data = append(response.Data, NewSubscriptionGroupIdentifierItem(dict))
			}
		}
	}
	return &response
}

func NewIAPHistoryAPIResponse(data map[string]interface{}) *IAPHistoryAPIResponse {
	var response IAPHistoryAPIResponse
	response.SignedTransactions = make([]*SignedTransactionInfo, 0)
	if err := mapstructure.Decode(data, &response); err != nil {
		return nil
	}
	base64Transactions, ok := data["signedTransactions"].([]interface{})
	if ok {
		for _, base64Transaction := range base64Transactions {
			base64TransactionStr, ok := base64Transaction.(string)
			if ok {
				response.SignedTransactions = append(response.SignedTransactions, NewSignedTransactionInfo(base64TransactionStr))
			}
		}
	}
	return &response
}

func NewSubscriptionGroupIdentifierItem(data map[string]interface{}) *SubscriptionGroupIdentifierItem {
	var response SubscriptionGroupIdentifierItem
	response.LastTransactions = make([]*LastTransactionsItem, 0)
	if err := mapstructure.Decode(data, &response); err != nil {
		return nil
	}
	lastTransactions, ok := data["lastTransactions"].([]interface{})
	if ok {
		for _, lastTransaction := range lastTransactions {
			dict, ok := lastTransaction.(map[string]interface{})
			if ok {
				response.LastTransactions = append(response.LastTransactions, NewLastTransactionsItem(dict))
			}
		}
	}
	return &response
}

func NewLastTransactionsItem(data map[string]interface{}) *LastTransactionsItem {
	var response LastTransactionsItem
	if err := mapstructure.Decode(data, &response); err != nil {
		return nil
	}
	base64Transaction, ok := data["signedTransactionInfo"].(string)
	if ok {
		response.SignedTransactionInfo = NewSignedTransactionInfo(base64Transaction)
	}
	base64Renewal, ok := data["signedRenewalInfo"].(string)
	if ok {
		response.SignedRenewalInfo = NewSignedRenewalInfo(base64Renewal)
	}
	return &response
}

func NewSignedTransactionInfo(data string) *SignedTransactionInfo {
	payloadString, err := base64.RawURLEncoding.DecodeString(strings.Split(data, ".")[1])
	if err != nil {
		return nil
	}
	var transactionInfo map[string]interface{}
	if err := json.Unmarshal(payloadString, &transactionInfo); err != nil {
		return nil
	}
	var info SignedTransactionInfo
	if err := mapstructure.Decode(transactionInfo, &info); err != nil {
		return nil
	}
	return &info
}

func NewSignedRenewalInfo(data string) *SignedRenewalInfo {
	payloadString, err := base64.RawURLEncoding.DecodeString(strings.Split(data, ".")[1])
	if err != nil {
		return nil
	}
	var renewalInfo map[string]interface{}
	if err := json.Unmarshal(payloadString, &renewalInfo); err != nil {
		return nil
	}
	var info SignedRenewalInfo
	if err := mapstructure.Decode(renewalInfo, &info); err != nil {
		return nil
	}
	return &info
}
