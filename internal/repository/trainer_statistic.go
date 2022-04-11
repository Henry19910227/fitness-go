package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type trainerStatistic struct {
	gorm tool.Gorm
}

func NewTrainerStatistic(gorm tool.Gorm) TrainerStatistic {
	return &trainerStatistic{gorm: gorm}
}

func (t *trainerStatistic) SaveTrainerStatistic(tx *gorm.DB, userID int64, param *model.SaveTrainerStatisticParam) error {
	if param == nil {
		return nil
	}
	db := t.gorm.DB()
	if tx != nil {
		db = tx
	}
	assignmentColumns := []string{"update_at"}
	var studentCount int
	if param.StudentCount != nil {
		studentCount = *param.StudentCount
		assignmentColumns = append(assignmentColumns, "student_count")
	}
	var courseCount int
	if param.CourseCount != nil {
		courseCount = *param.CourseCount
		assignmentColumns = append(assignmentColumns, "course_count")
	}
	var reviewScore float64
	if param.ReviewScore != nil {
		reviewScore = *param.ReviewScore
		assignmentColumns = append(assignmentColumns, "review_score")
	}
	stat := entity.TrainerStatistic{
		UserID:       userID,
		StudentCount: studentCount,
		CourseCount:  courseCount,
		ReviewScore:  reviewScore,
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns(assignmentColumns),
	}).Create(&stat).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerStatistic) CalculateTrainerStudentCount(tx *gorm.DB, userID int64) (int, error) {
	db := t.gorm.DB()
	if tx != nil {
		db = tx
	}
	var studentCount int
	if err := db.Table("workout_logs").
		Select("COUNT(DISTINCT workout_logs.user_id)").
		Joins("INNER JOIN workouts ON workout_logs.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.user_id = ?", userID).
		Take(&studentCount).Error; err != nil {
		return 0, err
	}
	return studentCount, nil
}
