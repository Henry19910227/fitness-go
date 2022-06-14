package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"time"
)

type user struct {
	sso  handler.SSO
	gorm tool.Gorm
}

func NewUser(gormTool tool.Gorm) User {
	return &user{gorm: gormTool}
}

func (u *user) CreateUser(accountType int, account string, nickname string, password string) (int64, error) {
	var email string
	if accountType == 1 {
		email = account
	}
	user := model.User{
		AccountType: accountType,
		Account:     account,
		Password:    password,
		UserStatus:  1,
		UserType:    1,
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
		Nickname:    nickname,
		Email:       email,
		Birthday:    "0000-01-01 00:00:00",
	}
	if err := u.gorm.DB().Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *user) UpdateUserByUID(tx *gorm.DB, uid int64, param *model.UpdateUserParam) error {
	var selects []interface{}
	if param.AccountType != nil {
		selects = append(selects, "account_type")
	}
	if param.Account != nil {
		selects = append(selects, "account")
	}
	if param.Password != nil {
		selects = append(selects, "password")
	}
	if param.DeviceToken != nil {
		selects = append(selects, "device_token")
	}
	if param.UserStatus != nil {
		selects = append(selects, "user_status")
	}
	if param.UserType != nil {
		selects = append(selects, "user_type")
	}
	if param.Email != nil {
		selects = append(selects, "email")
	}
	if param.Nickname != nil {
		selects = append(selects, "nickname")
	}
	if param.Avatar != nil {
		selects = append(selects, "avatar")
	}
	if param.Sex != nil {
		selects = append(selects, "sex")
	}
	if param.Birthday != nil {
		selects = append(selects, "birthday")
	}
	if param.Height != nil {
		selects = append(selects, "height")
	}
	if param.Weight != nil {
		selects = append(selects, "weight")
	}
	if param.Experience != nil {
		selects = append(selects, "experience")
	}
	if param.Target != nil {
		selects = append(selects, "target")
	}
	//插入更新時間
	if param != nil {
		selects = append(selects, "update_at")
		var updateAt = time.Now().Format("2006-01-02 15:04:05")
		param.UpdateAt = &updateAt
	}
	db := u.gorm.DB()
	if tx != nil {
		db = tx
	}
	if err := db.
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

func (u *user) FindUsers(result interface{}, param *model.FinsUsersParam, orderBy *model.OrderBy, paging *model.PagingParam) (int, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 userID 篩選條件
	if param.UserID != nil {
		query += "AND id = ? "
		params = append(params, *param.UserID)
	}
	//加入 nickname 篩選條件
	if param.Name != nil {
		query += "AND nickname LIKE ? "
		params = append(params, "%"+*param.Name+"%")
	}
	//加入 email 篩選條件
	if param.Email != nil {
		query += "AND email LIKE ? "
		params = append(params, "%"+*param.Email+"%")
	}
	//加入 user_status 篩選條件
	if param.UserStatus != nil {
		query += "AND user_status = ? "
		params = append(params, *param.UserStatus)
	}
	//加入 user_type 篩選條件
	if param.UserType != nil {
		query += "AND user_type = ? "
		params = append(params, *param.UserType)
	}

	var db *gorm.DB
	var amount int64
	db = u.gorm.DB().Model(&model.User{}).Where(query, params...).Count(&amount)

	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	//查詢數據
	if err := db.Find(result).Error; err != nil {
		return 0, nil
	}
	return int(amount), nil
}
