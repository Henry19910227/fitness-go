package dto

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

// IABProductAPIResponse Google Play Developer API 獲取購買商品回傳
// GET https://androidpublisher.googleapis.com/androidpublisher/v3/applications/{packageName}/purchases/products/{productId}/tokens/{token}
type IABProductAPIResponse struct {
	Kind                        string `mapstructure:"kind" json:"kind"`
	PurchaseTimeMillis          string `mapstructure:"purchaseTimeMillis" json:"purchaseTimeMillis"`
	PurchaseState               int    `mapstructure:"purchaseState" json:"purchaseState"`
	ConsumptionState            int    `mapstructure:"consumptionState" json:"consumptionState"`
	DeveloperPayload            string `mapstructure:"developerPayload" json:"developerPayload"`
	OrderId                     string `mapstructure:"orderId" json:"orderId"`
	PurchaseType                int    `mapstructure:"purchaseType" json:"purchaseType"`
	AcknowledgementState        int    `mapstructure:"acknowledgementState" json:"acknowledgementState"`
	PurchaseToken               string `mapstructure:"purchaseToken" json:"purchaseToken"`
	ProductId                   string `mapstructure:"productId" json:"productId"`
	Quantity                    int    `mapstructure:"quantity" json:"quantity"`
	ObfuscatedExternalAccountId string `mapstructure:"obfuscatedExternalAccountId" json:"obfuscatedExternalAccountId"`
	ObfuscatedExternalProfileId string `mapstructure:"obfuscatedExternalProfileId" json:"obfuscatedExternalProfileId"`
	RegionCode                  string `mapstructure:"regionCode" json:"regionCode"`
}

// IABSubscriptionAPIResponse Google Play Developer API 獲取訂閱狀態回傳
// GET https://androidpublisher.googleapis.com/androidpublisher/v3/applications/{packageName}/purchases/subscriptions/{subscriptionId}/tokens/{token}
type IABSubscriptionAPIResponse struct {
	Kind                        string                          `mapstructure:"kind" json:"kind"`
	StartTimeMillis             string                          `mapstructure:"startTimeMillis" json:"startTimeMillis"`
	ExpiryTimeMillis            string                          `mapstructure:"expiryTimeMillis" json:"expiryTimeMillis"`
	AutoResumeTimeMillis        string                          `mapstructure:"autoResumeTimeMillis" json:"autoResumeTimeMillis"`
	AutoRenewing                bool                            `mapstructure:"autoRenewing" json:"autoRenewing"`
	PriceCurrencyCode           string                          `mapstructure:"priceCurrencyCode" json:"priceCurrencyCode"`
	PriceAmountMicros           string                          `mapstructure:"priceAmountMicros" json:"priceAmountMicros"`
	IntroductoryPriceInfo       *IntroductoryPriceInfo          `mapstructure:"introductoryPriceInfo" json:"introductoryPriceInfo"`
	CountryCode                 string                          `mapstructure:"countryCode" json:"countryCode"`
	DeveloperPayload            string                          `mapstructure:"developerPayload" json:"developerPayload"`
	PaymentState                int                             `mapstructure:"paymentState" json:"paymentState"` //0. Payment pending 1. Payment received 2. Free trial 3. Pending deferred upgrade/downgrade
	CancelReason                int                             `mapstructure:"cancelReason" json:"cancelReason"`
	UserCancellationTimeMillis  string                          `mapstructure:"userCancellationTimeMillis" json:"userCancellationTimeMillis"`
	CancelSurveyResult          *SubscriptionCancelSurveyResult `mapstructure:"cancelSurveyResult" json:"cancelSurveyResult"`
	OrderId                     string                          `mapstructure:"orderId" json:"orderId"`
	LinkedPurchaseToken         string                          `mapstructure:"linkedPurchaseToken" json:"linkedPurchaseToken"`
	PurchaseType                int                             `mapstructure:"purchaseType" json:"purchaseType"`
	PriceChange                 *SubscriptionPriceChange        `mapstructure:"priceChange" json:"priceChange"`
	ProfileName                 string                          `mapstructure:"profileName" json:"profileName"`
	EmailAddress                string                          `mapstructure:"emailAddress" json:"emailAddress"`
	GivenName                   string                          `mapstructure:"givenName" json:"givenName"`
	FamilyName                  string                          `mapstructure:"familyName" json:"familyName"`
	ProfileId                   string                          `mapstructure:"profileId" json:"profileId"`
	AcknowledgementState        int                             `mapstructure:"acknowledgementState" json:"acknowledgementState"`
	ExternalAccountId           string                          `mapstructure:"externalAccountId" json:"externalAccountId"`
	PromotionType               int                             `mapstructure:"promotionType" json:"promotionType"`
	PromotionCode               string                          `mapstructure:"promotionCode" json:"promotionCode"`
	ObfuscatedExternalAccountId string                          `mapstructure:"obfuscatedExternalAccountId" json:"obfuscatedExternalAccountId"`
	ObfuscatedExternalProfileId string                          `mapstructure:"obfuscatedExternalProfileId" json:"obfuscatedExternalProfileId"`
}

type IABSubscribeNotificationResponse struct {
	EventTimeMillis          string                       `mapstructure:"eventTimeMillis" json:"eventTimeMillis"`
	PackageName              string                       `mapstructure:"packageName" json:"packageName"`
	Version                  string                       `mapstructure:"version" json:"version"`
	SubscriptionNotification *IABSubscriptionNotification `mapstructure:"subscriptionNotification" json:"subscriptionNotification"`
}

type IntroductoryPriceInfo struct {
	IntroductoryPriceCurrencyCode string `mapstructure:"introductoryPriceCurrencyCode" json:"introductoryPriceCurrencyCode"`
	IntroductoryPriceAmountMicros string `mapstructure:"introductoryPriceAmountMicros" json:"introductoryPriceAmountMicros"`
	IntroductoryPricePeriod       string `mapstructure:"introductoryPricePeriod" json:"introductoryPricePeriod"`
	IntroductoryPriceCycles       int    `mapstructure:"introductoryPriceCycles" json:"introductoryPriceCycles"`
}

type SubscriptionCancelSurveyResult struct {
	CancelSurveyReason    int    `mapstructure:"cancelSurveyReason" json:"cancelSurveyReason"`
	UserInputCancelReason string `mapstructure:"userInputCancelReason" json:"userInputCancelReason"`
}

type SubscriptionPriceChange struct {
	NewPrice Price `mapstructure:"newPrice" json:"newPrice"`
	State    int   `mapstructure:"state" json:"state"`
}

type IABSubscriptionNotification struct {
	NotificationType int    `mapstructure:"notificationType" json:"notificationType"`
	PurchaseToken    string `mapstructure:"purchaseToken" json:"purchaseToken"`
	SubscriptionId   string `mapstructure:"subscriptionId" json:"subscriptionId"`
	Version          string `mapstructure:"version" json:"version"`
}

type Price struct {
	PriceMicros string `mapstructure:"priceMicros" json:"priceMicros"`
	Currency    string `mapstructure:"currency" json:"currency"`
}

func NewIABProductAPIResponse(data map[string]interface{}) *IABProductAPIResponse {
	var response IABProductAPIResponse
	if err := mapstructure.Decode(data, &response); err != nil {
		return nil
	}
	return &response
}

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
