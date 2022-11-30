package trainer_status_update_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_status_update_log"
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

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return *item.ID, err
}

func (r *repository) Creates(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	err = r.db.Model(&model.Table{}).Create(&items).Error
	return err
}
