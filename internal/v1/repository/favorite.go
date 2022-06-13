package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"time"
)

type favorite struct {
	gorm tool.Gorm
}

func NewFavorite(gorm tool.Gorm) Favorite {
	return &favorite{gorm: gorm}
}

func (f *favorite) CreateFavoriteCourse(userID int64, courseID int64) error {
	course := entity.FavoriteCourse{
		UserID:   userID,
		CourseID: courseID,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := f.gorm.DB().Create(&course).Error; err != nil {
		return err
	}
	return nil
}

func (f *favorite) CreateFavoriteTrainer(userID int64, trainerID int64) error {
	trainer := entity.FavoriteTrainer{
		UserID:    userID,
		TrainerID: trainerID,
		CreateAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := f.gorm.DB().Create(&trainer).Error; err != nil {
		return err
	}
	return nil
}

func (f *favorite) CreateFavoriteAction(userID int64, actionID int64) error {
	action := entity.FavoriteAction{
		UserID:   userID,
		ActionID: actionID,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := f.gorm.DB().Create(&action).Error; err != nil {
		return err
	}
	return nil
}

func (f *favorite) FindFavoriteCourse(userID int64, courseID int64) (*model.FavoriteCourse, error) {
	var course model.FavoriteCourse
	if err := f.gorm.DB().
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Find(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (f *favorite) FindFavoriteTrainer(userID int64, trainerID int64) (*model.FavoriteTrainer, error) {
	var trainer model.FavoriteTrainer
	if err := f.gorm.DB().
		Where("user_id = ? AND trainer_id = ?", userID, trainerID).
		Find(&trainer).Error; err != nil {
		return nil, err
	}
	return &trainer, nil
}

func (f *favorite) DeleteFavoriteCourse(userID int64, courseID int64) error {
	if err := f.gorm.DB().
		Delete(&entity.FavoriteCourse{}, "user_id = ? AND course_id = ?", userID, courseID).Error; err != nil {
		return err
	}
	return nil
}

func (f *favorite) DeleteFavoriteTrainer(userID int64, trainerID int64) error {
	if err := f.gorm.DB().
		Delete(&entity.FavoriteTrainer{}, "user_id = ? AND trainer_id = ?", userID, trainerID).Error; err != nil {
		return err
	}
	return nil
}

func (f *favorite) DeleteFavoriteAction(userID int64, actionID int64) error {
	if err := f.gorm.DB().
		Delete(&entity.FavoriteAction{}, "user_id = ? AND action_id = ?", userID, actionID).Error; err != nil {
		return err
	}
	return nil
}
