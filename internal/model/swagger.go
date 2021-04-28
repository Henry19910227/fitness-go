package model

type SuccessLoginResult struct {
	Code  int         `json:"code" example:"0"`     // 狀態碼
	Data  interface{} `json:"data"`                   // 回傳資料
	Msg   string      `json:"msg" example:"success!"` // 成功訊息
	Token string     `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ0MDc3NjMsInN1YiI6IjEwMDEzIn0.Z5UeEC8ArCVYej9kI1paXD2f5FMFiTfeLpU6e_CZZw0"` // Token
}

type SuccessResult struct {
	Code int         `json:"code" example:"0"`     // 狀態碼
	Data interface{} `json:"data"`                   // 回傳資料
	Msg  string      `json:"msg" example:"success!"` // 成功訊息
}

type SuccessPagingResult struct {
	Code int         `json:"code" example:"0"`      // 狀態碼
	Data interface{} `json:"data"`                    // 回傳資料
	Msg  string      `json:"msg" example:"success!"`  // 成功訊息
	Paging Paging    `json:"paging"`                  // 分頁資訊
}

type ErrorResult struct {
	Code int         `json:"code" example:"9000"`         // 錯誤碼
	Data interface{} `json:"data"`                        // 回傳資料
	Msg  string      `json:"msg" example:"system error!"` // 錯誤訊息
}

type Paging struct {
	Total int `json:"total_page" example:"10"` // 總頁數
	Page  int `json:"page" example:"1"`   // 當前頁數
	Size  int `json:"size" example:"5"`   // 一頁筆數
}
