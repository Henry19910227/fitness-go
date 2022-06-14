package repository


import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
)

type admin struct {
	gorm tool.Gorm
}

func NewAdmin (gorm tool.Gorm) Admin {
	return &admin{gorm}
}

func (a *admin) GetAdminID(email string, password string) (int64, error) {
	var uid int64
	if err := a.gorm.DB().
		Table("admins").
		Select("id").
		Where("email = ? AND password = ?", email, password).
		Take(&uid).Error; err != nil {
		return 0, err
	}
	return uid, nil
}

func (a *admin) GetAdmin(uid int64, entity interface{}) error {
	if err := a.gorm.DB().
		Model(&model.Admin{}). //必須使用 Model 才能智能選擇字段
		Where("id = ?", uid).
		Take(entity).Error; err != nil {
		return err
	}
	return nil
}