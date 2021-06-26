package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
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
