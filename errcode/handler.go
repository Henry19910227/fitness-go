package errcode

import (
	"errors"
	logger "github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type handle struct {
	logger logger.Logger
	errMap map[int]Error
}

func NewHandler() Handler {
	return &handle{}
}

func NewErrHandler(logger logger.Logger) Handler {
	errHandler := &handle{logger: logger}
	errHandler.errMap = make(map[int]Error)
	//公共錯誤碼
	errHandler.errMap[SystemError] = NewError(SystemError, errors.New("系統發生錯誤"))
	errHandler.errMap[UpdateError] = NewError(UpdateError, errors.New("更新失敗"))
	errHandler.errMap[DataNotFound] = NewError(DataNotFound, errors.New("查無資料"))
	errHandler.errMap[DataAlreadyExists] = NewError(DataAlreadyExists, errors.New("資料已存在"))
	errHandler.errMap[InvalidToken] = NewError(InvalidToken, errors.New("無效的token"))
	errHandler.errMap[PermissionDenied] = NewError(PermissionDenied, errors.New("權限不足,存取遭拒"))
	return errHandler
}


/** 公用 */
func (h *handle) Set(c *gin.Context, tag string, err error) Error {
	//Mysql錯誤碼映射
	if errors.Is(err, gorm.ErrRecordNotFound){
		return NewError(DataNotFound, errors.New("查無資料"))
	}
	if strings.Contains(err.Error(), "1062")  {
		return NewError(DataAlreadyExists, errors.New("資料已存在"))
	}
	//自定義錯誤碼映射
	code, _ := strconv.Atoi(err.Error())
	if _, ok := h.errMap[code]; ok {
		return h.errMap[code]
	}
	h.logger.Set(c, logger.Error, tag, systemError.Code(), err.Error())
	return systemError
}

func (h *handle) Custom(code int, err error) Error {
	return NewError(code, err)
}

func (h *handle) SystemError() Error {
	return systemError
}

func (h *handle) UpdateError() Error {
	return updateError
}

func (h *handle) DataNotFound() Error {
	return dataNotFound
}

func (h *handle) DataAlreadyExists() Error {
	return dataAlreadyExists
}

func (h *handle) InvalidThirdParty() Error {
	return invalidThirdParty
}

func (h *handle) InvalidToken() Error {
	return invalidToken
}

func (h handle) PermissionDenied() Error {
	return permissionDenied
}

func (h handle) FileTypeError() Error {
	return FileTypeError
}

func (h handle) FileSizeError() Error {
	return FileSizeError
}

/** 註冊 */
func (h handle) RegisterFailure() Error {
	return RegisterFailure
}

func (h handle) SendOTPFailure() Error {
	return SendOTPFailure
}

func (h handle) OTPInvalid() Error {
	return OTPInvalid
}

func (h handle) NicknameDuplicate() Error {
	return NicknameDuplicate
}

func (h handle) EmailDuplicate() Error {
	return EmailDuplicate
}

func (h handle) AccountDuplicate() Error {
	return AccountDuplicate
}

func (h handle) LoginFailure() Error {
	return LoginFailure
}

func (h handle) LoginRoleFailure() Error {
	return LoginRoleFailure
}

func (h handle) LoginStatusFailure() Error {
	return LoginStatusFailure
}

/** 課表 */
func (h handle) ActionNotExist() Error {
	return ActionNotExist
}
