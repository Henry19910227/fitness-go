package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"time"
)

type userCourseAsset struct {
	gorm tool.Gorm
}

func NewUserCourseAsset(gorm tool.Gorm) UserCourseAsset {
	return &userCourseAsset{gorm: gorm}
}

func (p *userCourseAsset) CreateUserCourseAsset(tx *gorm.DB, param *model.CreateUserCourseAssetParam) (int64, error) {
	db := p.gorm.DB()
	if tx != nil {
		db = tx
	}
	asset := entity.UserCourseAsset{
		UserID:    param.UserID,
		CourseID:  param.CourseID,
		Available: 1,
		CreateAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&asset).Error; err != nil {
		return 0, err
	}
	return asset.ID, nil
}

func (p *userCourseAsset) FindUserCourseAsset(param *model.FindUserCourseAssetParam) (*model.UserCourseAsset, error) {
	if param == nil {
		return nil, nil
	}
	var asset model.UserCourseAsset
	if err := p.gorm.DB().
		Table("user_course_assets").
		Where("user_id = ? AND course_id = ?", param.UserID, param.CourseID).
		Take(&asset).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}
