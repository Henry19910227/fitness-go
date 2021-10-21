package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type action struct {
	gorm tool.Gorm
}

func NewAction(gorm tool.Gorm) Action {
	return &action{gorm: gorm}
}

func (a *action) CreateAction(courseID int64, param *model.CreateActionParam) (int64, error) {
	action := model.Action{
		CourseID: courseID,
		Name: param.Name,
		Source: 2,
		Type: param.Type,
		Category: param.Category,
		Body: param.Body,
		Equipment: param.Equipment,
		Intro: param.Intro,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if param.Cover != nil {
		action.Cover = *param.Cover
	}
	if param.Video != nil {
		action.Video = *param.Video
	}
	if err := a.gorm.DB().Create(&action).Error; err != nil {
		return 0, err
	}
	return action.ID, nil
}

func (a *action) FindActionByID(actionID int64, entity interface{}) error {
	if err := a.gorm.DB().
		Model(&model.Action{}).
		Where("id = ?", actionID).
		Take(entity).Error; err != nil{
			return err
	}
	return nil
}

func (a *action) UpdateActionByID(actionID int64, param *model.UpdateActionParam) error {
	var selects []interface{}
	if param.Name != nil { selects = append(selects, "name") }
	if param.Category != nil { selects = append(selects, "category") }
	if param.Body != nil { selects = append(selects, "body") }
	if param.Equipment != nil { selects = append(selects, "equipment") }
	if param.Intro != nil { selects = append(selects, "intro") }
	if param.Cover != nil { selects = append(selects, "cover") }
	if param.Video != nil { selects = append(selects, "video") }
	if param == nil || len(selects) == 0 { return nil }
	//插入更新時間
	selects = append(selects, "update_at")
	var updateAt = time.Now().Format("2006-01-02 15:04:05")
	param.UpdateAt = &updateAt
	if err := a.gorm.DB().
		Table("actions").
		Where("id = ?", actionID).
		Select("", selects...).
		Updates(param).Error; err != nil {
		return err
	}
	return nil
}

func (a *action) FindActionsByParam(courseID int64, param *model.FindActionsParam, entity interface{}) error {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 course_id 篩選條件
	query += "AND (course_id = ? OR course_id IS NULL) "
	params = append(params, courseID)
	//加入 source 篩選條件
	if len(*param.SourceOpt) > 0 {
		query += "AND source IN ? "
		params = append(params, *param.SourceOpt)
	}
	//加入 body 篩選條件
	if len(*param.BodyOpt) > 0 {
		query += "AND body IN ? "
		params = append(params, *param.BodyOpt)
	}
	//加入 category 篩選條件
	if len(*param.CategoryOpt) > 0 {
		query += "AND category IN ? "
		params = append(params, *param.CategoryOpt)
	}
	//加入 equipment 篩選條件
	if len(*param.EquipmentOpt) > 0 {
		query += "AND equipment IN ? "
		params = append(params, *param.EquipmentOpt)
	}
	//加入 name 篩選條件
	if param.Name != nil {
		query += "AND name LIKE ? "
		params = append(params, "%" + *param.Name + "%")
	}
	if err := a.gorm.DB().
		Model(&model.Action{}).
		Where(query, params...).
		Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (a *action) DeleteActionByID(actionID int64) error {
	if err := a.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//查找刪除此action後受影響之workout id (去除重複)
		workoutIDs := make([]int, 0)
		if err := tx.Table("workout_sets").
			Select("workout_id").
			Where("action_id = ?", actionID).
			Group("workout_id").
			Find(&workoutIDs).Error; err != nil {
			return err
		}
		//刪除action
		if err := tx.Where("id = ?", actionID).
			Delete(&model.Action{}).Error; err != nil {
			return err
		}
		//更新相關聯workout內訓練組數量
		for _, workoutID := range workoutIDs {
			if err := tx.Table("workouts").Where("id = ?", workoutID).Update("workout_set_count",
				tx.Table("workout_sets AS sets").Select("count(*)").Where("sets.workout_id = ?", workoutID),
			).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}