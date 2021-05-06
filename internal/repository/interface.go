package repository

import "github.com/Henry19910227/fitness-go/internal/dto"

type Admin interface {
	GetAdminID(email string, password string) (int64, error)
	GetAdmin(uid int64, entity interface{}) error
}

type User interface {
	CreateUser(param dto.CreateUser) (int64, error)
}