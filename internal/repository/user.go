package repository

import (
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type user struct {
	sso handler.SSO
	gorm  tool.Gorm
}

func NewUser(sso handler.SSO, gormTool  tool.Gorm) User {
	return &user{sso: sso, gorm: gormTool}
}

func (u *user) CreateUser(accountType int, account string, nickname string, password string) (int64, error) {
	user := model.User{
		AccountType: accountType,
		Account: account,
		Password: password,
		UserStatus: 1,
		UserType: 1,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		Nickname: nickname,
	}
	if err := u.gorm.DB().Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
