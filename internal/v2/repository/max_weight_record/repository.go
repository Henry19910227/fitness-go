package max_weight_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/max_weight_record"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
}

func (r *repository) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"weight", "update_at"}),
	}).Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item.ID, err
}

