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

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	err = r.db.
		Where("user_id = ?", input.UserID).
		Where("action_id = ?", input.ActionID).
		Delete(&model.Table{}).Error
	return err
}
