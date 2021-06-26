package repository

import (
	"errors"
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

func (a *action) DeleteActionByID(actionID int64) error {
	if err := a.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//獲取與動作關聯的課表狀態
		var courseStatus int
		if err := tx.
			Table("courses").
			Select("courses.course_status").
			Joins("INNER JOIN actions ON courses.id = actions.course_id").
			Where("actions.id = ?", actionID).
			Take(&courseStatus).Error; err != nil {
				return err
		}
		//課表狀態必須是"準備中"或"被退審"，才可刪除動作
		if !(courseStatus == 1 || courseStatus == 4)  {
			return errors.New("9006-權限不足,存取遭拒")
		}
		//刪除動作
		if err := tx.Where("id = ?", actionID).
			Delete(&model.Action{}).Error; err != nil {
				return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (a *action) CheckActionExistByUID(uid int64, actionID int64) (bool, error) {
	var result int
	if err := a.gorm.DB().
		Table("actions").
		Select("1").
		Joins("INNER JOIN courses ON actions.course_id = courses.id ").
		Joins("INNER JOIN users ON courses.user_id = users.id ").
		Where("actions.id = ? AND users.id = ?", actionID, uid).
		Find(&result).Error; err != nil {
		return false, err
	}
	return result > 0, nil
}
