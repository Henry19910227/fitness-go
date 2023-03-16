package user_subscribe_monthly_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic"
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
INSERT INTO user_subscribe_monthly_statistics (year, month, total, male, female, age_13_17, age_18_24, age_25_34, age_35_44, age_45_54, age_55_64, age_65_up)
SELECT
  2022 AS year,
  11 AS month,
  COUNT(DISTINCT users.id),
  COUNT(DISTINCT CASE WHEN users.sex = 'm' THEN users.id END),
  COUNT(DISTINCT CASE WHEN users.sex = 'f' THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) BETWEEN 13 AND 17 THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) BETWEEN 18 AND 24 THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) BETWEEN 25 AND 34 THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) BETWEEN 35 AND 44 THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) BETWEEN 45 AND 54 THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) BETWEEN 55 AND 64 THEN users.id END),
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 11, '-01'))) >= 65 THEN users.id END)
FROM users
INNER JOIN orders ON users.id = orders.user_id
WHERE YEAR(orders.create_at) = 2022 AND MONTH(orders.create_at) = 11
AND orders.order_type = 2 AND orders.order_status = 2
ON DUPLICATE KEY UPDATE
  total = VALUES(total),
  male = VALUES(male),
  female = VALUES(female),
  age_13_17 = VALUES(age_13_17),
  age_18_24 = VALUES(age_18_24),
  age_25_34 = VALUES(age_25_34),
  age_35_44 = VALUES(age_35_44),
  age_45_54 = VALUES(age_45_54),
  age_55_64 = VALUES(age_55_64),
  age_65_up = VALUES(age_65_up),
  update_at = CURRENT_TIMESTAMP;
*/
func (r *repository) Statistic(input *model.StatisticInput) (err error) {
	err = r.db.Exec("INSERT INTO user_subscribe_monthly_statistics (year, month, total, male, female, age_13_17, age_18_24, age_25_34, age_35_44, age_45_54, age_55_64, age_65_up) "+
		"SELECT "+
		"? AS year, "+
		"? AS month, "+
		"COUNT(DISTINCT users.id), "+
		"COUNT(DISTINCT CASE WHEN users.sex = 'm' THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN users.sex = 'f' THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 13 AND 17 THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 18 AND 24 THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 25 AND 34 THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 35 AND 44 THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 45 AND 54 THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 55 AND 64 THEN users.id END), "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) >= 65 THEN users.id END) "+
		"FROM users "+
		"INNER JOIN orders ON users.id = orders.user_id "+
		"WHERE YEAR(orders.create_at) = ? AND MONTH(orders.create_at) = ? "+
		"AND orders.order_type = 2 AND orders.order_status = 2 "+
		"ON DUPLICATE KEY UPDATE "+
		"total = VALUES(total), "+
		"male = VALUES(male), "+
		"female = VALUES(female), "+
		"age_13_17 = VALUES(age_13_17), "+
		"age_18_24 = VALUES(age_18_24), "+
		"age_25_34 = VALUES(age_25_34), "+
		"age_35_44 = VALUES(age_35_44), "+
		"age_45_54 = VALUES(age_45_54), "+
		"age_55_64 = VALUES(age_55_64), "+
		"age_65_up = VALUES(age_65_up), "+
		"update_at = CURRENT_TIMESTAMP",
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month,
		input.Year, input.Month).Error
	return err
}
