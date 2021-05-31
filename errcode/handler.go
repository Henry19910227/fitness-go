package errcode

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

/** 公用 */
func (handler *handler) Custom(code int, err error) Error {
	return NewError(code, err)
}

func (handler *handler) SystemError() Error {
	return systemError
}

func (handler *handler) UpdateError() Error {
	return updateError
}

func (handler *handler) DataNotFound() Error {
	return dataNotFound
}

func (handler *handler) DataAlreadyExists() Error {
	return dataAlreadyExists
}

func (handler *handler) InvalidThirdParty() Error {
	return InvalidThirdParty
}

func (handler *handler) InvalidToken() Error {
	return InvalidToken
}

func (handler handler) PermissionDenied() Error {
	return PermissionDenied
}

func (handler handler) FileTypeError() Error {
	return FileTypeError
}

func (handler handler) FileSizeError() Error {
	return FileSizeError
}

/** 註冊 */
func (handler handler) RegisterFailure() Error {
	return RegisterFailure
}

func (handler handler) SendOTPFailure() Error {
	return SendOTPFailure
}

func (handler handler) OTPInvalid() Error {
	return OTPInvalid
}

func (handler handler) NicknameDuplicate() Error {
	return NicknameDuplicate
}

func (handler handler) EmailDuplicate() Error {
	return EmailDuplicate
}

func (handler handler) AccountDuplicate() Error {
	return AccountDuplicate
}

func (handler handler) LoginFailure() Error {
	return LoginFailure
}

func (handler handler) LoginRoleFailure() Error {
	return LoginRoleFailure
}

func (handler handler) LoginStatusFailure() Error {
	return LoginStatusFailure
}
