package global

type PaymentOrderType int

const (
	BuyCourseOrderType PaymentOrderType = 1
	SubscribeOrderType PaymentOrderType = 2
)

type PaymentType int

const (
	NonePaymentType   PaymentType = 0
	ApplePaymentType  PaymentType = 1
	GooglePaymentType PaymentType = 2
)

type OrderStatus int

const (
	PendingOrderStatus OrderStatus = 1
	SuccessOrderStatus OrderStatus = 2
	ErrorOrderStatus   OrderStatus = 3
)

type SubscribeLogType string

const (
	Unknown         SubscribeLogType = "unknown"          // 未知情況
	InitialBuy      SubscribeLogType = "initial_buy"      // 初次訂閱
	Resubscribe     SubscribeLogType = "resubscribe"      // 恢復訂閱
	Renew           SubscribeLogType = "renew"            // 續訂
	Expired         SubscribeLogType = "expired"          // 過期
	Upgrade         SubscribeLogType = "upgrade"          // 訂閱升級
	Downgrade       SubscribeLogType = "downgrade"        // 訂閱降級
	DowngradeCancel SubscribeLogType = "downgrade_cancel" // 取消訂閱降級
	Refund          SubscribeLogType = "refund"           // 退費
	RenewEnable     SubscribeLogType = "renew_enable"     // 啟用續訂
	RenewDisable    SubscribeLogType = "renew_disable"    // 取消續訂
)
