package repository

import "github.com/Henry19910227/fitness-go/internal/model"

type Admin interface {
	GetAdminID(email string, password string) (int64, error)
	GetAdmin(uid int64, entity interface{}) error
}

type User interface {
	CreateUser(accountType int, account string, nickname string, password string) (int64, error)
	FindUserByEmailAndPassword(email string, password string) (*model.User, error)
	FindUserIDByNickname(nickname string) (int64, error)
	FindUserIDByEmail(email string) (int64, error)
}