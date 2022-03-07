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
		Take(&workoutLog, "id = ?", workoutLogID).Error; err != nil {
		return nil, err
	}
	return &workoutLog, nil
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

func (w *workoutLog) CalculateUserCourseStatistic(userID int64, workoutID int64) (*model.WorkoutLogCourseStatistic, error) {
	courseIDQuery := w.gorm.DB().Table("workout_logs").
		Select("courses.id").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("workouts.id = ?", workoutID).Limit(1)
	finishWorkoutCountQuery := w.gorm.DB().Table("workout_logs").
		Select("COUNT(DISTINCT workout_id)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = course_id AND workout_logs.user_id = ?", userID)
	totalFinishWorkoutCountQuery := w.gorm.DB().Table("workout_logs").
		Select("COUNT(*)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = course_id AND workout_logs.user_id = ?", userID)
	durationQuery := w.gorm.DB().Table("workout_logs").
		Select("SUM(duration)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = course_id AND workout_logs.user_id = ?", userID)
	var result model.WorkoutLogCourseStatistic
	if err := w.gorm.DB().Raw("SELECT (?) AS course_id, (?) AS finish_workout_count, (?) AS total_finish_workout_count, (?) AS duration",
		courseIDQuery,
		finishWorkoutCountQuery,
		totalFinishWorkoutCountQuery,
		durationQuery).
		Row().
		Scan(&result.CourseID, &result.FinishWorkoutCount, &result.TotalFinishWorkoutCount, &result.Duration); err != nil {
		return nil, err
	}
	return &result, nil
}

func (w *workoutLog) CalculateUserPlanStatistic(userID int64, workoutID int64) (*model.WorkoutLogPlanStatistic, error) {
	planIDQuery := w.gorm.DB().Table("workout_logs").
		Select("plans.id").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("workouts.id = ?", workoutID).Limit(1)
	finishWorkoutCountQuery := w.gorm.DB().Table("workout_logs").
		Select("COUNT(DISTINCT workout_id)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("plans.id = plan_id AND workout_logs.user_id = ?", userID)
	durationQuery := w.gorm.DB().Table("workout_logs").
		Select("SUM(duration)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Where("plans.id = plan_id AND workout_logs.user_id = ?", userID)
	var result model.WorkoutLogPlanStatistic
	if err := w.gorm.DB().Raw("SELECT (?) AS course_id, (?) AS finish_workout_count, (?) AS duration",
		planIDQuery,
		finishWorkoutCountQuery,
		durationQuery).
		Row().
		Scan(&result.PlanID, &result.FinishWorkoutCount, &result.Duration); err != nil {
		return nil, err
	}
	return &result, nil
}
