package errcode

type errs struct {
	ErrCode int
	Err     error
}

func NewError(code int, err error) Error {
	return &errs{code, err}
}

func (e *errs) Code() int {
	return e.ErrCode
}

func (e *errs) Msg() string {
	return e.Err.Error()
}
