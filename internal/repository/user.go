package repository

import (
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type user struct {
	sso handler.SSO
	mysql tool.Mysql
	gorm  tool.Gorm
}

func NewUser(sso handler.SSO, mysqlTool tool.Mysql, gormTool  tool.Gorm) User {
	return &user{sso: sso, mysql: mysqlTool, gorm: gormTool}
}

func (u *user) CreateUser(accountType int, account string, nickname string, password string) (int64, error) {
	user := model.User{
		AccountType: accountType,
		Account: account,
		Nickname: nickname,
		Password: password,
	}
	if err := u.gorm.DB().Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
