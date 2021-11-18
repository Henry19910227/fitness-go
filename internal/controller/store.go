package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Store struct {
	Base
	storeService service.Store
	courseService service.Course
	planService service.Plan
	workoutService service.Workout
	workoutSetService service.WorkoutSet
}

func NewStore(baseGroup *gin.RouterGroup, storeService service.Store, courseService service.Course, planService service.Plan, workoutService service.Workout, workoutSetService service.WorkoutSet, courseMidd midd.Course, planMidd midd.Plan) {
	store := Store{
		storeService: storeService,
		courseService: courseService,
		planService: planService,
		workoutService: workoutService,
		workoutSetService: workoutSetService,
	}
	baseGroup.GET("/store_home_page",store.GetHomePage)
	baseGroup.GET("/plan_product/:plan_id/workouts",
		planMidd.CourseStatusVerify(planService.GetPlanStatus, []global.CourseStatus{global.Sale}),
		store.GetWorkouts)
}

// GetHomePage 獲取商店首頁資料
// @Summary 獲取商店首頁資料
// @Description 獲取商店首頁資料
// @Tags Store
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.StoreHomePage} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /store_home_page [GET]
func (s *Store) GetHomePage(c *gin.Context) {
	homePage, err := s.storeService.GetHomePage(c)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, homePage, "success!")
}

// GetWorkouts 獲取訓練列表
// @Summary 獲取訓練列表
// @Description 獲取訓練列表
// @Tags Store
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=[]dto.Workout} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /plan_product/{plan_id}/workouts [GET]
func (s *Store) GetWorkouts(c *gin.Context) {
	var uri validator.PlanIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		s.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, err := s.workoutService.GetWorkoutsByPlanID(c, uri.PlanID)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, workouts, "success!")
}