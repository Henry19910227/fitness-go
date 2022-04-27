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

func randRange(min int64, max int64) int64 {
	if min > max || min < 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}
