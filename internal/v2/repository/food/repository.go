package food

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// 1.Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	//加入 name 篩選條件
	if input.Name != nil {
		db = db.Where("foods.name LIKE ?", "%"+*input.Name+"%")
	}
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("foods.user_id = ?", *input.UserID)
	}
	//加入 source 篩選條件
	if input.Source != nil {
		db = db.Where("foods.source = ?", *input.Source)
	}
	//加入 is_disabled 篩選條件
	if input.Status != nil {
		db = db.Where("foods.status = ?", *input.Status)
	}
	//加入 is_delete 篩選條件
	if input.IsDeleted != nil {
		db = db.Where("foods.is_deleted = ?", *input.IsDeleted)
	}
	// 2.Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	// 4.Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	// 5.Count
	db = db.Count(&amount)
	// 6.Select
	db = db.Select("foods.*")
	// 7.Paging
	if input.Page != nil && input.Size != nil {
		db = db.Offset((*input.Page - 1) * *input.Size).Limit(*input.Size)
	} else if input.Page != nil {
		db = db.Offset(0)
	} else if input.Size != nil {
		db = db.Limit(*input.Size)
	}
	// 8.Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("foods.%s %s", input.OrderField, input.OrderType))
	}
	// 9.Custom Order
	if input.Orders != nil {
		for _, orderBy := range input.Orders {
			db = db.Order(orderBy.Value)
		}
	}
	// 查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return *item.ID, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}
