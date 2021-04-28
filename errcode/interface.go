package errcode

import "errors"

var (

	// Common
	systemError       = NewError(9000, errors.New("系統發生錯誤"))
	updateError       = NewError(9001, errors.New("更新失敗"))
	dataNotFound      = NewError(9002, errors.New("查無資料"))
	dataAlreadyExists = NewError(9003, errors.New("資料已存在"))
	InvalidThirdParty = NewError(9004, errors.New("無效的第三方驗證"))
	InvalidToken      = NewError(9005, errors.New("無效的token"))
)

type Error interface {
	Code() int
	Msg() string
}

type Common interface {
	Custom(code int, err error) Error
	// SystemError 9000 - 系統發生錯誤
	SystemError() Error
	// UpdateError 9001 - 更新失敗
	UpdateError() Error
	// DataNotFound 9002 - 查無資料
	DataNotFound() Error
	// DataAlreadyExists 9003 - 資料已存在
	DataAlreadyExists() Error
	// InvalidThirdParty 9004 - 無效的第三方驗證
	InvalidThirdParty() Error
	// InvalidToken 9005 - 無效的token
	InvalidToken() Error
}
