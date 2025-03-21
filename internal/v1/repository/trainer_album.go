package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"time"
)

type trainerAlbum struct {
	gorm tool.Gorm
}

func NewTrainerAlbum(gorm tool.Gorm) TrainerAlbum {
	return &trainerAlbum{gorm: gorm}
}

func (t *trainerAlbum) CreateAlbumPhoto(uid int64, imageNamed string) error {
	photo := model.TrainerAlbumPhoto{
		UserID:   uid,
		Photo:    imageNamed,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := t.gorm.DB().Create(&photo).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerAlbum) FindAlbumPhotoByUID(uid int64) ([]*model.TrainerAlbumPhotoEntity, error) {
	photos := make([]*model.TrainerAlbumPhotoEntity, 0)
	if err := t.gorm.DB().Table("trainer_album").
		Select("id", "photo", "create_at").
		Where("user_id = ?", uid).
		Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (t *trainerAlbum) FindAlbumPhotosByUID(uid int64, input interface{}) error {
	if err := t.gorm.DB().Model(&entity.TrainerAlbum{}).
		Where("user_id = ?", uid).
		Find(input).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerAlbum) FindAlbumPhotoByID(photoID int64, entity interface{}) error {
	if err := t.gorm.DB().
		Model(&model.TrainerAlbumPhoto{}).
		Where("id = ?", photoID).
		Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerAlbum) FindAlbumPhotosByIDs(photoIDs []int64, entity interface{}) error {
	if err := t.gorm.DB().
		Model(&model.TrainerAlbumPhoto{}).
		Where("id IN (?)", photoIDs).
		Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerAlbum) DeleteAlbumPhotoByID(photoID int64) error {
	if err := t.gorm.DB().Where("id = ?", photoID).
		Delete(&model.TrainerAlbumPhoto{}).Error; err != nil {
		return err
	}
	return nil
}
