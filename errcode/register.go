package errcode

type register struct {
	common
}

func NewRegister() Register {
	register := &register{}
	return register
}
