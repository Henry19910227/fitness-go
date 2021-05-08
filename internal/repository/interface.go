package repository

type Admin interface {
	GetAdminID(email string, password string) (int64, error)
	GetAdmin(uid int64, entity interface{}) error
}

type User interface {
	CreateUser(accountType int, account string, nickname string, password string) (int64, error)
	FindUserIDByNickname(nickname string) (int64, error)
}