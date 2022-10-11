package favorite_action

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_action"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return err
	}
	return err
}
