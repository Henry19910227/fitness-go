package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
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
		workout := entity.Workout{
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
	if param.SaleID != nil { selects = append(selects, "sale_id") }
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

func (c *course) FindCourseAmountByUserID(uid int64) (int, error) {
	var amount int
	if err := c.gorm.DB().
		Table("courses").
		Select("COUNT(*)").
		Where("user_id = ?", uid).
		Find(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}

func (c *course) FindCourseSummaries(param *model.FindCourseSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseSummaryEntity, error) {
	if param == nil {
		return nil, nil
	}
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 user_id 篩選條件
	if param.UID != nil {
		query += "AND courses.user_id = ? "
		params = append(params, param.UID)
	}
	//加入 status 篩選條件
	if param.Status != nil {
		query += "AND courses.course_status = ? "
		params = append(params, *param.Status)
	}
	var db *gorm.DB
	//基本查詢
	db = c.gorm.DB().
		Table("courses").
		Select("courses.id", "courses.course_status", "courses.category",
			"courses.schedule_type", "courses.`name`", "courses.cover",
			"courses.`level`", "courses.plan_count", "courses.workout_count",
			"IFNULL(sale.id,0)", "IFNULL(sale.type,0)", "IFNULL(sale.name,'')",
			"IFNULL(sale.twd,0)", "IFNULL(sale.identifier,'')",
			"IFNULL(sale.create_at,'')", "IFNULL(sale.update_at,'')",
			"trainers.user_id", "trainers.nickname", "trainers.avatar", "trainers.skill").
		Joins("INNER JOIN trainers ON courses.user_id = trainers.user_id").
		Joins("LEFT JOIN sale_items AS sale ON courses.sale_id = sale.id").
		Where(query, params...)
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	//查詢數據
	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}
	courses := make([]*model.CourseSummaryEntity, 0)
	for rows.Next() {
		var course model.CourseSummaryEntity
		if err := rows.Scan(&course.ID, &course.CourseStatus, &course.Category,
			&course.ScheduleType, &course.Name, &course.Cover, &course.Level,
			&course.PlanCount, &course.WorkoutCount,
			&course.Sale.ID, &course.Sale.Type, &course.Sale.Name, &course.Sale.Twd, &course.Sale.Identifier,
			&course.Sale.CreateAt, &course.Sale.UpdateAt,
			&course.Trainer.UserID, &course.Trainer.Nickname, &course.Trainer.Avatar, &course.Trainer.Skill); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (c *course) FindCourseProductSummaries(param model.FindCourseProductSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseProductSummary, error) {
	var db *gorm.DB
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 course_status 篩選條件
	query += "AND courses.course_status = ? "
	params = append(params, global.Sale)
	//加入 name 篩選條件
	if param.Name != nil {
		query += "AND courses.name LIKE ? "
		params = append(params, "%" + *param.Name + "%")
	}
	//加入 score 篩選條件
	if param.Score != nil {
		query += "AND FLOOR(review.score_total / review.amount) >= ? "
		params = append(params, *param.Score)
	}
	//加入 level 篩選條件
	if len(param.Level) > 0 {
		query += "AND courses.level IN ? "
		params = append(params, param.Level)
	}
	//加入 category 篩選條件
	if len(param.Category) > 0 {
		query += "AND courses.category IN ? "
		params = append(params, param.Category)
	}
	//加入 suit 篩選條件
	if len(param.Suit) > 0 {
		query += "AND courses.suit IN ? "
		params = append(params, param.Suit)
	}
	//加入 Equipment 篩選條件
	if len(param.Equipment) > 0 {
		query += "AND courses.equipment IN ? "
		params = append(params, param.Equipment)
	}
	//加入 Place 篩選條件
	if len(param.Equipment) > 0 {
		query += "AND courses.place IN ? "
		params = append(params, param.Place)
	}
	//加入 TrainTarget 篩選條件
	if len(param.TrainTarget) > 0 {
		query += "AND courses.train_target IN ? "
		params = append(params, param.TrainTarget)
	}
	//加入 BodyTarget 篩選條件
	if len(param.BodyTarget) > 0 {
		query += "AND courses.body_target IN ? "
		params = append(params, param.BodyTarget)
	}
	//加入 SaleType 篩選條件
	if len(param.SaleType) > 0 {
		query += "AND sale.type IN ? "
		params = append(params, param.SaleType)
	}
	//加入 TrainerSex 篩選條件
	if len(param.TrainerSex) > 0 {
		query += "AND users.sex IN ? "
		params = append(params, param.TrainerSex)
	}
	//加入 TrainerSkill 篩選條件
	if len(param.TrainerSkill) > 0 {
		query += "AND trainers.skill IN ? "
		params = append(params, param.TrainerSkill)
	}
	//基本查詢
	db = c.gorm.DB().
		Table("courses").
		Select("courses.id", "courses.course_status", "courses.category",
			"courses.schedule_type", "courses.`name`", "courses.cover",
			"courses.`level`", "courses.plan_count", "courses.workout_count",
			"IFNULL(sale.id,0)", "IFNULL(sale.type,0)", "IFNULL(sale.name,'')",
			"IFNULL(sale.twd,0)", "IFNULL(sale.identifier,'')",
			"IFNULL(sale.create_at,'')", "IFNULL(sale.update_at,'')",
		    "IFNULL(review.score_total,0)", "IFNULL(review.amount,0)",
			"trainers.user_id", "trainers.nickname", "trainers.avatar", "trainers.skill").
		Joins("INNER JOIN trainers ON courses.user_id = trainers.user_id").
		Joins("INNER JOIN users ON courses.user_id = users.id").
		Joins("LEFT JOIN sale_items AS sale ON courses.sale_id = sale.id").
		Joins("LEFT JOIN review_statistics AS review ON courses.id = review.course_id").
		Where(query, params...)

	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("courses.%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	//查詢數據
	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}
	courses := make([]*model.CourseProductSummary, 0)
	for rows.Next() {
		var course model.CourseProductSummary
		if err := rows.Scan(&course.ID, &course.CourseStatus, &course.Category,
			&course.ScheduleType, &course.Name, &course.Cover, &course.Level,
			&course.PlanCount, &course.WorkoutCount,
			&course.Sale.ID, &course.Sale.Type, &course.Sale.Name, &course.Sale.Twd, &course.Sale.Identifier,
			&course.Sale.CreateAt, &course.Sale.UpdateAt,
			&course.ReviewStatistic.ScoreTotal, &course.ReviewStatistic.Amount,
			&course.Trainer.UserID, &course.Trainer.Nickname, &course.Trainer.Avatar, &course.Trainer.Skill); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (c *course) FindCourseProduct(courseID int64) (*model.CourseProduct, error) {
	var course model.CourseProduct
	if err := c.gorm.DB().
		Preload("Trainer").
		Preload("Sale").
		Preload("Plans").
		Preload("Review").
		Where("id = ?", courseID).
		Take(&course).Error; err != nil {
			return nil, err
	}
	return &course, nil
}

func (c *course) FindCourseDetailByCourseID(courseID int64) (*model.CourseDetailEntity, error) {
	var course model.CourseDetailEntity
	if err := c.gorm.DB().
		Table("courses").
		Select("courses.id", "courses.course_status", "courses.category",
			"courses.schedule_type", "courses.`name`", "courses.cover", "courses.intro",
			"courses.food", "courses.level", "courses.suit", "courses.equipment",
			"courses.place", "courses.train_target", "courses.body_target", "courses.notice",
			"courses.plan_count", "courses.workout_count", "courses.create_at", "courses.update_at",
		    "IFNULL(sale.id,0)", "IFNULL(sale.type,0)", "IFNULL(sale.name,'')",
		    "IFNULL(sale.twd,0)", "IFNULL(sale.identifier,'')",
		    "IFNULL(sale.create_at,'')", "IFNULL(sale.update_at,'')",
			"trainers.user_id", "trainers.nickname", "trainers.avatar", "trainers.skill").
		Joins("INNER JOIN trainers ON courses.user_id = trainers.user_id").
		Joins("LEFT JOIN sale_items AS sale ON courses.sale_id = sale.id").
		Where("courses.id = ?", courseID).
		Row().
		Scan(&course.ID, &course.CourseStatus, &course.Category, &course.ScheduleType, &course.Name,
			&course.Cover, &course.Intro, &course.Food, &course.Level, &course.Suit, &course.Equipment,
			&course.Place, &course.TrainTarget, &course.BodyTarget, &course.Notice, &course.PlanCount,
			&course.WorkoutCount, &course.CreateAt, &course.UpdateAt,
		    &course.Sale.ID, &course.Sale.Type, &course.Sale.Name, &course.Sale.Twd, &course.Sale.Identifier,
		    &course.Sale.CreateAt, &course.Sale.UpdateAt,
			&course.Trainer.UserID, &course.Trainer.Nickname, &course.Trainer.Avatar, &course.Trainer.Skill); err != nil {
			return nil, err
	}
	return &course, nil
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

func (c *course) FindCourseByWorkoutSetID(setID int64, entity interface{}) error {
	var courseID int64
	if err := c.gorm.DB().
		Table("workout_sets AS `set`").
		Select("plans.course_id").
		Joins("INNER JOIN workouts ON `set`.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("`set`.id = ?", setID).
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