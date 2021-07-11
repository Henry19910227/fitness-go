package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type course struct {
	gorm  tool.Gorm
}

func NewCourse(gorm tool.Gorm) Course {
	return &course{gorm: gorm}
}

func (c *course) CreateCourse(uid int64, param *model.CreateCourseParam) (int64, error) {
	course := model.Course{
		UserID: uid,
		Name: param.Name,
		Level: param.Level,
		Category: param.Category,
		ScheduleType: 2,
		CourseStatus: 1,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := c.gorm.DB().Create(&course).Error; err != nil {
		return 0, err
	}
	return course.ID, nil
}

func (c *course) CreateSingleWorkoutCourse(uid int64, param *model.CreateCourseParam) (int64, error) {
	course := model.Course{
		UserID: uid,
		Name: param.Name,
		Level: param.Level,
		Category: param.Category,
		ScheduleType: 1,
		CourseStatus: 1,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := c.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建課表
		if err := tx.Create(&course).Error; err != nil {
			return err
		}
		//創建計畫
		plan := model.Plan{
			CourseID: course.ID,
			Name: course.Name + "計畫",
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := tx.Create(&plan).Error; err != nil {
			return err
		}
		//創建訓練
		workout := model.Workout{
			PlanID: plan.ID,
			Name: "單一訓練",
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := tx.Create(&workout).Error; err != nil {
			return err
		}
		plantCountQuery := tx.Table("plans").
			Select("COUNT(*) AS plan_count").
			Where("course_id = ?", course.ID)
		workoutCountQuery := tx.Table("workouts").
			Select("COUNT(*) AS workout_count").
			Where("plan_id = ?", plan.ID)
		//更新課表的計畫與訓練數量
		if err := tx.Table("courses").
			Where("id = ?", course.ID).
			Updates(map[string]interface{}{
				"plan_count": plantCountQuery,
				"workout_count": workoutCountQuery,
		}).Error; err != nil {
				return err
		}
		//更新計畫的訓練數量
		if err := tx.Table("plans").
			Where("id = ?", plan.ID).
			Updates(map[string]interface{}{
				"workout_count": workoutCountQuery,
			}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return course.ID, nil
}


func (c *course) UpdateCourseByID(courseID int64, param *model.UpdateCourseParam) error {
	var selects []interface{}
	if param.CourseStatus != nil { selects = append(selects, "course_status") }
	if param.Category != nil { selects = append(selects, "category") }
	if param.ScheduleType != nil { selects = append(selects, "schedule_type") }
	if param.SaleType != nil { selects = append(selects, "sale_type") }
	if param.Price != nil { selects = append(selects, "price") }
	if param.Name != nil { selects = append(selects, "name") }
	if param.Cover != nil { selects = append(selects, "cover") }
	if param.Intro != nil { selects = append(selects, "intro") }
	if param.Food != nil { selects = append(selects, "food") }
	if param.Level != nil { selects = append(selects, "level") }
	if param.Suit != nil { selects = append(selects, "suit") }
	if param.Equipment != nil { selects = append(selects, "equipment") }
	if param.Place != nil { selects = append(selects, "place") }
	if param.TrainTarget != nil { selects = append(selects, "train_target") }
	if param.BodyTarget != nil { selects = append(selects, "body_target") }
	if param.Notice != nil { selects = append(selects, "notice") }
	//插入更新時間
	if param != nil {
		selects = append(selects, "update_at")
		var updateAt = time.Now().Format("2006-01-02 15:04:05")
		param.UpdateAt = &updateAt
	}
	if err := c.gorm.DB().
		Table("courses").
		Where("id = ?", courseID).
		Select("", selects...).
		Updates(param).Error; err != nil {
			return err
	}
	return nil
}

func (c *course) FindCourses(uid int64, entity interface{}, status *int) error {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 user_id 篩選條件
	query += "AND user_id = ? "
	params = append(params, uid)
	//加入 status 篩選條件
	if status != nil {
		query += "AND course_status = ? "
		params = append(params, *status)
	}
	if err := c.gorm.DB().
		Model(entity).
		Where(query, params...).
		Find(entity).Error; err != nil {
			return err
	}
	return nil
}

func (c *course) FindCourseByID(courseID int64, entity interface{}) error {
	if err := c.gorm.DB().
		Model(&model.Course{}).
		Where("id = ?", courseID).
		Take(entity).Error; err != nil {
		return err
	}
	return nil
}

func (c *course) FindCourseByPlanID(planID int64, entity interface{}) error {
	var courseID int64
	if err := c.gorm.DB().
		Table("plans").
		Select("course_id").
		Where("id = ?", planID).
		Take(&courseID).Error; err != nil {
			return err
	}
	return c.FindCourseByID(courseID, entity)
}

func (c *course) FindCourseByWorkoutID(workoutID int64, entity interface{}) error {
	var courseID int64
	if err := c.gorm.DB().
		Table("workouts").
		Select("plans.course_id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("workouts.id = ?", workoutID).
		Take(&courseID).Error; err != nil {
		return err
	}
	return c.FindCourseByID(courseID, entity)
}

func (c *course) FindCourseByActionID(actionID int64, entity interface{}) error {
	var courseID int64
	if err := c.gorm.DB().
		Table("actions").
		Select("course_id").
		Where("id = ?", actionID).
		Take(&courseID).Error; err != nil {
		return err
	}
	return c.FindCourseByID(courseID, entity)
}

func (c *course) DeleteCourseByID(courseID int64) error {
	if err := c.gorm.DB().
		Where("id = ?", courseID).
		Delete(&model.Course{}).Error; err != nil {
		return err
	}
	return nil
}