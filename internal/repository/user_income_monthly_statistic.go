package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userIncomeMonthlyStatistic struct {
	gorm tool.Gorm
}

func NewUserIncomeMonthlyStatistic(gorm tool.Gorm) UserIncomeMonthlyStatistic {
	return &userIncomeMonthlyStatistic{gorm: gorm}
}

// CalculateUserIncomeMonthlyCount SQL
//SELECT
//courses.user_id AS user_id,
//SUM(product_labels.twd) AS income,
//MAX(DATE_FORMAT(orders.create_at, '%Y')) AS `year`,
//MAX(DATE_FORMAT(orders.create_at, '%c')) AS `month`
//FROM orders
//INNER JOIN order_courses ON orders.id = order_courses.order_id
//INNER JOIN sale_items ON order_courses.sale_item_id = sale_items.id
//INNER JOIN product_labels ON sale_items.product_label_id = product_labels.id
//INNER JOIN courses ON order_courses.course_id = courses.id
//WHERE
//DATE_FORMAT(orders.create_at, '%Y-%m') = DATE_FORMAT('2022-03-01', '%Y-%m')
//AND
//orders.order_status = 2
//AND
//orders.order_type = 1
//GROUP BY courses.user_id
func (u userIncomeMonthlyStatistic) CalculateUserIncomeMonthlyCount(tx *gorm.DB, date string) ([]*model.UserIncomeMonthlyStatisticResult, error) {
	db := u.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.UserIncomeMonthlyStatisticResult, 0)
	if err := db.Table("orders").
		Select("courses.user_id AS user_id",
			"SUM(product_labels.twd) AS income",
			"MAX(DATE_FORMAT(orders.create_at, '%Y')) AS `year`",
			"MAX(DATE_FORMAT(orders.create_at, '%c')) AS `month`").
		Joins("INNER JOIN order_courses ON orders.id = order_courses.order_id").
		Joins("INNER JOIN sale_items ON order_courses.sale_item_id = sale_items.id").
		Joins("INNER JOIN product_labels ON sale_items.product_label_id = product_labels.id").
		Joins("INNER JOIN courses ON order_courses.course_id = courses.id").
		Where("DATE_FORMAT(orders.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m') "+
			"AND orders.order_status = ? "+
			"AND orders.order_type = ?", date, global.SuccessOrderStatus, global.BuyCourseOrderType).
		Group("courses.user_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (u userIncomeMonthlyStatistic) Save(tx *gorm.DB, values []*model.UserIncomeMonthlyStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := u.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"user_id": v.UserID, "income": v.Income, "month": v.Month, "year": v.Year})
	}
	if err := db.Table("user_income_monthly_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "year"}, {Name: "month"}},
		DoUpdates: clause.AssignmentColumns([]string{"income", "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (u *userIncomeMonthlyStatistic) Find(userID int64, output interface{}) error {
	if err := u.gorm.DB().
		Model(&entity.UserIncomeMonthlyStatistic{}).
		Where("user_id = ?", userID).
		Take(output).Error; err != nil {
		return err
	}
	return nil
}
