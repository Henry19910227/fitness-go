package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type actionPR struct {
	gorm tool.Gorm
}

func NewActionPR(gorm tool.Gorm) ActionPR {
	return &actionPR{gorm: gorm}
}

func (a *actionPR) FindActionPR(tx *gorm.DB, userID int64, actionID int64) (*model.ActionPR, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	var actionPR model.ActionPR
	if err := db.
		Table("action_personal_records").
		Where("user_id = ? AND action_id = ?", userID, actionID).
		Take(&actionPR).Error; err != nil {
		return nil, err
	}
	return &actionPR, nil
}

func (a *actionPR) FindActionPRs(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.ActionPR, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	actionPRs := make([]*model.ActionPR, 0)
	if err := db.
		Table("action_personal_records").
		Where("user_id = ? AND action_id IN (?)", userID, actionIDs).
		Find(&actionPRs).Error; err != nil {
		return nil, err
	}
	return actionPRs, nil
}

func (a *actionPR) FindActionBestPRs(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.ActionBestPR, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	actionPRs := make([]*model.ActionBestPR, 0)
	if err := db.
		Table("actions").
		Select("IFNULL(actions.id,0) AS action_id",
			"IFNULL(max_rm_records.rm,0) AS max_rm",
			"IFNULL(max_reps_records.reps,0) AS max_reps",
			"IFNULL(max_speed_records.speed,0) AS max_speed",
			"IFNULL(max_weight_records.weight,0) AS max_weight",
			"IFNULL(min_duration_records.duration,0)AS min_duration").
		Joins("LEFT JOIN max_rm_records ON actions.id = max_rm_records.action_id AND max_rm_records.user_id = ?", userID).
		Joins("LEFT JOIN max_reps_records ON actions.id = max_reps_records.action_id AND max_reps_records.user_id = ?", userID).
		Joins("LEFT JOIN max_speed_records ON actions.id = max_speed_records.action_id AND max_speed_records.user_id = ?", userID).
		Joins("LEFT JOIN max_weight_records ON actions.id = max_weight_records.action_id AND max_weight_records.user_id = ?", userID).
		Joins("LEFT JOIN min_duration_records ON actions.id = min_duration_records.action_id AND min_duration_records.user_id = ?", userID).
		Where("actions.id IN (?)", actionIDs).
		Find(&actionPRs).Error; err != nil {
		return nil, err
	}
	return actionPRs, nil
}

func (a *actionPR) SaveActionPRs(tx *gorm.DB, userID int64, params []*model.CreateActionPRParam) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	actionPRs := make([]*entity.ActionPR, 0)
	for _, param := range params {
		actionPR := entity.ActionPR{
			UserID:   userID,
			ActionID: param.ActionID,
			Weight:   param.Weight,
			Reps:     param.Reps,
			Distance: param.Distance,
			Duration: param.Duration,
			Incline:  param.Incline,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		actionPRs = append(actionPRs, &actionPR)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"weight", "reps", "distance", "duration", "incline", "update_at"}),
	}).Create(&actionPRs).Error; err != nil {
		return err
	}
	return nil
}

func (a *actionPR) SaveMaxRMRecords(tx *gorm.DB, params []*model.SaveMaxRmRecord) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	records := make([]*entity.MaxRmRecord, 0)
	for _, param := range params {
		record := entity.MaxRmRecord{
			UserID:   param.UserID,
			ActionID: param.ActionID,
			RM:       param.RM,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		records = append(records, &record)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"rm", "update_at"}),
	}).Create(&records).Error; err != nil {
		return err
	}
	return nil
}

func (a *actionPR) SaveMaxRepsRecords(tx *gorm.DB, params []*model.SaveMaxRepsRecord) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	records := make([]*entity.MaxRepsRecord, 0)
	for _, param := range params {
		record := entity.MaxRepsRecord{
			UserID:   param.UserID,
			ActionID: param.ActionID,
			Reps:     param.Reps,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		records = append(records, &record)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"reps", "update_at"}),
	}).Create(&records).Error; err != nil {
		return err
	}
	return nil
}

func (a *actionPR) SaveMaxWeightRecords(tx *gorm.DB, params []*model.SaveMaxWeightRecord) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	records := make([]*entity.MaxWeightRecord, 0)
	for _, param := range params {
		record := entity.MaxWeightRecord{
			UserID:   param.UserID,
			ActionID: param.ActionID,
			Weight:   param.Weight,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		records = append(records, &record)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"weight", "update_at"}),
	}).Create(&records).Error; err != nil {
		return err
	}
	return nil
}

func (a *actionPR) SaveMinDurationRecords(tx *gorm.DB, params []*model.SaveMinDurationRecord) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	records := make([]*entity.MinDurationRecord, 0)
	for _, param := range params {
		record := entity.MinDurationRecord{
			UserID:   param.UserID,
			ActionID: param.ActionID,
			Duration: param.Duration,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		records = append(records, &record)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"duration", "update_at"}),
	}).Create(&records).Error; err != nil {
		return err
	}
	return nil
}

func (a *actionPR) SaveMaxSpeedRecords(tx *gorm.DB, params []*model.SaveMaxSpeedRecord) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	records := make([]*entity.MaxSpeedRecord, 0)
	for _, param := range params {
		record := entity.MaxSpeedRecord{
			UserID:   param.UserID,
			ActionID: param.ActionID,
			Speed:    param.Speed,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		records = append(records, &record)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"speed", "update_at"}),
	}).Create(&records).Error; err != nil {
		return err
	}
	return nil
}

func (a *actionPR) CalculateMaxRM(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.MaxRmRecord, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	var records []*model.MaxRmRecord
	if err := db.
		Table("actions").
		Select("MAX(workout_logs.user_id) AS user_id",
			"MAX(actions.id) AS action_id",
			"TRUNCATE(MAX(workout_set_logs.weight * (1 + 0.333 * workout_set_logs.reps)),1) AS rm").
		Joins("INNER JOIN workout_sets ON workout_sets.action_id = actions.id").
		Joins("INNER JOIN workout_set_logs ON workout_set_logs.workout_set_id = workout_sets.id").
		Joins("INNER JOIN workout_logs ON workout_logs.id = workout_set_logs.workout_log_id").
		Where("actions.id IN (?) AND workout_logs.user_id = ?", actionIDs, userID).
		Group("actions.id").
		Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (a *actionPR) CalculateMaxReps(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.MaxRepsRecord, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	var records []*model.MaxRepsRecord
	if err := db.
		Table("actions").
		Select("MAX(workout_logs.user_id) AS user_id",
			"MAX(actions.id) AS action_id",
			"MAX(workout_set_logs.reps) AS reps").
		Joins("INNER JOIN workout_sets ON workout_sets.action_id = actions.id").
		Joins("INNER JOIN workout_set_logs ON workout_set_logs.workout_set_id = workout_sets.id").
		Joins("INNER JOIN workout_logs ON workout_logs.id = workout_set_logs.workout_log_id").
		Where("actions.id IN (?) AND workout_logs.user_id = ?", actionIDs, userID).
		Group("actions.id").
		Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (a *actionPR) CalculateMaxWeight(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.MaxWeightRecord, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	var records []*model.MaxWeightRecord
	if err := db.
		Table("actions").
		Select("MAX(workout_logs.user_id) AS user_id",
			"MAX(actions.id) AS action_id",
			"MAX(workout_set_logs.weight) AS weight").
		Joins("INNER JOIN workout_sets ON workout_sets.action_id = actions.id").
		Joins("INNER JOIN workout_set_logs ON workout_set_logs.workout_set_id = workout_sets.id").
		Joins("INNER JOIN workout_logs ON workout_logs.id = workout_set_logs.workout_log_id").
		Where("actions.id IN (?) AND workout_logs.user_id = ?", actionIDs, userID).
		Group("actions.id").
		Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (a *actionPR) CalculateMinDuration(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.MinDurationRecord, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	var records []*model.MinDurationRecord
	if err := db.
		Table("actions").
		Select("MAX(workout_logs.user_id) AS user_id",
			"MAX(actions.id) AS action_id",
			"MIN(workout_set_logs.duration) AS duration").
		Joins("INNER JOIN workout_sets ON workout_sets.action_id = actions.id").
		Joins("INNER JOIN workout_set_logs ON workout_set_logs.workout_set_id = workout_sets.id").
		Joins("INNER JOIN workout_logs ON workout_logs.id = workout_set_logs.workout_log_id").
		Where("actions.id IN (?) AND workout_logs.user_id = ?", actionIDs, userID).
		Group("actions.id").
		Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (a *actionPR) CalculateMaxSpeed(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.MaxSpeedRecord, error) {
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	var records []*model.MaxSpeedRecord
	if err := db.
		Table("actions").
		Select("MAX(workout_logs.user_id) AS user_id",
			"MAX(actions.id) AS action_id",
			"TRUNCATE(MAX(workout_set_logs.distance * 1000 / workout_set_logs.duration * 3600 / 1000),1) AS speed").
		Joins("INNER JOIN workout_sets ON workout_sets.action_id = actions.id").
		Joins("INNER JOIN workout_set_logs ON workout_set_logs.workout_set_id = workout_sets.id").
		Joins("INNER JOIN workout_logs ON workout_logs.id = workout_set_logs.workout_log_id").
		Where("actions.id IN (?) AND workout_logs.user_id = ?", actionIDs, userID).
		Group("actions.id").
		Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
