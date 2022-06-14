package global

type SubscribeStatus int
const (
	NoneSubscribeStatus  SubscribeStatus = 0 // 無會員狀態
	ValidSubscribeStatus SubscribeStatus = 1 // 付費會員狀態
)
