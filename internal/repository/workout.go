package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type workout struct {
	gorm tool.Gorm
}

func NewWorkout(gorm tool.Gorm) Workout {
	return &workout{gorm: gorm}
}

func (w *workout) CreateWorkout(planID int64, name string) (int64, error) {
	workout := entity.Workout{
		PlanID:   planID,
		Name:     name,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := w.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//取得關聯課表id
		var courseID int64
		if err := tx.
			Table("plans").
			Select("course_id").
			Where("id = ?", planID).
			Take(&courseID).Error; err != nil {
			return err
		}
		//創建訓練
		if err := tx.Create(&workout).Error; err != nil {
			return err
		}
		//查詢關聯計畫當前包含的訓練總量
		var workoutCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM workouts WHERE plan_id = ? FOR UPDATE", planID).
			Scan(&workoutCount).Error; err != nil {
			return err
		}
		//查詢關聯課表當前包含的訓練總量
		var courseWorkoutCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM plans "+
				"INNER JOIN courses ON courses.id = plans.course_id "+
				"INNER JOIN workouts ON plans.id = workouts.plan_id "+
				"WHERE courses.id = ? FOR UPDATE", courseID).
			Scan(&courseWorkoutCount).Error; err != nil {
			return err
		}
		//更新計畫擁有的訓練數量
		if err := tx.
			Table("plans").
			Where("id = ?", planID).
			Update("workout_count", workoutCount).Error; err != nil {
			return err
		}
		//更新課表擁有的訓練數量
		if err := tx.
			Table("courses").
			Where("id = ?", courseID).
			Update("workout_count", courseWorkoutCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return workout.ID, nil
}

func (w *workout) FindWorkoutsByPlanID(planID int64) ([]*entity.Workout, error) {
	workouts := make([]*entity.Workout, 0)
	if err := w.gorm.DB().
		Table("workouts").
		Select("*").
		Where("plan_id = ?", planID).
		Find(&workouts).Error; err != nil {
		return nil, err
	}
	return workouts, nil
}

func (w *workout) FindWorkoutAssets(userID int64, planID int64) ([]*model.WorkoutAsset, error) {
	workouts := make([]*model.WorkoutAsset, 0)
	if err := w.gorm.DB().
		Table("workouts").
		Select("DISTINCT(workouts.id) AS id", "workouts.plan_id AS plan_id",
			"workouts.name AS name", "workouts.equipment AS equipment",
			"workouts.start_audio AS start_audio", "workouts.end_audio AS end_audio",
			"workouts.workout_set_count AS workout_set_count", "IF(log.duration > 0,1,0) AS finish",
			"workouts.create_at AS create_at", "workouts.update_at AS update_at").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
		Joins("LEFT JOIN workout_logs AS log ON workouts.id = log.workout_id AND log.user_id = ?", userID).
		Where("plan_id = ?", planID).
		Find(&workouts).Error; err != nil {
		return nil, err
	}
	return workouts, nil
}

func (w *workout) FindWorkoutByID(workoutID int64, obj interface{}) error {
	if err := w.gorm.DB().
		Model(&entity.Workout{}).
		Where("id = ?", workoutID).
		Take(obj).Error; err != nil {
		return err
	}
	return nil
}

func (w *workout) UpdateWorkoutByID(workoutID int64, param *model.UpdateWorkoutParam) error {
	var selects []interface{}
	if param.Name != nil {
		selects = append(selects, "name")
	}
	if param.Equipment != nil {
		selects = append(selects, "equipment")
	}
	if param.StartAudio != nil {
		selects = append(selects, "start_audio")
	}
	if param.EndAudio != nil {
		selects = append(selects, "end_audio")
	}
	if param == nil || len(selects) == 0 {
		return nil
	}
	//插入更新時間
	selects = append(selects, "update_at")
	var updateAt = time.Now().Format("2006-01-02 15:04:05")
	param.UpdateAt = &updateAt
	if err := w.gorm.DB().
		Table("workouts").
		Where("id = ?", workoutID).
		Select("", selects...).
		Updates(param).Error; err != nil {
		return err
	}
	return nil
}

func (w *workout) DeleteWorkoutByID(workoutID int64) error {
	if err := w.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//查詢plan id & course id
		var courseID int64
		var planID int64
		if err := tx.
			Table("workouts").
			Select("plans.id, plans.course_id").
			Joins("INNER JOIN plans ON workouts.plan_id = plans.id").
			Where("workouts.id = ?", workoutID).
			Row().
			Scan(&planID, &courseID); err != nil {
			return err
		}
		//刪除訓練
		if err := tx.
			Where("id = ?", workoutID).
			Delete(&entity.Workout{}).Error; err != nil {
			return err
		}
		//查詢關聯計畫的訓練數量
		var workoutCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM workouts WHERE plan_id = ? FOR UPDATE", planID).
			Scan(&workoutCount).Error; err != nil {
			return err
		}
		//查詢關聯課表的目前總訓練數量
		var courseWorkoutCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM courses "+
				"INNER JOIN plans ON courses.id = plans.course_id "+
				"INNER JOIN workouts ON plans.id = workouts.plan_id "+
				"WHERE courses.id = ? FOR UPDATE", courseID).
			Scan(&courseWorkoutCount).Error; err != nil {
			return err
		}
		//更新計畫擁有的訓練數量
		if err := tx.
			Table("plans").
			Where("id = ?", planID).
			Update("workout_count", workoutCount).Error; err != nil {
			return err
		}
		//更新課表擁有的訓練數量
		if err := tx.
			Table("courses").
			Where("id = ?", courseID).
			Update("workout_count", courseWorkoutCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (w *workout) FindWorkoutOwnerByID(workoutID int64) (int64, error) {
	var userID int64
	if err := w.gorm.DB().
		Table("workouts").
		Select("courses.user_id").
		Joins("INNER JOIN plans ON workouts.plan_id = plans.id ").
		Joins("INNER JOIN courses ON plans.course_id = courses.id ").
		Where("workouts.id = ?", workoutID).
		Take(&userID).Error; err != nil {
		return 0, err
	}
	return userID, nil
}

func (w *workout) FindStartAudioCountByAudioName(audioName string) (int, error) {
	var startAudioCount int
	if err := w.gorm.DB().
		Table("workouts").
		Select("COUNT(*)").
		Where("start_audio = ?", audioName).
		Take(&startAudioCount).Error; err != nil {
		return 0, err
	}
	return startAudioCount, nil
}

func (w *workout) FindEndAudioCountByAudioName(audioName string) (int, error) {
	var endAudioCount int
	if err := w.gorm.DB().
		Table("workouts").
		Select("COUNT(*)").
		Where("end_audio = ?", audioName).
		Take(&endAudioCount).Error; err != nil {
		return 0, err
	}
	return endAudioCount, nil
}
