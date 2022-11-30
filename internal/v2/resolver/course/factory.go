package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fcm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_status_update_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseService := course.NewService(db)
	courseStatusLogService := course_status_update_log.NewService(db)
	planService := plan.NewService(db)
	workoutService := workout.NewService(db)
	subscribeInfoService := user_subscribe_info.NewService(db)
	saleItemService := sale_item.NewService(db)
	trainerService := trainer.NewService(db)
	uploadTool := uploader.NewCourseCoverTool()
	redisTool := redis.NewTool()
	fcmTool := fcm.NewTool()
	return New(courseService, courseStatusLogService,
		planService, workoutService, subscribeInfoService,
		saleItemService, trainerService,
		uploadTool, redisTool, fcmTool)
}
