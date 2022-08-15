package subscribe_log

// Type Enum
const (
	Unknown         string = "unknown"          // 未知情況
	InitialBuy      string = "initial_buy"      // 初次訂閱
	Resubscribe     string = "resubscribe"      // 恢復訂閱
	Renew           string = "renew"            // 續訂
	Expired         string = "expired"          // 過期
	Upgrade         string = "upgrade"          // 訂閱升級
	Downgrade       string = "downgrade"        // 訂閱降級
	DowngradeCancel string = "downgrade_cancel" // 取消訂閱降級
	Refund          string = "refund"           // 退費
	RenewEnable     string = "renew_enable"     // 啟用續訂
	RenewDisable    string = "renew_disable"    // 取消續訂
)
