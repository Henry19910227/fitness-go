package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type course struct {
	gorm tool.Gorm
}

func NewCourse(gorm tool.Gorm) Course {
	return &course{gorm: gorm}
}

func (c *course) CreateCourse(uid int64, param *model.CreateCourseParam) (int64, error) {
	course := entity.Course{
		UserID:       uid,
		Name:         param.Name,
		Level:        param.Level,
		Category:     param.Category,
		ScheduleType: int(global.MultipleScheduleType),
		CourseStatus: int(global.Preparing),
		SaleType:     int(global.SaleTypeNone),
		CreateAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := c.gorm.DB().Create(&course).Error; err != nil {
		return 0, err
	}
	return course.ID, nil
}

func (c *course) CreateSingleWorkoutCourse(uid int64, param *model.CreateCourseParam) (int64, error) {
	course := entity.Course{
		UserID:       uid,
		Name:         param.Name,
		Level:        param.Level,
		Category:     param.Category,
		ScheduleType: int(global.SingleScheduleType),
		CourseStatus: int(global.Preparing),
		SaleType:     int(global.SaleTypeNone),
		CreateAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := c.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建課表
		if err := tx.Create(&course).Error; err != nil {
			return err
		}
		//創建計畫
		plan := model.Plan{
			CourseID: course.ID,
			Name:     course.Name + "計畫",
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := tx.Create(&plan).Error; err != nil {
			return err
		}
		//創建訓練
		workout := entity.Workout{
			PlanID:   plan.ID,
			Name:     "單一訓練",
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
				"plan_count":    plantCountQuery,
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

func (c *course) UpdateCourseByID(tx *gorm.DB, courseID int64, param *model.UpdateCourseParam) error {
	var selects []interface{}
	if param.CourseStatus != nil {
		selects = append(selects, "course_status")
	}
	if param.Category != nil {
		selects = append(selects, "category")
	}
	if param.ScheduleType != nil {
		selects = append(selects, "schedule_type")
	}
	if param.SaleType != nil {
		selects = append(selects, "sale_type")
		selects = append(selects, "sale_id")
	}
	if param.Name != nil {
		selects = append(selects, "name")
	}
	if param.Cover != nil {
		selects = append(selects, "cover")
	}
	if param.Intro != nil {
		selects = append(selects, "intro")
	}
	if param.Food != nil {
		selects = append(selects, "food")
	}
	if param.Level != nil {
		selects = append(selects, "level")
	}
	if param.Suit != nil {
		selects = append(selects, "suit")
	}
	if param.Equipment != nil {
		selects = append(selects, "equipment")
	}
	if param.Place != nil {
		selects = append(selects, "place")
	}
	if param.TrainTarget != nil {
		selects = append(selects, "train_target")
	}
	if param.BodyTarget != nil {
		selects = append(selects, "body_target")
	}
	if param.Notice != nil {
		selects = append(selects, "notice")
	}
	//插入更新時間
	if param != nil {
		selects = append(selects, "update_at")
		var updateAt = time.Now().Format("2006-01-02 15:04:05")
		param.UpdateAt = &updateAt
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	if err := db.
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

func (c *course) FindCourseSummaries(param *model.FindCourseSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseSummary, error) {
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
		Select("courses.id AS id", "courses.user_id AS user_id", "courses.sale_id AS sale_id",
			"courses.sale_type AS sale_type", "courses.course_status AS course_status", "courses.category AS category",
			"courses.schedule_type AS schedule_type", "courses.name AS name", "courses.cover AS cover",
			"courses.`level` AS level", "courses.plan_count AS plan_count", "courses.workout_count AS workout_count").
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel")
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	var courses []*model.CourseSummary
	if err := db.Where(query, params...).Find(&courses).Error; err != nil {
		return nil, err
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
	//加入 教練ID 篩選條件
	if param.UserID != nil {
		query += "AND courses.user_id = ? "
		params = append(params, *param.UserID)
	}
	//加入 name 篩選條件
	if param.Name != nil {
		query += "AND courses.name LIKE ? "
		params = append(params, "%"+*param.Name+"%")
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
		query += "AND courses.suit LIKE ? "
		params = append(params, "%"+transformFilterParams(param.Suit)+"%")
	}
	//加入 Equipment 篩選條件
	if len(param.Equipment) > 0 {
		query += "AND courses.equipment LIKE ? "
		params = append(params, "%"+transformFilterParams(param.Equipment)+"%")
	}
	//加入 Place 篩選條件
	if len(param.Place) > 0 {
		query += "AND courses.place LIKE ? "
		params = append(params, "%"+transformFilterParams(param.Place)+"%")
	}
	//加入 TrainTarget 篩選條件
	if len(param.TrainTarget) > 0 {
		query += "AND courses.train_target LIKE ? "
		params = append(params, "%"+transformFilterParams(param.TrainTarget)+"%")
	}
	//加入 BodyTarget 篩選條件
	if len(param.BodyTarget) > 0 {
		query += "AND courses.body_target LIKE ? "
		params = append(params, "%"+transformFilterParams(param.BodyTarget)+"%")
	}
	//加入 SaleType 篩選條件
	if len(param.SaleType) > 0 {
		query += "AND courses.sale_type IN ? "
		params = append(params, param.SaleType)
	}
	//加入 TrainerSex 篩選條件
	if len(param.TrainerSex) > 0 {
		query += "AND users.sex IN ? "
		params = append(params, param.TrainerSex)
	}
	//加入 TrainerSkill 篩選條件
	if len(param.TrainerSkill) > 0 {
		query += "AND trainers.skill LIKE ? "
		params = append(params, "%"+transformFilterParams(param.TrainerSkill)+"%")
	}
	//基本查詢
	db = c.gorm.DB().
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel").
		Preload("Review").
		Joins("INNER JOIN users ON courses.user_id = users.id").
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
	var courses []*model.CourseProductSummary
	if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *course) FindCourseProductCount(param model.FindCourseProductCountParam) (int, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 course_status 篩選條件
	query += "AND courses.course_status = ? "
	params = append(params, global.Sale)
	//加入 教練ID 篩選條件
	if param.UserID != nil {
		query += "AND courses.user_id = ? "
		params = append(params, *param.UserID)
	}
	//加入 name 篩選條件
	if param.Name != nil {
		query += "AND courses.name LIKE ? "
		params = append(params, "%"+*param.Name+"%")
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
		query += "AND courses.suit LIKE ? "
		params = append(params, "%"+transformFilterParams(param.Suit)+"%")
	}
	//加入 Equipment 篩選條件
	if len(param.Equipment) > 0 {
		query += "AND courses.equipment LIKE ? "
		params = append(params, "%"+transformFilterParams(param.Equipment)+"%")
	}
	//加入 Place 篩選條件
	if len(param.Place) > 0 {
		query += "AND courses.place LIKE ? "
		params = append(params, "%"+transformFilterParams(param.Place)+"%")
	}
	//加入 TrainTarget 篩選條件
	if len(param.TrainTarget) > 0 {
		query += "AND courses.train_target LIKE ? "
		params = append(params, "%"+transformFilterParams(param.TrainTarget)+"%")
	}
	//加入 BodyTarget 篩選條件
	if len(param.BodyTarget) > 0 {
		query += "AND courses.body_target LIKE ? "
		params = append(params, "%"+transformFilterParams(param.BodyTarget)+"%")
	}
	//加入 SaleType 篩選條件
	if len(param.SaleType) > 0 {
		query += "AND courses.sale_type IN ? "
		params = append(params, param.SaleType)
	}
	//加入 TrainerSex 篩選條件
	if len(param.TrainerSex) > 0 {
		query += "AND users.sex IN ? "
		params = append(params, param.TrainerSex)
	}
	//加入 TrainerSkill 篩選條件
	if len(param.TrainerSkill) > 0 {
		query += "AND trainers.skill LIKE ? "
		params = append(params, "%"+transformFilterParams(param.TrainerSkill)+"%")
	}
	//基本查詢
	var count int64
	if err := c.gorm.DB().
		Table("courses").
		Select("courses.id", "courses.course_status", "courses.category",
			"courses.schedule_type", "courses.`name`", "courses.cover",
			"courses.`level`", "courses.plan_count", "courses.workout_count",
			"IFNULL(sale.id,0)", "IFNULL(sale.type,0)", "IFNULL(sale.name,'')",
			"IFNULL(sale.twd,0)", "IFNULL(sale.product_id,'')",
			"IFNULL(sale.create_at,'')", "IFNULL(sale.update_at,'')",
			"IFNULL(review.score_total,0)", "IFNULL(review.amount,0)",
			"trainers.user_id", "trainers.nickname", "trainers.avatar", "trainers.skill").
		Joins("INNER JOIN trainers ON courses.user_id = trainers.user_id").
		Joins("INNER JOIN users ON courses.user_id = users.id").
		Joins("LEFT JOIN sale_items AS sale ON courses.sale_id = sale.id").
		Joins("LEFT JOIN review_statistics AS review ON courses.id = review.course_id").
		Where(query, params...).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (c *course) FindProgressCourseAssetSummaries(userID int64, paging *model.PagingParam) ([]*model.CourseAssetSummary, error) {
	var courses []*model.CourseAssetSummary
	db := c.gorm.DB().
		//Select("courses.id AS id", "courses.user_id AS user_id", "courses.sale_id AS sale_id",
		//	"courses.sale_type AS sale_type", "courses.course_status AS course_status", "courses.category AS category",
		//	"courses.schedule_type AS schedule_type", "courses.name AS name", "courses.cover AS cover", "courses.level AS level",
		//	"courses.plan_count AS plan_count", "courses.workout_count AS workout_count", "stat.finish_workout_count AS finish_workout_count",
		//	"stat.duration AS duration").
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel").
		Preload("Review").
		Joins("LEFT JOIN user_course_statistics AS stat ON courses.id = stat.course_id AND stat.user_id = ?", userID).
		Order("stat.update_at DESC").
		Where("stat.user_id = ?", userID)
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *course) FindChargeCourseAssetSummaries(userID int64, paging *model.PagingParam) ([]*model.CourseAssetSummary, error) {
	var courses []*model.CourseAssetSummary
	db := c.gorm.DB().
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel").
		Preload("Review").
		Joins("INNER JOIN user_course_assets AS asset ON courses.id = asset.course_id AND asset.user_id = ?", userID).
		Order("asset.create_at DESC").
		Where("asset.user_id = ? AND asset.available = ? AND courses.sale_type = ?", userID, 1, global.SaleTypeCharge)
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *course) FindProgressCourseAssetCount(userID int64) (int, error) {
	var count int64
	if err := c.gorm.DB().
		Table("courses").
		Joins("LEFT JOIN user_course_statistics AS stat ON courses.id = stat.course_id AND stat.user_id = ?", userID).
		Where("stat.user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (c *course) FindChargeCourseAssetCount(userID int64) (int, error) {
	var count int64
	if err := c.gorm.DB().
		Table("courses").
		Joins("INNER JOIN user_course_assets AS asset ON courses.id = asset.course_id AND asset.user_id = ?", userID).
		Order("asset.create_at DESC").
		Where("asset.user_id = ? AND asset.available = ? AND courses.sale_type = ?", userID, 1, global.SaleTypeCharge).
		Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}

func (c *course) FindCourseProduct(courseID int64) (*model.CourseProduct, error) {
	var course model.CourseProduct
	if err := c.gorm.DB().
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel").
		Preload("Review").
		Where("id = ?", courseID).
		Take(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (c *course) FindCourseAsset(courseID int64, userID int64) (*model.CourseAsset, error) {
	var course model.CourseAsset
	if err := c.gorm.DB().
		Select("courses.id AS id", "courses.user_id AS user_id", "courses.sale_id AS sale_id",
			"courses.sale_type AS sale_type", "courses.course_status AS course_status", "courses.category AS category",
			"courses.schedule_type AS schedule_type", "courses.name AS name", "courses.cover AS cover", "courses.level AS level",
			"courses.plan_count AS plan_count", "courses.workout_count AS workout_count",
			"IFNULL(stat.finish_workout_count, 0) AS finish_workout_count", "IFNULL(stat.duration, 0) AS duration").
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel").
		Joins("LEFT JOIN user_course_statistics AS stat ON courses.id = stat.course_id AND stat.user_id = ?", userID).
		Where("courses.id = ?", courseID).
		Take(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (c *course) FindCourseByCourseID(courseID int64) (*model.Course, error) {
	var course model.Course
	if err := c.gorm.DB().
		Preload("Trainer").
		Preload("Sale").
		Preload("Sale.ProductLabel").
		Take(&course, courseID).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (c *course) FindCourseByID(tx *gorm.DB, courseID int64, entity interface{}) error {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	if err := db.
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
	return c.FindCourseByID(nil, courseID, entity)
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
	return c.FindCourseByID(nil, courseID, entity)
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
	return c.FindCourseByID(nil, courseID, entity)
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
	return c.FindCourseByID(nil, courseID, entity)
}

func (c *course) DeleteCourseByID(courseID int64) error {
	if err := c.gorm.DB().
		Where("id = ?", courseID).
		Delete(&entity.Course{}).Error; err != nil {
		return err
	}
	return nil
}

func transformFilterParams(params []int) string {
	var result string
	for _, param := range params {
		result += strconv.Itoa(param) + ","
	}
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}
