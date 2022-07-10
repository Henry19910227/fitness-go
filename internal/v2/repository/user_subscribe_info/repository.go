package user_subscribe_info

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("id = ?", *input.UserID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}


