package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type certificate struct {
	gorm tool.Gorm
}

func NewCertificate(gorm tool.Gorm) Certificate {
	return &certificate{gorm: gorm}
}

func (c *certificate) CreateCertificate(uid int64, name string, imageNamed string) (int64, error) {
	certificate := model.Certificate{
		UserID: uid,
		Name: name,
		Image: imageNamed,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := c.gorm.DB().Create(&certificate).Error; err != nil {
		return 0, err
	}
	return certificate.ID, nil
}

func (c *certificate) DeleteCertificateByID(cerID int64) error {
	if err := c.gorm.DB().
		Where("id = ?", cerID).
		Delete(&model.Certificate{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *certificate) FindCertificate(cerID int64, entity interface{}) error {
	if err := c.gorm.DB().
		Model(&model.Certificate{}).
		Where("id = ?", cerID).
		Take(entity).Error; err != nil{
		return err
	}
	return nil
}

