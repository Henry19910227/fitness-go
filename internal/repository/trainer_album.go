package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
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
		UserID: uid,
		Photo: imageNamed,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := t.gorm.DB().Create(&photo).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerAlbum) FindAlbumPhotoByUID(uid int64) ([]*model.TrainerAlbumPhotoEntity, error) {
	photos := make([]*model.TrainerAlbumPhotoEntity, 0)
	if err := t.gorm.DB().Table("trainer_albums").
		Select("id", "photo", "create_at").
		Where("user_id = ?", uid).
		Find(&photos).Error; err != nil {
			return nil, err
	}
	return photos, nil
}

func (t *trainerAlbum) FindAlbumPhotosByUID(uid int64, entity interface{}) error {
	if err := t.gorm.DB().Model(&model.TrainerAlbumPhoto{}).
		Where("user_id = ?", uid).
		Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainerAlbum) FindAlbumPhotoByID(photoID int64, entity interface{}) error {
	if err := t.gorm.DB().
		Model(&model.TrainerAlbumPhoto{}).
		Where("id = ?", photoID).
		Take(entity).Error; err != nil {
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
