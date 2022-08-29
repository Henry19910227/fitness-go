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

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
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
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
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
	// id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	// user_id 篩選條件
	if input.ID != nil {
		db = db.Where("user_id = ?", *input.UserID)
	}
	// name 篩選條件
	if input.Name != nil {
		db = db.Where("name LIKE ?", "%"+*input.Name+"%")
	}
	// course_status 篩選條件
	if input.CourseStatus != nil {
		db = db.Where("course_status = ?", *input.CourseStatus)
	}
	// trainer_status 篩選條件
	if input.SaleType != nil {
		db = db.Where("sale_type = ?", *input.SaleType)
	}
	if len(input.IDs) > 0 {
		db = db.Where("id IN (?)", input.IDs)
	}
	if len(input.IgnoredCourseStatus) > 0 {
		db = db.Where("course_status NOT IN (?)", input.IgnoredCourseStatus)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
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

func (r *repository) FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{}).Joins("INNER JOIN favorite_courses ON courses.id = favorite_courses.course_id")
	// id 篩選條件
	if input.UserID != nil {
		db = db.Where("favorite_courses.user_id = ?", *input.UserID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
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
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("favorite_courses.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) ProgressList(input *model.ProgressListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// JOIN
	db = db.Joins("LEFT JOIN user_course_statistics AS statistic ON courses.id = statistic.course_id AND statistic.user_id = ?", input.UserID)
	// user_id 篩選條件
	db = db.Where("statistic.user_id = ?", input.UserID)
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
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
		db = db.Order(fmt.Sprintf("statistic.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) ChargeList(input *model.ChargeListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// JOIN
	db = db.Joins("INNER JOIN user_course_assets AS asset ON courses.id = asset.course_id AND asset.user_id = ?", input.UserID)
	// user_id 篩選條件
	db = db.Where("asset.user_id = ?", input.UserID)
	// available 篩選條件
	db = db.Where("asset.available = ?", 1)
	// sale_type 篩選條件
	db = db.Where("courses.sale_type = ?", model.SaleTypeCharge)
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
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
		db = db.Order(fmt.Sprintf("asset.%s %s", input.OrderField, input.OrderType))
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
