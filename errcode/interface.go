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

	// Login
	LoginFailure     = NewError(1100, errors.New("登入失敗, 帳號或密碼錯誤"))
	LoginRoleFailure = NewError(1101, errors.New("登入身份錯誤"))
	LoginStatusFailure = NewError(1102, errors.New("帳號無法使用"))


	// Register
	RegisterFailure    = NewError(1400, errors.New("註冊失敗"))
	SendOTPFailure     = NewError(1401, errors.New("信箱驗證碼發送失敗"))
	OTPInvalid         = NewError(1402, errors.New("無效的信箱驗證碼"))
	NicknameDuplicate  = NewError(1405, errors.New("該暱稱已被使用"))
	EmailDuplicate     = NewError(1406, errors.New("該信箱已被使用"))
	AccountDuplicate   = NewError(1407, errors.New("該帳號已被使用"))
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


type Handler interface {
	/** 公共 */
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

	/** 註冊 */
	// NewError(1400, errors.New("註冊失敗"))
	RegisterFailure() Error
	// NewError(1401, errors.New("手機驗證碼發送失敗"))
	SendOTPFailure() Error
	// NewError(1402, errors.New("無效的手機驗證碼"))
	OTPInvalid() Error
	// NewError(1405, errors.New("該暱稱已被使用"))
	NicknameDuplicate() Error
	// NewError(1406, errors.New("該信箱已被使用"))
	EmailDuplicate() Error
	// NewError(1407, errors.New("該帳號已被使用"))
	AccountDuplicate() Error
}


type Login interface {
	Common
	// LoginFailure 1100 - 登入失敗, 帳號或密碼錯誤
	LoginFailure() Error
	// NewError(1101, errors.New("登入身份錯誤"))
	LoginRoleFailure() Error
	//NewError(1102, errors.New("帳號無法使用"))
	LoginStatusFailure() Error
}

type Register interface {
	Common
	//// NewError(1400, errors.New("註冊失敗"))
	//RegisterFailure() Error
	//// NewError(1401, errors.New("手機驗證碼發送失敗"))
	//SendOTPFailure() Error
	//// NewError(1402, errors.New("無效的手機驗證碼"))
	//MobileInvalid() Error
	//// NewError(1405, errors.New("該暱稱已被使用"))
	//NicknameDuplicate() Error
	//// NewError(1406, errors.New("該信箱已被使用"))
	//EmailDuplicate() Error
	//// NewError(1407, errors.New("該帳號已被使用"))
	//AccountDuplicate() Error
}