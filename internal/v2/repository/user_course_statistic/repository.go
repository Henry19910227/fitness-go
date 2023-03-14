package user_course_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_statistic"
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
/*
INSERT INTO user_course_statistics (user_id, course_id, finish_workout_count, total_finish_workout_count, duration)
SELECT
10001 AS user_id,
1 AS course_id,
COUNT(DISTINCT workout_logs.workout_id) AS finish_workout_count,
COUNT(workout_logs.workout_id) AS total_finish_workout_count,
IFNULL(SUM(workout_logs.duration),0) AS duration
FROM courses
INNER JOIN plans ON courses.id = plans.course_id
INNER JOIN workouts ON plans.id = workouts.plan_id
INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
WHERE courses.id = 1 AND workout_logs.user_id = 10001
ON DUPLICATE KEY UPDATE
finish_workout_count = VALUES(finish_workout_count),
total_finish_workout_count = VALUES(total_finish_workout_count),
duration = VALUES(duration), update_at = CURRENT_TIMESTAMP;
**/
func (r *repository) Statistic(input *model.Statistic) (err error) {
	err = r.db.Exec("INSERT INTO user_course_statistics (user_id, course_id, finish_workout_count, total_finish_workout_count, duration) "+
		"SELECT "+
		"? AS user_id, "+
		"? AS course_id, "+
		"COUNT(DISTINCT workout_logs.workout_id) AS finish_workout_count, "+
		"COUNT(workout_logs.workout_id) AS total_finish_workout_count, "+
		"IFNULL(SUM(workout_logs.duration),0) AS duration "+
		"FROM courses "+
		"INNER JOIN plans ON courses.id = plans.course_id "+
		"INNER JOIN workouts ON plans.id = workouts.plan_id "+
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id "+
		"WHERE courses.id = ? AND workout_logs.user_id = ? "+
		"ON DUPLICATE KEY UPDATE "+
		"finish_workout_count = VALUES(finish_workout_count), "+
		"total_finish_workout_count = VALUES(total_finish_workout_count), "+
		"duration = VALUES(duration), "+
		"update_at = CURRENT_TIMESTAMP",
		*input.UserID, *input.CourseID, *input.CourseID, *input.UserID).Error
	return err
}
