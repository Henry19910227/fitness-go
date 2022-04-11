package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type workoutLog struct {
	gorm tool.Gorm
}

func NewWorkoutLog(gorm tool.Gorm) WorkoutLog {
	return &workoutLog{gorm: gorm}
}

func (w *workoutLog) CreateWorkoutLog(tx *gorm.DB, param *model.CreateWorkoutLogParam) (int64, error) {
	if param == nil {
		return 0, nil
	}
	db := w.gorm.DB()
	if tx != nil {
		db = tx
	}
	workoutLog := entity.WorkoutLog{
		UserID:    param.UserID,
		WorkoutID: param.WorkoutID,
		Duration:  param.Duration,
		Intensity: param.Intensity,
		Place:     param.Place,
		CreateAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&workoutLog).Error; err != nil {
		return 0, err
	}
	return workoutLog.ID, nil
}

func (w *workoutLog) FindWorkoutLog(workoutLogID int64) (*model.WorkoutLog, error) {
	var workoutLog model.WorkoutLog
	if err := w.gorm.DB().
		Preload("Workout").
		Take(&workoutLog, "id = ?", workoutLogID).Error; err != nil {
		return nil, err
	}
	return &workoutLog, nil
}

func (w *workoutLog) FindWorkoutLogsByDate(userID int64, startDate string, endDate string) ([]*model.WorkoutLog, error) {
	workoutLogs := make([]*model.WorkoutLog, 0)
	if err := w.gorm.DB().
		Preload("Workout").
		Where("user_id = ? AND create_at BETWEEN ? AND ?", userID, startDate, endDate).
		Find(&workoutLogs).Error; err != nil {
		return nil, err
	}
	return workoutLogs, nil
}

func (w *workoutLog) FindWorkoutLogsByPlanID(planID int64) ([]*model.WorkoutLog, error) {
	workoutLogs := make([]*model.WorkoutLog, 0)
	if err := w.gorm.DB().
		Table("workout_logs").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("plans.id = ?", planID).
		Find(&workoutLogs).Error; err != nil {
		return nil, err
	}
	return workoutLogs, nil
}

func (w *workoutLog) CalculateUserCourseStatistic(tx *gorm.DB, userID int64, workoutID int64) (*model.WorkoutLogCourseStatistic, error) {
	db := w.gorm.DB()
	if tx != nil {
		db = tx
	}
	var courseID int64
	if err := db.Table("workout_logs").
		Select("courses.id").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("workouts.id = ?", workoutID).Take(&courseID).Error; err != nil {
		return nil, err
	}
	finishWorkoutCountQuery := db.Table("workout_logs").
		Select("COUNT(DISTINCT workout_id)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = ? AND workout_logs.user_id = ?", courseID, userID)
	totalFinishWorkoutCountQuery := db.Table("workout_logs").
		Select("COUNT(*)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = ? AND workout_logs.user_id = ?", courseID, userID)
	durationQuery := db.Table("workout_logs").
		Select("SUM(duration)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = ? AND workout_logs.user_id = ?", courseID, userID)
	result := model.WorkoutLogCourseStatistic{
		CourseID: courseID,
	}
	if err := db.Raw("SELECT (?) AS finish_workout_count, (?) AS total_finish_workout_count, (?) AS duration",
		finishWorkoutCountQuery,
		totalFinishWorkoutCountQuery,
		durationQuery).
		Row().
		Scan(&result.FinishWorkoutCount, &result.TotalFinishWorkoutCount, &result.Duration); err != nil {
		return nil, err
	}
	return &result, nil
}

func (w *workoutLog) CalculateUserPlanStatistic(tx *gorm.DB, userID int64, workoutID int64) (*model.WorkoutLogPlanStatistic, error) {
	db := w.gorm.DB()
	if tx != nil {
		db = tx
	}
	var planID int64
	if err := db.Table("workout_logs").
		Select("plans.id").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("workouts.id = ?", workoutID).Take(&planID).Error; err != nil {
		return nil, err
	}
	finishWorkoutCountQuery := db.Table("workout_logs").
		Select("COUNT(DISTINCT workout_id)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("plans.id = ? AND workout_logs.user_id = ?", planID, userID)
	durationQuery := db.Table("workout_logs").
		Select("SUM(duration)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("plans.id = ? AND workout_logs.user_id = ?", planID, userID)
	result := model.WorkoutLogPlanStatistic{
		PlanID: planID,
	}
	if err := db.Raw("SELECT (?) AS finish_workout_count, (?) AS duration",
		finishWorkoutCountQuery,
		durationQuery).
		Row().
		Scan(&result.FinishWorkoutCount, &result.Duration); err != nil {
		return nil, err
	}
	return &result, nil
}
