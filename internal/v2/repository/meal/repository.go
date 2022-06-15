package meal

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"gorm.io/gorm"
)

type repository struct {
	gorm orm.Tool
}

func New(gormTool orm.Tool) Repository {
	return &repository{gorm: gormTool}
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.gorm.DB().Model(&model.Output{})
	//加入 id 篩選條件
	if input.DietID != nil {
		db = db.Where("diet_id = ?", *input.DietID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("%s %s", preload.OrderBy.OrderField, preload.OrderBy.OrderType))
				})
				continue
			}
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1)*input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Create(items []*model.Table) (err error) {
	err = r.gorm.DB().Model(&model.Table{}).Create(items).Error
	return err
}

func (r *repository) Update(items []*model.Table) (err error) {
	err = r.gorm.DB().Model(&model.Table{}).Save(&items).Error
	return err
}

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	db := r.gorm.DB()
	if input.ID != nil{
		db = db.Where("id = ?", *input.ID)
	}
	if input.DietID != nil{
		db = db.Where("diet_id = ?", *input.DietID)
	}
	err = db.Delete(&model.Table{}).Error
	return err
}
