package user

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
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
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//加入 is_deleted 篩選條件
	if input.IsDeleted != nil {
		db = db.Where("is_deleted = ?", *input.IsDeleted)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return *item.ID, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	//加入 account 篩選條件
	if input.Account != nil {
		db = db.Where("account = ?", *input.Account)
	}
	//加入 password 篩選條件
	if input.Password != nil {
		db = db.Where("password = ?", *input.Password)
	}
	//加入 nickname 篩選條件
	if input.Nickname != nil {
		db = db.Where("nickname = ?", *input.Nickname)
	}
	//加入 is_deleted 篩選條件
	if input.IsDeleted != nil {
		db = db.Where("is_deleted = ?", *input.IsDeleted)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}
