package user_promote_monthly_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_promote_monthly_statistic"
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

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 year 篩選條件
	if input.Year != nil {
		db = db.Where("year = ?", *input.Year)
	}
	//加入 month 篩選條件
	if input.Month != nil {
		db = db.Where("month = ?", *input.Month)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	//加入 year 篩選條件
	if input.Year != nil {
		db = db.Where("user_promote_monthly_statistics.year = ?", *input.Year)
	}
	//加入 month 篩選條件
	if input.Month != nil {
		db = db.Where("user_promote_monthly_statistics.month = ?", *input.Month)
	}
	// Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("user_promote_monthly_statistics.*")
	// Paging
	if input.Page != nil && input.Size != nil {
		db = db.Offset((*input.Page - 1) * *input.Size).Limit(*input.Size)
	} else if input.Page != nil {
		db = db.Offset(0)
	} else if input.Size != nil {
		db = db.Limit(*input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("user_promote_monthly_statistics.%s %s", input.OrderField, input.OrderType))
	}
	// Custom Order
	if input.Orders != nil {
		for _, orderBy := range input.Orders {
			db = db.Order(orderBy.Value)
		}
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

// Statistic SQL
/*
INSERT INTO user_promote_monthly_statistics (year, month, total, male, female, exp_1_3, exp_4_6, exp_7_10, exp_11_15, exp_16_19, exp_20_up, age_13_17, age_18_24, age_25_34, age_35_44, age_45_54, age_55_64, age_65_up)
SELECT
  2022 AS year,
  6 AS month,
  COUNT(DISTINCT users.id) AS total,
  COUNT(DISTINCT CASE WHEN users.sex = 'm' THEN users.id END) AS male,
  COUNT(DISTINCT CASE WHEN users.sex = 'f' THEN users.id END) AS female,
  COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 1 AND 3 THEN users.id END) AS exp_1_3,
  COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 4 AND 6 THEN users.id END) AS exp_4_6,
  COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 7 AND 10 THEN users.id END) AS exp_7_10,
  COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 11 AND 15 THEN users.id END) AS exp_11_15,
  COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 16 AND 19 THEN users.id END) AS exp_16_19,
  COUNT(DISTINCT CASE WHEN trainers.experience >= 20 THEN users.id END) AS exp_20_up,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 13 AND 17 THEN users.id END) AS age_13_17,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 18 AND 24 THEN users.id END) AS age_18_24,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 25 AND 34 THEN users.id END) AS age_25_34,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 35 AND 44 THEN users.id END) AS age_35_44,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 45 AND 54 THEN users.id END) AS age_45_54,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 55 AND 64 THEN users.id END) AS age_55_64,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) >= 65 THEN users.id END) AS age_65_up
FROM users
INNER JOIN trainers ON users.id = trainers.user_id
WHERE YEAR(trainers.create_at) = 2022 AND MONTH(trainers.create_at) = 6
ON DUPLICATE KEY UPDATE
  total = VALUES(total),
  male = VALUES(male),
  female = VALUES(female),
  exp_1_3 = VALUES(exp_1_3),
  exp_4_6 = VALUES(exp_4_6),
  exp_7_10 = VALUES(exp_7_10),
  exp_11_15 = VALUES(exp_11_15),
  exp_16_19 = VALUES(exp_16_19),
  exp_20_up = VALUES(exp_20_up),
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
	err = r.db.Exec("INSERT INTO user_promote_monthly_statistics (year, month, total, male, female, exp_1_3, exp_4_6, exp_7_10, exp_11_15, exp_16_19, exp_20_up, age_13_17, age_18_24, age_25_34, age_35_44, age_45_54, age_55_64, age_65_up) "+
		"SELECT "+
		"? AS year, "+
		"? AS month, "+
		"COUNT(DISTINCT users.id) AS total, "+
		"COUNT(DISTINCT CASE WHEN users.sex = 'm' THEN users.id END) AS male, "+
		"COUNT(DISTINCT CASE WHEN users.sex = 'f' THEN users.id END) AS female, "+
		"COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 1 AND 3 THEN users.id END) AS exp_1_3, "+
		"COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 4 AND 6 THEN users.id END) AS exp_4_6, "+
		"COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 7 AND 10 THEN users.id END) AS exp_7_10, "+
		"COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 11 AND 15 THEN users.id END) AS exp_11_15, "+
		"COUNT(DISTINCT CASE WHEN trainers.experience BETWEEN 16 AND 19 THEN users.id END) AS exp_16_19, "+
		"COUNT(DISTINCT CASE WHEN trainers.experience >= 20 THEN users.id END) AS exp_20_up, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 13 AND 17 THEN users.id END) AS age_13_17, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 18 AND 24 THEN users.id END) AS age_18_24, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 25 AND 34 THEN users.id END) AS age_25_34, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 35 AND 44 THEN users.id END) AS age_35_44, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 45 AND 54 THEN users.id END) AS age_45_54, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) BETWEEN 55 AND 64 THEN users.id END) AS age_55_64, "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(?, '-', ?, '-01'))) >= 65 THEN users.id END) AS age_65_up "+
		"FROM users "+
		"INNER JOIN trainers ON users.id = trainers.user_id "+
		"WHERE YEAR(trainers.create_at) = ? AND MONTH(trainers.create_at) = ? "+
		"ON DUPLICATE KEY UPDATE "+
		"total = VALUES(total), "+
		"male = VALUES(male), "+
		"female = VALUES(female), "+
		"exp_1_3 = VALUES(exp_1_3), "+
		"exp_4_6 = VALUES(exp_4_6), "+
		"exp_7_10 = VALUES(exp_7_10), "+
		"exp_11_15 = VALUES(exp_11_15), "+
		"exp_16_19 = VALUES(exp_16_19), "+
		"exp_20_up = VALUES(exp_20_up), "+
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
