package global

type PaymentOrderType int
const (
	BuyCourseOrderType PaymentOrderType = 1
	SubscribeOrderType PaymentOrderType = 2
)

type PaymentType int
const (
	NonePaymentType PaymentType = 0
	ApplePaymentType PaymentType = 1
	GooglePaymentType PaymentType = 2
)

type OrderStatus int
const (
	PendingOrderStatus OrderStatus = 1
	SuccessOrderStatus OrderStatus = 2
	ErrorOrderStatus OrderStatus = 3
)

type SubscribeLogType string
const (
	NormalSubscribeLogType SubscribeLogType = "" // 訂閱
	ExpiredSubscribeLogType SubscribeLogType = "" // 過期
	RefundSubscribeLogType SubscribeLogType = "" // 退費
)
