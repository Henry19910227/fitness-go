package action

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
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
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// ID List 篩選條件
	if len(input.IDs) > 0 {
		db = db.Where("id IN (?)", input.IDs)
	}
	// Name 篩選條件
	if input.Name != nil {
		db = db.Where("name LIKE ?", "%"+*input.Name+"%")
	}
	// Source 篩選條件
	if len(input.SourceList) > 0 {
		db = db.Where("source IN (?)", input.SourceList)
	}
	// Category 篩選條件
	if len(input.CategoryList) > 0 {
		db = db.Where("category IN (?)", input.CategoryList)
	}
	// Body 篩選條件
	if len(input.BodyList) > 0 {
		db = db.Where("body IN (?)", input.BodyList)
	}
	// Equipment 篩選條件
	if len(input.EquipmentList) > 0 {
		db = db.Where("equipment IN (?)", input.EquipmentList)
	}
	// UserID 篩選條件
	if input.UserID != nil {
		db = db.Where("user_id = ? OR user_id IS NULL", *input.UserID)
	}
	// Type 篩選條件
	if input.Type != nil {
		db = db.Where("type = ?", *input.Type)
	}
	// Source 篩選條件
	if input.Source != nil {
		db = db.Where("source = ?", *input.Source)
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

func (r *repository) UserActionList(input *model.UserActionListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// UserID 篩選條件
	if input.UserID != nil {
		db = db.Where("user_id = ? OR user_id IS NULL", *input.UserID)
	}
	// Name 篩選條件
	if input.Name != nil {
		db = db.Where("name LIKE ?", "%"+*input.Name+"%")
	}
	// Source 篩選條件
	if len(input.Source) > 0 {
		db = db.Where("source IN (?)", input.Source)
	}
	// Category 篩選條件
	if len(input.Category) > 0 {
		db = db.Where("category IN (?)", input.Category)
	}
	// Body 篩選條件
	if len(input.Body) > 0 {
		db = db.Where("body IN (?)", input.Body)
	}
	// Equipment 篩選條件
	if len(input.Equipment) > 0 {
		db = db.Where("equipment IN (?)", input.Equipment)
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
