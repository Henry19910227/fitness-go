package errcode

import "github.com/gin-gonic/gin"

type mockHandle struct {
}

func NewMockHandler() Handler {
	return &mockHandle{}
}

func (m mockHandle) Custom(code int, err error) Error {
	return systemError
}

func (m mockHandle) Set(c *gin.Context, tag string, err error) Error {
	return systemError
}

func (m mockHandle) SystemError() Error {
	return systemError
}

func (m mockHandle) DataNotFound() Error {
	return systemError
}

func (m mockHandle) DataAlreadyExists() Error {
	return systemError
}

func (m mockHandle) InvalidToken() Error {
	return systemError
}

func (m mockHandle) PermissionDenied() Error {
	return systemError
}

func (m mockHandle) FileTypeError() Error {
	return systemError
}

func (m mockHandle) FileSizeError() Error {
	return systemError
}

func (m mockHandle) SendOTPFailure() Error {
	return systemError
}

func (m mockHandle) OTPInvalid() Error {
	return systemError
}

func (m mockHandle) NicknameDuplicate() Error {
	return systemError
}

func (m mockHandle) EmailDuplicate() Error {
	return systemError
}

func (m mockHandle) LoginFailure() Error {
	return systemError
}

func (m mockHandle) ActionNotExist() Error {
	return systemError
}
