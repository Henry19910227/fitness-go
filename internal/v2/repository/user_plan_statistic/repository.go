package user_plan_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_plan_statistic"
	"gorm.io/gorm"
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

// Statistic SQL
/**
INSERT INTO user_plan_statistics (user_id, plan_id, finish_workout_count, duration)
SELECT
10001 AS user_id,
1999 AS plan_id,
COUNT(DISTINCT workout_logs.workout_id) AS finish_workout_count,
IFNULL(SUM(workout_logs.duration),0) AS duration
FROM plans
INNER JOIN workouts ON plans.id = workouts.plan_id
INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
WHERE plans.id = 1999 AND workout_logs.user_id = 10001
ON DUPLICATE KEY UPDATE finish_workout_count = VALUES(finish_workout_count), duration = VALUES(duration), update_at = CURRENT_TIMESTAMP;
*/
func (r *repository) Statistic(input *model.Statistic) (err error) {
	err = r.db.Exec("INSERT INTO user_plan_statistics (user_id, plan_id, finish_workout_count, duration) "+
		"SELECT "+
		"? AS user_id, "+
		"? AS plan_id, "+
		"COUNT(DISTINCT workout_logs.workout_id) AS finish_workout_count, "+
		"IFNULL(SUM(workout_logs.duration),0) AS duration "+
		"FROM plans "+
		"INNER JOIN workouts ON plans.id = workouts.plan_id "+
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id "+
		"WHERE plans.id = ? AND workout_logs.user_id = ? "+
		"ON DUPLICATE KEY UPDATE "+
		"finish_workout_count = VALUES(finish_workout_count), "+
		"duration = VALUES(duration), update_at = CURRENT_TIMESTAMP", *input.UserID, *input.PlanID, *input.PlanID, *input.UserID).Error
	return err
}
