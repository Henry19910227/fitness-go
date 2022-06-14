package base

type CodeField struct {
	Code int `json:"code" example:"9000"` // 狀態碼
}
type MsgField struct {
	Msg string `json:"msg" example:"message.."` // 訊息
}
