package order

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order"
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

func (r *repository) Create(item *model.Table) (id string, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return "", err
	}
	return *item.ID, err
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("orders.id = ?", *input.ID)
	}
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("orders.user_id = ?", *input.UserID)
	}
	//加入 order_type 篩選條件
	if input.Type != nil {
		db = db.Where("orders.order_type = ?", *input.Type)
	}
	//加入 order_status 篩選條件
	if input.OrderStatus != nil {
		db = db.Where("orders.order_status = ?", *input.OrderStatus)
	}
	//加入 course_id 篩選條件
	if input.CourseID != nil {
		db = db.Joins("INNER JOIN order_courses ON orders.id = order_courses.order_id")
		db = db.Where("order_courses.course_id = ?", *input.CourseID)
	}
	//加入 OriginalTransactionID 篩選條件
	if input.OriginalTransactionID != nil {
		db = db.Joins("INNER JOIN receipts ON receipts.order_id = orders.id")
		db = db.Where("receipts.original_transaction_id = ?", *input.OriginalTransactionID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("orders.%s %s", preload.OrderBy.OrderField, preload.OrderBy.OrderType))
				})
				continue
			}
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("orders.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("orders.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}
