package feedback_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback_image"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
}

func (r *repository) Create(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	err = r.db.Model(&model.Table{}).Create(&items).Error
	if err != nil {
		return err
	}
	return err
}
