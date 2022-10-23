package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseService := course.NewService(db)
	planService := plan.NewService(db)
	workoutService := workout.NewService(db)
	subscribeInfoService := user_subscribe_info.NewService(db)
	saleItemService := sale_item.NewService(db)
	trainerService := trainer.NewService(db)
	uploadTool := uploader.NewCourseCoverTool()
	return New(courseService, planService, workoutService, subscribeInfoService, saleItemService, trainerService, uploadTool)
}
