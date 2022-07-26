package order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_asset"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	orderService := order.NewService(db)
	courseService := course.NewService(db)
	orderCourseService := order_course.NewService(db)
	courseAssetService := user_course_asset.NewService(db)
	subscribeInfoService := user_subscribe_info.NewService(db)
	orderSubscribePlanService := order_subscribe_plan.NewService(db)
	return New(orderService, courseService, orderCourseService, courseAssetService, subscribeInfoService, orderSubscribePlanService)
}
