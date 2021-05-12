package repository

import (
	"github.com/Henry19910227/fitness-go/internal/dto/userdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type user struct {
	sso handler.SSO
	gorm  tool.Gorm
}

func NewUser(gormTool  tool.Gorm) User {
	return &user{gorm: gormTool}
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
		Birthday: "0000-01-01 00:00:00",
	}
	if err := u.gorm.DB().Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *user) UpdateUserByUID(uid int64, param *model.UpdateUserParam) error {
	var selects []interface{}
	if param.AccountType != nil { selects = append(selects, "account_type") }
	if param.Account != nil { selects = append(selects, "account") }
	if param.Password != nil { selects = append(selects, "password") }
	if param.DeviceToken != nil { selects = append(selects, "device_token") }
	if param.UserStatus != nil { selects = append(selects, "user_status") }
	if param.UserType != nil { selects = append(selects, "user_type") }
	if param.Email != nil { selects = append(selects, "email") }
	if param.Nickname != nil { selects = append(selects, "nickname") }
	if param.Sex != nil { selects = append(selects, "sex") }
	if param.Birthday != nil { selects = append(selects, "birthday") }
	if param.Height != nil { selects = append(selects, "height") }
	if param.Weight != nil { selects = append(selects, "weight") }
	if param.Experience != nil { selects = append(selects, "experience") }
	if param.Target != nil { selects = append(selects, "target") }
	//插入更新時間
	if param != nil {
		selects = append(selects, "update_at")
		var updateAt = time.Now().Format("2006-01-02 15:04:05")
		param.UpdateAt = &updateAt
	}
	if err := u.gorm.DB().
		Table("users").
		Where("id = ?", uid).
		Select("", selects...).
		Updates(param).Error; err != nil {
		return err
	}
	return nil
}

func (u *user) FindUserByUID(uid int64, entity interface{}) error {
	if err := u.gorm.DB().
		Model(&model.User{}).
		Where("id = ?", uid).
		Take(entity).Error; err != nil {
		return err
	}
	return nil
}

func (u *user) FindUserByAccountAndPassword(account string, password string, entity interface{}) error {
	if err := u.gorm.DB().
		Model(&model.User{}).
		Where("account = ? AND password = ?", account, password).
		Take(entity).Error; err != nil {
		return err
	}
	return nil
}

func (u *user) FindUserIDByNickname(nickname string) (int64, error) {
	var uid int64
	if err := u.gorm.DB().
		Table("users").
		Select("users.id").
		Where("users.nickname = ?", nickname).
		Take(&uid).Error; err != nil {
			return 0, err
	}
	return uid, nil
}

func (u *user) FindUserIDByEmail(email string) (int64, error) {
	var uid int64
	if err := u.gorm.DB().
		Table("users").
		Select("users.id").
		Where("users.account = ? OR users.email = ?", email, email).
		Take(&uid).Error; err != nil {
		return 0, err
	}
	return uid, nil
}
