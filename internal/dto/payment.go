package dto

import "time"

type AppleReceiptResponse struct {
	Environment string `mapstructure:"environment"`
	Receipt     struct {
		BundleID string         `mapstructure:"bundle_id"`
		InApp    []*ReceiptInfo `mapstructure:"in_app"`
	} `mapstructure:"receipt"`
	LatestReceiptInfo  []*ReceiptInfo        `mapstructure:"latest_receipt_info"`
	PendingRenewalInfo []*PendingRenewalInfo `mapstructure:"pending_renewal_info"`
	Status             int                   `mapstructure:"status"`
}

type GoogleReceiptResponse struct {
	Kind                        string      `mapstructure:"kind"`
	PurchaseTimeMillis          string      `mapstructure:"purchaseTimeMillis"`
	PurchaseState               int         `mapstructure:"purchaseState"`
	ConsumptionState            int         `mapstructure:"consumptionState"`
	DeveloperPayload            string      `mapstructure:"developerPayload"`
	OrderID                     string      `mapstructure:"orderId"`
	PurchaseType                int         `mapstructure:"purchaseType"`
	AcknowledgementState        int         `mapstructure:"acknowledgementState"`
	PurchaseToken               string      `mapstructure:"purchaseToken"`
	ProductID                   string      `mapstructure:"productId"`
	Quantity                    interface{} `mapstructure:"quantity"`
	ObfuscatedExternalAccountId string      `mapstructure:"obfuscatedExternalAccountId"`
	ObfuscatedExternalProfileId string      `mapstructure:"obfuscatedExternalProfileId"`
	RegionCode                  string      `mapstructure:"regionCode"`
}

type ReceiptInfo struct {
	Quantity               string `mapstructure:"quantity"`
	ProductID              string `mapstructure:"product_id"`
	TransactionID          string `mapstructure:"transaction_id"`
	OriginalTransactionID  string `mapstructure:"original_transaction_id"`
	PurchaseDateMS         string `mapstructure:"purchase_date_ms"`
	OriginalPurchaseDateMS string `mapstructure:"original_purchase_date_ms"`
	ExpiresDateMS          string `mapstructure:"expires_date_ms"`
	OriginalPurchaseDate   *time.Time
	PurchaseDate           *time.Time
	ExpiresDate            *time.Time
}

type PendingRenewalInfo struct {
	ExpirationIntent       string `mapstructure:"expiration_intent"`
	AutoRenewProductID     string `mapstructure:"auto_renew_product_id"`
	IsInBillingRetryPeriod string `mapstructure:"is_in_billing_retry_period"`
	ProductID              string `mapstructure:"product_id"`
	OriginalTransactionID  string `mapstructure:"original_transaction_id"`
	AutoRenewStatus        string `mapstructure:"auto_renew_status"`
}

type SignedRenewalInfo struct {
	AutoRenewProductId     string `mapstructure:"autoRenewProductId"`     // 在下一個計費周期續訂的產品的產品標識符
	AutoRenewStatus        int    `mapstructure:"autoRenewStatus"`        // 自動續訂訂閱的續訂狀態(0:自動續訂已關閉，客戶已關閉訂閱的自動續訂，並且在當前訂閱期結束時不會續訂/1:自動續訂已開啟，訂閱在當前訂閱期結束時續訂)
	ExpirationIntent       int    `mapstructure:"expirationIntent"`       // 訂閱過期的原因(1:客戶取消了訂閱/2:發生計費錯誤；例如，客戶的付款信息不再有效/3:客戶不同意最近的價格上漲/4:該產品在續訂時無法購買)
	GracePeriodExpiresDate int64  `mapstructure:"gracePeriodExpiresDate"` // 訂閱續訂的計費寬限期到期的時間
	IsInBillingRetryPeriod bool   `mapstructure:"isInBillingRetryPeriod"` // 一個布爾值，指示 App Store 是否正在嘗試自動續訂過期的訂閱
	OfferIdentifier        string `mapstructure:"offerIdentifier"`        // 包含促銷代碼或促銷優惠標識符的標識符
	OfferType              string `mapstructure:"offerType"`              //促銷優惠的類型
	OriginalTransactionId  string `mapstructure:"originalTransactionId"`  // 購買的原始交易標識符
	PriceIncreaseStatus    int64  `mapstructure:"priceIncreaseStatus"`    // 指示客戶是否已批准訂閱價格上漲的狀態
	ProductId              string `mapstructure:"productId"`              // 應用內購買的產品標識符
	SignedDate             int64  `mapstructure:"signedDate"`             // App Store 對 JSON Web 簽名數據進行簽名的 UNIX 時間（以毫秒為單位）
}

type SignedTransactionInfo struct {
	AppAccountToken             string `mapstructure:"appAccountToken"`
	BundleId                    string `mapstructure:"bundleId"`
	ExpiresDate                 int64  `mapstructure:"expiresDate"`                 // 訂閱到期或續訂的 UNIX 時間（以毫秒為單位）
	InAppOwnershipType          string `mapstructure:"inAppOwnershipType"`          // 描述交易是由用戶購買還是可通過家庭共享提供給他們的字符串(FAMILY_SHARED:該交易屬於受益於該服務的家庭成員/PURCHASED:交易屬於買方)
	IsUpgraded                  bool   `mapstructure:"isUpgraded"`                  // 用戶是否升級到另一個訂閱
	OfferIdentifier             string `mapstructure:"offerIdentifier"`             // 包含促銷代碼或促銷優惠標識符的標識符
	OfferType                   int    `mapstructure:"offerType"`                   // 促銷優惠的類型(1:介紹性報價/2:促銷優惠/3:帶有訂閱優惠代碼的優惠)
	OriginalPurchaseDate        int64  `mapstructure:"originalPurchaseDate"`        // UNIX 時間，以毫秒為單位，表示原始交易標識符的購買日期
	OriginalTransactionId       string `mapstructure:"originalTransactionId"`       // 原始購買的交易標識符
	ProductId                   string `mapstructure:"productId"`                   // 應用內購買的產品標識符
	PurchaseDate                int64  `mapstructure:"purchaseDate"`                // App Store 向用戶帳戶收取購買、恢復產品、訂閱或訂閱續訂的時間
	Quantity                    int64  `mapstructure:"quantity"`                    // 用戶購買的消耗品數量
	RevocationDate              int64  `mapstructure:"revocationDate"`              // App Store 退還交易或從家庭共享中撤銷交易的 UNIX 時間（以毫秒為單位）
	RevocationReason            string `mapstructure:"revocationReason"`            // App Store 退還交易或從家庭共享中撤銷交易的原因(0:Apple Support 因其他原因代表客戶退還交易；例如，意外購買/1:由於您的應用程序中存在實際或感知的問題，Apple 支持代表客戶退還了交易)
	SignedDate                  int64  `mapstructure:"signedDate"`                  // App Store 對 JSON Web 簽名 (JWS) 數據進行簽名的 UNIX 時間（以毫秒為單位)
	SubscriptionGroupIdentifier string `mapstructure:"subscriptionGroupIdentifier"` // 訂閱所屬訂閱組的標識
	TransactionId               string `mapstructure:"transactionId"`               // 交易的唯一標識符
	Type                        string `mapstructure:"type"`                        // 應用內購買的產品類型(Auto-Renewable Subscription/Non-Consumable/Consumable/Non-Renewing Subscription)
	WebOrderLineItemId          string `mapstructure:"webOrderLineItemId"`          // 跨設備訂閱購買事件的唯一標識符，包括訂閱續訂
}

type IAPNotificationResponse struct {
	NotificationType    string `mapstructure:"notificationType"`    // 應用內購買事件 https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
	Subtype             string `mapstructure:"subtype"`             // 通知類型的詳細信息的字符串 https://developer.apple.com/documentation/appstoreservernotifications/subtype
	NotificationUUID    string `mapstructure:"notificationUUID"`    // 通知的唯一標識符。使用此值來識別重複通知
	NotificationVersion string `mapstructure:"notificationVersion"` // 通知的版本號
	Data                struct {
		AppAppleId            string                `mapstructure:"appAppleId"`  // 通知適用的應用的唯一標識符
		BundleId              string                `mapstructure:"bundleId"`    // 應用程序的捆綁標識符
		BundleVersion         string                `mapstructure:"bundleId"`    // 標識捆綁包迭代的構建版本
		Environment           string                `mapstructure:"environment"` // 通知適用的服務器環境，沙盒或生產環境
		SignedRenewalInfo     SignedRenewalInfo     `mapstructure:"omitempty"`   // 由 App Store 簽名的訂閱續訂信息，採用 JSON Web 簽名格式
		SignedTransactionInfo SignedTransactionInfo `mapstructure:"omitempty"`   // 由 App Store 簽名的交易信息，採用 JSON Web 簽名格式
	} `mapstructure:"data"` // 包含應用元數據和簽名續訂和交易信息的對象
}

type IAPSubscribeResponse struct {
	Status                int                   // 訂閱狀態 (1:訂閱處於活動狀態/2:訂閱已過期/3:訂閱處於計費重試期/4:訂閱處於計費寬限期/5:訂閱被撤銷)
	OriginalTransactionID string                //原始購買的交易標識符
	SignedRenewalInfo     SignedRenewalInfo     `mapstructure:"omitempty"`
	SignedTransactionInfo SignedTransactionInfo `mapstructure:"omitempty"`
}
