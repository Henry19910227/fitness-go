package review_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_statistic"
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

func (r *repository) Statistic(input *model.StatisticInput) (err error) {
	err = r.db.Exec("INSERT INTO review_statistics (course_id, score_total, amount, five_total, four_total, three_total, two_total, one_total, update_at) " +
		"SELECT course_id,score_total,amount,five_total,four_total,three_total,two_total,one_total,NOW() AS update_at " +
		"FROM " +
		"(SELECT ? AS course_id) AS course_id_table," +
		"(SELECT SUM(score) AS score_total, COUNT(*) AS amount FROM reviews WHERE course_id = ?) AS total_table, " +
		"(SELECT COUNT(*) AS five_total FROM reviews WHERE course_id = ? AND score = '5') AS five_total_table, " +
		"(SELECT COUNT(*) AS four_total FROM reviews WHERE course_id = ? AND score = '4') AS four_total_table, " +
		"(SELECT COUNT(*) AS three_total FROM reviews WHERE course_id = ? AND score = '3') AS three_total_table, " +
		"(SELECT COUNT(*) AS two_total FROM reviews WHERE course_id = ? AND score = '2') AS two_total_table, " +
		"(SELECT COUNT(*) AS one_total FROM reviews WHERE course_id = ? AND score = '1') AS one_total_table " +
		"ON DUPLICATE KEY UPDATE " +
		"score_total = total_table.score_total, " +
		"amount = total_table.amount, " +
		"five_total = five_total_table.five_total, " +
		"four_total = four_total_table.four_total, " +
		"three_total = three_total_table.three_total, " +
		"two_total = two_total_table.two_total, " +
		"one_total = one_total_table.one_total, " +
		"update_at = update_at",input.CourseID, input.CourseID, input.CourseID, input.CourseID, input.CourseID, input.CourseID, input.CourseID).Error
	return err
}
