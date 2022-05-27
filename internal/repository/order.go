package repository

import (
	"crypto/rand"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"time"
)

type order struct {
	gorm tool.Gorm
}

func NewOrder(gorm tool.Gorm) Order {
	return &order{gorm: gorm}
}

func (o *order) CreateCourseOrder(param *model.CreateOrderParam) (string, error) {
	if param == nil {
		return "", nil
	}
	random := randRange(100000, 999999)
	order := entity.Order{
		ID:          time.Now().Format("20060102150405") + strconv.Itoa(int(random)),
		UserID:      param.UserID,
		Quantity:    1,
		OrderType:   int(global.BuyCourseOrderType),
		OrderStatus: int(global.PendingOrderStatus),
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := o.gorm.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		orderCourse := entity.OrderCourse{
			OrderID:    order.ID,
			SaleItemID: param.SaleItemID,
			CourseID:   param.CourseID,
		}
		if err := tx.Create(&orderCourse).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	return order.ID, nil
}

func (o *order) CreateSubscribeOrder(userID int64) (string, error) {
	random := randRange(100000, 999999)
	order := entity.Order{
		ID:          time.Now().Format("20060102150405") + strconv.Itoa(int(random)),
		UserID:      userID,
		Quantity:    1,
		OrderType:   int(global.SubscribeOrderType),
		OrderStatus: int(global.PendingOrderStatus),
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := o.gorm.DB().Create(&order).Error; err != nil {
		return "", err
	}
	return order.ID, nil
}

func (o *order) UpdateOrderStatus(tx *gorm.DB, orderID string, orderStatus global.OrderStatus) error {
	db := o.gorm.DB()
	if tx != nil {
		db = tx
	}
	if err := db.
		Table("orders").
		Where("id = ?", orderID).
		Update("order_status", int(orderStatus)).Error; err != nil {
		return err
	}
	return nil
}

func (o *order) UpdateOrderSubscribePlan(tx *gorm.DB, orderID string, subscribePlanID int64) error {
	db := o.gorm.DB()
	if tx != nil {
		db = tx
	}
	if err := db.
		Table("order_subscribe_plans").
		Where("order_id = ?", orderID).
		Update("subscribe_plan_id", subscribePlanID).Error; err != nil {
		return err
	}
	return nil
}

func (o *order) FindOrder(orderID string) (*model.Order, error) {
	var order model.Order
	if err := o.gorm.DB().
		Preload("OrderCourse").
		Preload("OrderSubscribe").
		Preload("OrderCourse.SaleItem").
		Preload("OrderCourse.SaleItem.ProductLabel").
		Preload("OrderCourse.Course.Trainer").
		Preload("OrderCourse.Course.Sale").
		Preload("OrderCourse.Course.Sale.ProductLabel").
		Preload("OrderCourse.Course.Review").
		Preload("OrderSubscribe").
		Preload("OrderSubscribe.SubscribePlan").
		Preload("OrderSubscribe.SubscribePlan.ProductLabel").
		Take(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *order) FindOrderByOriginalTransactionID(originalTransactionID string) (*model.Order, error) {
	var order model.Order
	if err := o.gorm.DB().
		Preload("OrderCourse").
		Preload("OrderSubscribe").
		Preload("OrderCourse.SaleItem").
		Preload("OrderCourse.SaleItem.ProductLabel").
		Preload("OrderCourse.Course.Trainer").
		Preload("OrderCourse.Course.Sale").
		Preload("OrderCourse.Course.Sale.ProductLabel").
		Preload("OrderCourse.Course.Review").
		Preload("OrderSubscribe").
		Preload("OrderSubscribe.SubscribePlan").
		Preload("OrderSubscribe.SubscribePlan.ProductLabel").
		Joins("INNER JOIN receipts ON orders.id = receipts.order_id").
		Take(&order, "receipts.original_transaction_id = ?", originalTransactionID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *order) FindOrderByCourseID(userID int64, courseID int64) (*model.Order, error) {
	var order model.Order
	if err := o.gorm.DB().
		Preload("OrderCourse").
		Preload("OrderSubscribe").
		Preload("OrderCourse.SaleItem").
		Preload("OrderCourse.SaleItem.ProductLabel").
		Preload("OrderCourse.Course.Trainer").
		Preload("OrderCourse.Course.Sale").
		Preload("OrderCourse.Course.Sale.ProductLabel").
		Preload("OrderCourse.Course.Review").
		Preload("OrderSubscribe").
		Preload("OrderSubscribe.SubscribePlan").
		Preload("OrderSubscribe.SubscribePlan.ProductLabel").
		Joins("INNER JOIN order_courses ON orders.id = order_courses.order_id").
		Order("orders.create_at DESC").
		Take(&order, "orders.user_id = ? AND order_courses.course_id = ? AND orders.order_status = ?",
			userID, courseID, int(global.PendingOrderStatus)).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *order) FindOrders(userID int64, param *model.FindOrdersParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.Order, error) {
	db := o.gorm.DB().
		Table("orders").
		Preload("OrderCourse").
		Preload("OrderSubscribe").
		Preload("OrderCourse.SaleItem").
		Preload("OrderCourse.SaleItem.ProductLabel").
		Preload("OrderCourse.Course").
		Preload("OrderSubscribe").
		Preload("OrderSubscribe.SubscribePlan").
		Preload("OrderSubscribe.SubscribePlan.ProductLabel").
		Joins("LEFT JOIN order_subscribe_plans ON orders.id = order_subscribe_plans.order_id").
		Joins("LEFT JOIN order_courses ON orders.id = order_courses.order_id")
	query := "1=1 "
	params := make([]interface{}, 0)
	query += "AND orders.user_id = ? "
	params = append(params, userID)
	//加入 order type 篩選條件
	if param.PaymentOrderType != nil {
		query += "AND orders.order_type = ? "
		params = append(params, *param.PaymentOrderType)
	}
	//加入 order status 篩選條件
	if param.OrderStatus != nil {
		query += "AND orders.order_status = ? "
		params = append(params, *param.OrderStatus)
	}
	//加入 subscribe plan id 篩選條件
	if param.SubscribePlanID != nil {
		query += "AND order_subscribe_plans.subscribe_plan_id = ? "
		params = append(params, *param.SubscribePlanID)
	}
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//頁數
	if paging != nil {
		if paging.Offset > 0 {
			db = db.Offset(paging.Offset)
		}
	}
	//筆數
	if paging != nil {
		if paging.Limit > 0 {
			db = db.Limit(paging.Limit)
		}
	}
	orders := make([]*model.Order, 0)
	db = db.Where(query, params...).Find(&orders)
	return orders, nil
}

func (o *order) FindCMSUserOrdersAPIItems(userID int64, result interface{}, orderBy *model.OrderBy, paging *model.PagingParam) (int, error) {
	db := o.gorm.DB().
		Table("orders").
		Select("orders.id AS id", "courses.name AS course_name", "trainers.nickname AS trainer_name", "orders.create_at AS create_at").
		Joins("INNER JOIN order_courses ON orders.id = order_courses.order_id").
		Joins("INNER JOIN courses ON order_courses.course_id = courses.id").
		Joins("INNER JOIN trainers ON courses.user_ID = trainers.user_id")
	db = db.Where("orders.user_id = ? AND orders.order_status = ? AND orders.order_type = ?", userID, global.SuccessOrderStatus, 1)
	//個數
	var amount int64
	db = db.Count(&amount)
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	//查詢數據
	if err := db.Find(result).Error; err != nil {
		return 0, nil
	}
	return int(amount), nil
}

func (o *order) List(listResult interface{}, param *model.FindOrderListParam, preloads []*model.Preload, orderBy *model.OrderBy, paging *model.PagingParam) (int, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 userID 篩選條件
	if param.UserID != nil {
		query += "AND user_id = ? "
		params = append(params, *param.UserID)
	}
	//創建orm
	db := o.gorm.DB().Model(&entity.OrderTemplate{})
	//關聯加載
	for _, preload := range preloads {
		db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
			return db.Select("", preload.Selects...)
		})
	}
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	//查詢數據
	if err := db.Where(query, params...).Find(listResult).Error; err != nil {
		return 0, nil
	}
	//查詢資料數量
	var amount int64
	if err := o.gorm.DB().
		Model(&entity.OrderTemplate{}).
		Where(query, params...).
		Count(&amount).Error; err != nil {
		return 0, err
	}
	return int(amount), nil
}

func randRange(min int64, max int64) int64 {
	if min > max || min < 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}
