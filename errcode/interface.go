package errcode

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	SystemError int = 9000 // 系統發生錯誤
	UpdateError int = 9001 // 更新失敗
	DataNotFound int = 9002 // 查無資料
	DataAlreadyExists int = 9003 // 資料已存在
	InvalidThirdParty int = 9004 // 無效的第三方驗證
	InvalidToken int = 9005 // 無效的token
	PermissionDenied int = 9006 // 權限不足,存取遭拒
	FileTypeError int = 9007 // 上傳檔案類型不符合規範
	FileSizeError int = 9008 // 上傳檔案大小超過限制
	FileCountError int = 9009 // 上傳檔案數量超過限制
)

var (

	// Common
	systemError       = NewError(SystemError, errors.New("系統發生錯誤"))
	updateError       = NewError(UpdateError, errors.New("更新失敗"))
	dataNotFound      = NewError(DataNotFound, errors.New("查無資料"))
	dataAlreadyExists = NewError(DataAlreadyExists, errors.New("資料已存在"))
	invalidThirdParty = NewError(InvalidThirdParty, errors.New("無效的第三方驗證"))
	invalidToken      = NewError(InvalidToken, errors.New("無效的token"))
	permissionDenied  = NewError(PermissionDenied, errors.New("權限不足,存取遭拒"))
	fileTypeError     = NewError(FileTypeError, errors.New("上傳檔案類型不符合規範"))
	fileSizeError     = NewError(FileSizeError, errors.New("上傳檔案大小超過限制"))

	// Login
	LoginFailure     = NewError(1100, errors.New("登入失敗, 帳號或密碼錯誤"))
	LoginRoleFailure = NewError(1101, errors.New("登入身份錯誤"))
	LoginStatusFailure = NewError(1102, errors.New("帳號無法使用"))


	// Register
	RegisterFailure    = NewError(1400, errors.New("註冊失敗"))
	SendOTPFailure     = NewError(1401, errors.New("信箱驗證碼發送失敗"))
	OTPInvalid         = NewError(1402, errors.New("無效的信箱驗證碼"))
	NicknameDuplicate  = NewError(1405, errors.New("此名稱已有人使用，請試試其他名稱"))
	EmailDuplicate     = NewError(1406, errors.New("此郵件已有註冊紀錄，請返回登入"))
	AccountDuplicate   = NewError(1407, errors.New("該帳號已被使用"))

	// Course
	ActionNotExist = NewError(1500, errors.New("不存在的動作"))
)

type Error interface {
	Code() int
	Msg() string
}

type Handler interface {
	/** 公共 */
	Custom(code int, err error) Error
	Set(c *gin.Context, tag string, err error) Error
	// SystemError 9000 - 系統發生錯誤
	SystemError() Error
	DataNotFound() Error
	// DataAlreadyExists 9003 - 資料已存在
	DataAlreadyExists() Error
	// InvalidToken 9005 - 無效的token
	InvalidToken() Error
	// NewError(9006, errors.New("權限不足,存取遭拒"))
	PermissionDenied() Error
    // NewError(9007, errors.New("上傳檔案類型不符合規範"))
	FileTypeError() Error
    // NewError(9008, errors.New("上傳檔案大小超過限制"))
	FileSizeError() Error

	/** 註冊 */
	// NewError(1401, errors.New("手機驗證碼發送失敗"))
	SendOTPFailure() Error
	// NewError(1402, errors.New("無效的手機驗證碼"))
	OTPInvalid() Error
	// NewError(1405, errors.New("該暱稱已被使用"))
	NicknameDuplicate() Error
	// NewError(1406, errors.New("該信箱已被使用"))
	EmailDuplicate() Error

	/** 登入 */
	// LoginFailure 1100 - 登入失敗, 帳號或密碼錯誤
	LoginFailure() Error

	/** 課表 */
	// NewError(1500, errors.New("不存在的動作"))
	ActionNotExist() Error
}