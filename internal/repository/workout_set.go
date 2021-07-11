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

func (s *set) FindWorkoutSetByID(setID int64) (*model.WorkoutSetEntity, error) {
	row := s.gorm.DB().
		Table("workout_sets AS `set`").
		Select("`set`.id", "`set`.workout_id", "`set`.type",
			"`set`.auto_next", "`set`.start_audio", "`set`.progress_audio",
			"`set`.remark", "`set`.weight", "`set`.reps",
			"`set`.distance", "`set`.duration", "`set`.incline",
			"IFNULL(actions.id, 0)", "IFNULL(actions.name, '')", "IFNULL(actions.source, 0)",
			"IFNULL(actions.type, 0)", "IFNULL(actions.intro, '')", "IFNULL(actions.cover, '')",
			"IFNULL(actions.video, '')").
		Joins("LEFT JOIN actions ON set.action_id = actions.id").
		Where("`set`.id = ?", setID).Row()
	var set model.WorkoutSetEntity
	var action model.WorkoutSetAction
	if err := row.Scan(&set.ID, &set.WorkoutID, &set.Type,
		&set.AutoNext, &set.StartAudio, &set.ProgressAudio,
		&set.Remark, &set.Weight, &set.Reps,
		&set.Distance, &set.Duration, &set.Incline,
		&action.ID, &action.Name, &action.Source,
		&action.Type, &action.Intro, &action.Cover,
		&action.Video); err != nil {
		return nil, err
	}
	if action.ID != 0 {
		set.Action = &action
	}
	return &set, nil
}

func (s *set) FindWorkoutSetsByIDs(setIDs []int64) ([]*model.WorkoutSetEntity, error) {
	rows, err := s.gorm.DB().
		Table("workout_sets AS `set`").
		Select("`set`.id", "`set`.workout_id", "`set`.type",
			"`set`.auto_next", "`set`.start_audio", "`set`.progress_audio",
			"`set`.remark", "`set`.weight", "`set`.reps",
			"`set`.distance", "`set`.duration", "`set`.incline",
			"IFNULL(actions.id, 0)", "IFNULL(actions.name, '')", "IFNULL(actions.source, 0)",
			"IFNULL(actions.type, 0)", "IFNULL(actions.intro, '')", "IFNULL(actions.cover, '')",
			"IFNULL(actions.video, '')").
		Joins("LEFT JOIN actions ON set.action_id = actions.id").
		Where("`set`.id IN (?)", setIDs).Rows()
	if err != nil {
		return nil, err
	}
	var sets []*model.WorkoutSetEntity
	for rows.Next() {
		var set model.WorkoutSetEntity
		var action model.WorkoutSetAction
		rows.Scan(&set.ID, &set.WorkoutID, &set.Type,
			&set.AutoNext, &set.StartAudio, &set.ProgressAudio,
			&set.Remark, &set.Weight, &set.Reps,
			&set.Distance, &set.Duration, &set.Incline,
			&action.ID, &action.Name, &action.Source,
			&action.Type, &action.Intro, &action.Cover,
			&action.Video)
		if action.ID != 0 {
			set.Action = &action
		}
		sets = append(sets, &set)
	}
	return sets, nil
}