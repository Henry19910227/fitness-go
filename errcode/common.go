package errcode

type common struct {
}

func NewCommon() Common {
	return &common{}
}

func (handler *common) Custom(code int, err error) Error {
	return NewError(code, err)
}

func (handler *common) SystemError() Error {
	return systemError
}

func (handler *common) UpdateError() Error {
	return updateError
}

func (handler *common) DataNotFound() Error {
	return dataNotFound
}

func (handler *common) DataAlreadyExists() Error {
	return dataAlreadyExists
}

func (handler *common) InvalidThirdParty() Error {
	return InvalidThirdParty
}

func (handler *common) InvalidToken() Error {
	return InvalidToken
}
