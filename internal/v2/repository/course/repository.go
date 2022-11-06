package course

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
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

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	err = r.db.Where("id = ?", input.ID).Delete(&model.Table{}).Error
	return err
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("courses.id = ?", *input.ID)
	}
	//加入 plan_id 篩選條件
	if input.PlanID != nil {
		db = db.Joins("INNER JOIN plans ON courses.id = plans.course_id")
		db = db.Where("plans.id = ?", *input.PlanID)
	}
	//加入 workout_id 篩選條件
	if input.WorkoutID != nil {
		db = db.Joins("INNER JOIN plans ON courses.id = plans.course_id")
		db = db.Joins("INNER JOIN workouts ON plans.id = workouts.plan_id")
		db = db.Where("workouts.id = ?", *input.WorkoutID)
	}
	//加入 workout_set_id 篩選條件
	if input.WorkoutSetID != nil {
		db = db.Joins("INNER JOIN plans ON courses.id = plans.course_id")
		db = db.Joins("INNER JOIN workouts ON plans.id = workouts.plan_id")
		db = db.Joins("INNER JOIN workout_sets ON workouts.id = workout_sets.workout_id")
		db = db.Where("workout_sets.id = ?", *input.WorkoutSetID)
	}
	// Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	// Select
	db = db.Select("courses.*")
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// id 篩選條件
	if input.ID != nil {
		db = db.Where("courses.id = ?", *input.ID)
	}
	// user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("courses.user_id = ?", *input.UserID)
	}
	// name 篩選條件
	if input.Name != nil {
		db = db.Where("courses.name LIKE ?", "%"+*input.Name+"%")
	}
	// course_status 篩選條件
	if input.CourseStatus != nil {
		db = db.Where("courses.course_status = ?", *input.CourseStatus)
	}
	// trainer_status 篩選條件
	if input.SaleType != nil {
		db = db.Where("courses.sale_type = ?", *input.SaleType)
	}
	// schedule_type 篩選條件
	if input.ScheduleType != nil {
		db = db.Where("courses.schedule_type = ?", *input.ScheduleType)
	}
	// Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("courses.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("courses.%s %s", input.OrderField, input.OrderType))
	}
	// Custom Order
	if input.Orders != nil {
		for _, orderBy := range input.Orders {
			db = db.Order(orderBy.Value)
		}
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Updates(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Save(&items).Error
	return err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}

func (r *repository) UpdateSaleID(id int64, saleItemID *int64) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", id).Updates(map[string]interface{}{"sale_id": saleItemID}).Error
	return err
}
