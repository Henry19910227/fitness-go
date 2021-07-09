package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type set struct {
	gorm tool.Gorm
}

func NewWorkoutSet(gorm tool.Gorm) WorkoutSet {
	return &set{gorm: gorm}
}

func (s *set) CreateWorkoutSetsByWorkoutID(workoutID int64, actionIDs []int64) ([]int64, error) {
	sets := make([]*model.WorkoutSet, 0)
	for _, v := range actionIDs {
		var actionID = v
		set := model.WorkoutSet{
			WorkoutID: workoutID,
			ActionID: &actionID,
			Type: 1,
			AutoNext: "N",
			Weight: 1,
			Reps: 1,
			Distance: 1,
			Duration: 1,
			Incline: 1,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		sets = append(sets, &set)
	}
	if err := s.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建訓練組
		if err := tx.Create(&sets).Error; err != nil {
			return err
		}
		//更新訓練的訓練組個數
		countQuery := tx.Table("workout_sets").
			Select("COUNT(*) AS workout_set_count").
			Where("workout_id = ?", workoutID)
		if err := tx.Table("workouts").
			Where("id = ?", workoutID).
			Update("workout_set_count", countQuery).Error; err != nil {
				return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	workoutIDs := make([]int64, 0)
	for _, set := range sets{
		workoutIDs = append(workoutIDs, set.ID)
	}
	return workoutIDs, nil
}

func (s *set) CreateRestSetByWorkoutID(workoutID int64) (int64, error) {
	set := model.WorkoutSet{
		WorkoutID: workoutID,
		Type: 2,
		AutoNext: "N",
		Duration: 30,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.gorm.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&set).Error; err != nil {
			return err
		}
		var setCount int
		if err := tx.Raw("SELECT COUNT(*) FROM workout_sets WHERE workout_id = ? FOR UPDATE", workoutID).
			Scan(&setCount).Error; err != nil {
				return err
		}
		if err := tx.
			Table("workouts").
			Where("id = ?", workoutID).
			Update("workout_set_count", setCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return set.ID, nil
}

func (s *set) FindWorkoutSetByID(setID int64) (*model.WorkoutSet, error) {
	var set model.WorkoutSet
	if err := s.gorm.DB().
		Table("workout_sets").
		Select("*").
		Where("id = ?", setID).
		Take(&set).Error; err != nil {
			return nil, err
	}
	return &set, nil
}