package repository

type Admin interface {
	GetAdminID(email string, password string) (int64, error)
	GetAdmin(uid int64, entity interface{}) error
}
