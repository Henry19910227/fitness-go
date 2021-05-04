package errcode

type login struct {
	common
}

func NewLoginError() Login {
	return &login{}
}

func (l *login) LoginFailure() Error {
	return LoginFailure
}

func (l *login) LoginRoleFailure() Error {
	return LoginRoleFailure
}

func (l *login) LoginStatusFailure() Error {
	return LoginStatusFailure
}
