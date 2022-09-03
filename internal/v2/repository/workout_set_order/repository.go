package workout_set_order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_order"
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
	err = r.db.Model(&model.Table{}).Create(items).Error
	return err
}

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	err = r.db.Where("workout_id = ?", input.WorkoutID).Delete(&model.Table{}).Error
	return err
}
