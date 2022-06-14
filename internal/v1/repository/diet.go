package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type diet struct {
	gorm tool.Gorm
}

func NewDiet(gorm tool.Gorm) Diet {
	return &diet{gorm: gorm}
}

func (d *diet) CreateDiet(tx *gorm.DB, userID int64, rdaID int64, scheduleTime string) (int64, error) {
	db := d.gorm.DB()
	if tx != nil {
		db = tx
	}
	diet := entity.Diet{
		UserID:     userID,
		RdaID:      rdaID,
		ScheduleAt: scheduleTime,
		CreateAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "schedule_at"}},
		DoUpdates: clause.AssignmentColumns([]string{"update_at"}),
	}).Create(&diet).Error; err != nil {
		return 0, err
	}
	return diet.ID, nil
}

func (d *diet) SaveDiets(tx *gorm.DB, param *model.SaveDietsParam) error {
	if len(param.Diets) == 0 {
		return nil
	}
	db := d.gorm.DB()
	if tx != nil {
		db = tx
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"rda_id", "update_at"}),
	}).Create(&param.Diets).Error; err != nil {
		return err
	}
	return nil
}

func (d *diet) FindDiet(tx *gorm.DB, param *model.FindDietParam, preloads []*model.Preload) (*model.Diet, error) {
	db := d.gorm.DB()
	if tx != nil {
		db = tx
	}
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if param.ID != nil {
		query += "AND id = ? "
		params = append(params, *param.ID)
	}
	//加入 user_id 篩選條件
	if param.UserID != nil {
		query += "AND user_id = ? "
		params = append(params, *param.UserID)
	}
	//加入 schedule_at 篩選條件
	if param.ScheduleAt != nil {
		query += "AND DATE_FORMAT(schedule_at,'%Y-%m-%d') = DATE_FORMAT(?,'%Y-%m-%d')"
		params = append(params, *param.ScheduleAt)
	}
	//設置表
	db.Model(&model.Diet{})
	//關聯加載
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查找數據
	var diet model.Diet
	if err := db.Where(query, params...).Take(&diet).Error; err != nil {
		return nil, err
	}
	return &diet, nil
}

func (d *diet) FindDiets(tx *gorm.DB, param *model.FindDietsParam) ([]*model.Diet, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 user_id 篩選條件
	if param.UserID != nil {
		query += "AND user_id = ? "
		params = append(params, *param.UserID)
	}
	//加入 schedule_at 篩選條件
	if param.AfterScheduleAt != nil {
		query += "AND DATE_FORMAT(?,'%Y-%m-%d') <= DATE_FORMAT(schedule_at,'%Y-%m-%d')"
		params = append(params, *param.AfterScheduleAt)
	}
	db := d.gorm.DB()
	if tx != nil {
		db = tx
	}
	db = db.Model(&model.Diet{})
	//關聯加載
	if len(param.Preloads) > 0 {
		for _, preload := range param.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查找數據
	diets := make([]*model.Diet, 0)
	if err := db.Where(query, params...).Find(&diets).Error; err != nil {
		return nil, err
	}
	return diets, nil
}
