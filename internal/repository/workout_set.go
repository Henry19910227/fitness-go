package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
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
	sets := make([]*entity.WorkoutSet, 0)
	for _, v := range actionIDs {
		var actionID = v
		set := entity.WorkoutSet{
			WorkoutID: workoutID,
			ActionID: &actionID,
			Type: 1,
			AutoNext: "N",
			Weight: 10,
			Reps: 10,
			Distance: 1,
			Duration: 60,
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
			Where("workout_id = ? AND type = ?", workoutID, 1)
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

func (s *set) CreateWorkoutSetsByWorkoutIDAndSets(workoutID int64, sets []*entity.WorkoutSet) ([]int64, error) {
	if len(sets) == 0 {
		return []int64{}, nil
	}
	if err := s.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建訓練組
		if err := tx.Create(&sets).Error; err != nil {
			return err
		}
		//更新訓練的訓練組個數
		countQuery := tx.Table("workout_sets").
			Select("COUNT(*) AS workout_set_count").
			Where("workout_id = ? AND type = ?", workoutID, 1)
		if err := tx.Table("workouts").
			Where("id = ?", workoutID).
			Update("workout_set_count", countQuery).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	setIDs := make([]int64, 0)
	for _, set := range sets{
		setIDs = append(setIDs, set.ID)
	}
	return setIDs, nil
}

func (s *set) CreateRestSetByWorkoutID(workoutID int64) (int64, error) {
	set := entity.WorkoutSet{
		WorkoutID: workoutID,
		Type: 2,
		AutoNext: "N",
		Duration: 30,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.gorm.DB().Create(&set).Error; err != nil {
		return 0, err
	}
	//更新訓練的訓練組個數
	countQuery := s.gorm.DB().Table("workout_sets").
		Select("COUNT(*) AS workout_set_count").
		Where("workout_id = ? AND type = ?", workoutID, 1)
	if err := s.gorm.DB().Table("workouts").
		Where("id = ?", workoutID).
		Update("workout_set_count", countQuery).Error; err != nil {
		return 0, err
	}
	return set.ID, nil
}

func (s *set) FindWorkoutSetByID(setID int64) (*model.WorkoutSet, error) {
	row := s.gorm.DB().
		Table("workout_sets AS `set`").
		Select("`set`.id", "`set`.workout_id", "`set`.type",
			"`set`.auto_next", "`set`.start_audio", "`set`.progress_audio",
			"`set`.remark", "`set`.weight", "`set`.reps",
			"`set`.distance", "`set`.duration", "`set`.incline",
			"IFNULL(actions.id, 0)", "IFNULL(actions.name, '')", "IFNULL(actions.source, 0)",
			"IFNULL(actions.type, 0)", "IFNULL(actions.category, 0)", "IFNULL(actions.body, 0)",
			"IFNULL(actions.equipment, 0)", "IFNULL(actions.intro, '')", "IFNULL(actions.cover, '')",
			"IFNULL(actions.video, '')").
		Joins("LEFT JOIN actions ON set.action_id = actions.id").
		Where("`set`.id = ?", setID).Row()
	var set model.WorkoutSet
	var action model.Action
	if err := row.Scan(&set.ID, &set.WorkoutID, &set.Type,
		&set.AutoNext, &set.StartAudio, &set.ProgressAudio,
		&set.Remark, &set.Weight, &set.Reps,
		&set.Distance, &set.Duration, &set.Incline,
		&action.ID, &action.Name, &action.Source, &action.Type, &action.Category, &action.Equipment,
		&action.Body, &action.Intro, &action.Cover, &action.Video); err != nil {
		return nil, err
	}
	if action.ID != 0 {
		set.Action = &action
	}
	return &set, nil
}

func (s *set) FindWorkoutSetsByIDs(setIDs []int64) ([]*model.WorkoutSet, error) {
	rows, err := s.gorm.DB().
		Table("workout_sets AS `set`").
		Select("`set`.id", "`set`.workout_id", "`set`.type",
			"`set`.auto_next", "`set`.start_audio", "`set`.progress_audio",
			"`set`.remark", "`set`.weight", "`set`.reps",
			"`set`.distance", "`set`.duration", "`set`.incline",
		    "IFNULL(actions.id, 0)", "IFNULL(actions.name, '')", "IFNULL(actions.source, 0)",
		    "IFNULL(actions.type, 0)", "IFNULL(actions.category, 0)", "IFNULL(actions.body, 0)",
		    "IFNULL(actions.equipment, 0)", "IFNULL(actions.intro, '')", "IFNULL(actions.cover, '')",
		    "IFNULL(actions.video, '')").
		Joins("LEFT JOIN actions ON set.action_id = actions.id").
		Where("`set`.id IN (?)", setIDs).Rows()
	if err != nil {
		return nil, err
	}
	var sets []*model.WorkoutSet
	for rows.Next() {
		var set model.WorkoutSet
		var action model.Action
		if err := rows.Scan(&set.ID, &set.WorkoutID, &set.Type,
			&set.AutoNext, &set.StartAudio, &set.ProgressAudio,
			&set.Remark, &set.Weight, &set.Reps,
			&set.Distance, &set.Duration, &set.Incline,
			&action.ID, &action.Name, &action.Source, &action.Type, &action.Category, &action.Equipment,
			&action.Body, &action.Intro, &action.Cover, &action.Video); err != nil {
			return nil, err
		}
		if action.ID != 0 {
			set.Action = &action
		}
		sets = append(sets, &set)
	}
	return sets, nil
}

func (s *set) FindWorkoutSetsByWorkoutID(workoutID int64) ([]*model.WorkoutSet, error) {
	rows, err := s.gorm.DB().
		Table("workout_sets AS `set`").
		Select("`set`.id", "`set`.workout_id", "`set`.type",
			"`set`.auto_next", "`set`.start_audio", "`set`.progress_audio",
			"`set`.remark", "`set`.weight", "`set`.reps",
			"`set`.distance", "`set`.duration", "`set`.incline",
		    "IFNULL(actions.id, 0)", "IFNULL(actions.name, '')", "IFNULL(actions.source, 0)",
		    "IFNULL(actions.type, 0)", "IFNULL(actions.category, 0)", "IFNULL(actions.body, 0)",
		    "IFNULL(actions.equipment, 0)", "IFNULL(actions.intro, '')", "IFNULL(actions.cover, '')",
		    "IFNULL(actions.video, '')").
		Joins("LEFT JOIN actions ON `set`.action_id = actions.id").
		Joins("LEFT JOIN workout_set_orders AS orders ON orders.workout_set_id = `set`.id").
		Where("`set`.workout_id = ?", workoutID).
		// ORDER BY orders.seq IS NULL 表示 seq 不為空則為0，空則為1，使用ASC排序之後，1的會被排後面，達到將null值的資料放到最後的需求
		Order("orders.seq IS NULL ASC, orders.seq ASC, `set`.create_at ASC").
		Rows()
	if err != nil {
		return nil, err
	}
	var sets []*model.WorkoutSet
	for rows.Next() {
		var set model.WorkoutSet
		var action model.Action
		if err := rows.Scan(&set.ID, &set.WorkoutID, &set.Type,
			&set.AutoNext, &set.StartAudio, &set.ProgressAudio,
			&set.Remark, &set.Weight, &set.Reps,
			&set.Distance, &set.Duration, &set.Incline,
			&action.ID, &action.Name, &action.Source, &action.Type, &action.Category, &action.Equipment,
			&action.Body, &action.Intro, &action.Cover, &action.Video); err != nil {
			return nil, err
		}
		if action.ID != 0 {
			set.Action = &action
		}
		sets = append(sets, &set)
	}
	return sets, nil
}

func (s *set) FindWorkoutSetsByCourseID(courseID int64) ([]*model.WorkoutSet, error) {
	var sets []*model.WorkoutSet
	if err := s.gorm.DB().
		Table("workout_sets AS sets").
		Joins("INNER JOIN workouts ON sets.workout_id = workouts.id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("courses.id = ? AND sets.type = ?", courseID, 1).
		Preload("Action").
		Find(&sets).Error; err != nil {
			return nil, err
	}
	return sets, nil
}

func (s *set) FindStartAudioCountByAudioName(audioName string) (int, error) {
	var startAudioCount int
	if err := s.gorm.DB().
		Table("workout_sets").
		Select("COUNT(*)").
		Where("start_audio = ?", audioName).
		Take(&startAudioCount).Error; err != nil {
			return 0, err
	}
	return startAudioCount, nil
}

func (s *set) FindProgressAudioCountByAudioName(audioName string) (int, error) {
	var progressAudioCount int
	if err := s.gorm.DB().
		Table("workout_sets").
		Select("COUNT(*)").
		Where("progress_audio = ?", audioName).
		Take(&progressAudioCount).Error; err != nil {
		return 0, err
	}
	return progressAudioCount, nil
}

func (s *set) UpdateWorkoutSetByID(setID int64, param *model.UpdateWorkoutSetParam) error {
	var selects []interface{}
	if param.AutoNext != nil { selects = append(selects, "auto_next") }
	if param.StartAudio != nil { selects = append(selects, "start_audio") }
	if param.ProgressAudio != nil { selects = append(selects, "progress_audio") }
	if param.Remark != nil { selects = append(selects, "remark") }
	if param.Weight != nil { selects = append(selects, "weight") }
	if param.Reps != nil { selects = append(selects, "reps") }
	if param.Distance != nil { selects = append(selects, "distance") }
	if param.Duration != nil { selects = append(selects, "duration") }
	if param.Incline != nil { selects = append(selects, "incline") }
	//插入更新時間
	if param != nil {
		selects = append(selects, "update_at")
		var updateAt = time.Now().Format("2006-01-02 15:04:05")
		param.UpdateAt = &updateAt
	}
	if err := s.gorm.DB().
		Table("workout_sets").
		Where("id = ?", setID).
		Select("", selects...).
		Updates(param).Error; err != nil {
		return err
	}
	return nil
}

func (s *set) DeleteWorkoutSetByID(setID int64) error {
	if err := s.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//查詢workout id
		var workoutID int64
		if err := tx.
			Table("workout_sets").
			Select("workout_id").
			Where("id = ?", setID).
			Row().
			Scan(&workoutID); err != nil {
			return err
		}
		//刪除訓練組
		if err := tx.
			Where("id = ?", setID).
			Delete(&entity.WorkoutSet{}).Error; err != nil {
			return err
		}
		//查詢訓練數量
		var workoutSetCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM workout_sets WHERE workout_id = ? AND type = ? FOR UPDATE", workoutID, 1).
			Scan(&workoutSetCount).Error; err != nil {
			return err
		}
		//更新訓練擁有的訓練組數量
		if err := tx.
			Table("workouts").
			Where("id = ?", workoutID).
			Update("workout_set_count", workoutSetCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *set) UpdateWorkoutSetOrdersByWorkoutID(workoutID int64, params []*model.WorkoutSetOrder) error {
	if err := s.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//刪除舊有排序
		if err := tx.
			Where("workout_id = ?", workoutID).
			Delete(&model.WorkoutSetOrder{}).Error; err != nil {
			return err
		}
		//添加新的排序
		if err := tx.Create(params).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}