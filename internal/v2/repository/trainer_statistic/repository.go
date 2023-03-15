package trainer_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
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

func (r *repository) StatisticReviewScore(input *model.StatisticReviewScoreInput) (err error) {
	err = r.db.Exec("INSERT INTO trainer_statistics (user_id,review_score,update_at) "+
		"SELECT user_id,review_score,NOW() AS update_at "+
		"FROM "+
		"( "+
		"SELECT courses.user_id AS user_id,FORMAT(SUM(rs.score_total) / SUM(rs.amount),1) AS review_score "+
		"FROM review_statistics AS rs "+
		"INNER JOIN courses ON rs.course_id = courses.id "+
		"WHERE courses.user_id = ? "+
		"GROUP BY courses.user_id"+
		") AS ts "+
		"ON DUPLICATE KEY UPDATE "+
		"review_score = ts.review_score, "+
		"update_at = NOW()", input.UserID).Error
	return err
}

// StatisticStudentCount SQL
/*
INSERT INTO trainer_statistics (user_id, student_count)
SELECT courses.user_id, COUNT(DISTINCT workout_logs.user_id) AS student_count
FROM courses
INNER JOIN plans ON courses.id = plans.course_id
INNER JOIN workouts ON plans.id = workouts.plan_id
INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
WHERE courses.course_status = 3
GROUP BY courses.user_id
ON DUPLICATE KEY UPDATE student_count = VALUES(student_count), update_at = CURRENT_TIMESTAMP;
*/
func (r *repository) StatisticStudentCount() (err error) {
	err = r.db.Exec("INSERT INTO trainer_statistics (user_id, student_count) " +
		"SELECT " +
		"courses.user_id, " +
		"COUNT(DISTINCT workout_logs.user_id) AS student_count " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"WHERE courses.course_status = 3 " +
		"GROUP BY courses.user_id " +
		"ON DUPLICATE KEY UPDATE student_count = VALUES(student_count), update_at = CURRENT_TIMESTAMP").Error
	return err
}
