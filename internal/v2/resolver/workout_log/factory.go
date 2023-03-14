package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_distance_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_reps_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_rm_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_speed_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_weight_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/min_duration_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_plan_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set_log"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	workoutLogService := workout_log.NewService(db)
	workoutSetLogService := workout_set_log.NewService(db)
	workoutSetService := workout_set.NewService(db)
	courseService := course.NewService(db)
	planService := plan.NewService(db)
	maxDistanceService := max_distance_record.NewService(db)
	maxRepsService := max_reps_record.NewService(db)
	maxRMService := max_rm_record.NewService(db)
	maxSpeedService := max_speed_record.NewService(db)
	maxWeightService := max_weight_record.NewService(db)
	minDurationService := min_duration_record.NewService(db)
	userCourseStatisticService := user_course_statistic.NewService(db)
	userPlanStatisticService := user_plan_statistic.NewService(db)
	return New(workoutLogService, workoutSetLogService, workoutSetService,
		courseService, planService, maxDistanceService, maxRepsService,
		maxRMService, maxSpeedService, maxWeightService,
		minDurationService, userCourseStatisticService, userPlanStatisticService)
}
