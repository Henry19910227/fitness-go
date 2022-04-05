package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/tool"
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
		UserID: userID,
		CourseID: courseID,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := f.gorm.DB().Create(&course).Error; err != nil {
		return err
	}
	return nil
}

