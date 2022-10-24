package review_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
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

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	db := r.db
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	err = db.Delete(&model.Table{}).Error
	return err
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
