package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
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
		ScheduleType: param.ScheduleType,
		CourseStatus: 1,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := c.gorm.DB().Create(&course).Error; err != nil {
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

func (c *course) FindCourseStatusByPlanID(planID int64) (int, error) {
	var status int
	if err := c.gorm.DB().
		Table("courses").
		Select("course_status").
		Joins("INNER JOIN plans ON courses.id = plans.course_id").
		Where("plans.id = ?", planID).
		Take(&status).Error; err != nil {
		return 0, err
	}
	return status, nil
}

func (c *course) DeleteCourseByID(courseID int64) error {
	if err := c.gorm.DB().
		Where("id = ?", courseID).
		Delete(&model.Course{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *course) CheckCourseExistByIDAndUID(courseID int64, uid int64) (bool, error) {
	var result int
	if err := c.gorm.DB().
		Table("courses").
		Select("1").
		Where("id = ? AND user_id = ?", courseID, uid).
		Find(&result).Error; err != nil {
			return false, err
	}
	return result > 0, nil
}