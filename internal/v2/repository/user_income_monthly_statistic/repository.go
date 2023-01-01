package user_income_monthly_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_income_monthly_statistic"
	"gorm.io/gorm"
	"time"
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

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_income_monthly_statistics.user_id = ?", *input.UserID)
	}
	// year 篩選條件
	if input.Year != nil {
		db = db.Where("user_income_monthly_statistics.year = ?", *input.Year)
	}
	// month 篩選條件
	if input.Month != nil {
		db = db.Where("user_income_monthly_statistics.month = ?", *input.Month)
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
	db = db.Select("user_income_monthly_statistics.*")
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
		db = db.Order(fmt.Sprintf("user_income_monthly_statistics.%s %s", input.OrderField, input.OrderType))
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

// Statistic https://kind-bass-788.notion.site/user_income_monthly_statistic-067c3292ca294027aea1f73abcdd95e4
func (r *repository) Statistic() (err error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	err = r.db.Exec("INSERT INTO user_income_monthly_statistics " +
		"( " +
		"user_id, " +
		"income, " +
		"year, " +
		"month " +
		") " +
		"SELECT " +
		"a.user_id AS user_id, " +
		"a.income AS income, " +
		"a.year AS year, " +
		"a.month AS month " +
		"FROM " +
		"( " + tableA() + " ) AS a " +
		"ON DUPLICATE KEY UPDATE " +
		"income = a.income", timeStr).Error
	return err
}

func tableA() string {
	return 	"SELECT " +
		"courses.user_id AS user_id, " +
		"SUM(product_labels.twd) AS income, " +
		"MAX(DATE_FORMAT(orders.create_at, '%Y')) AS `year`, " +
		"MAX(DATE_FORMAT(orders.create_at, '%c')) AS `month` " +
		"FROM orders " +
		"INNER JOIN order_courses ON orders.id = order_courses.order_id " +
		"INNER JOIN sale_items ON order_courses.sale_item_id = sale_items.id " +
		"INNER JOIN product_labels ON sale_items.product_label_id = product_labels.id " +
		"INNER JOIN courses ON order_courses.course_id = courses.id " +
		"WHERE DATE_FORMAT(orders.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m') AND orders.order_status = 2 AND orders.order_type = 1 " +
		"GROUP BY courses.user_id"
}

