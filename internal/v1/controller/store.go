package controller

import (
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/gin-gonic/gin"
)

type Store struct {
	Base
	storeService      service.Store
	courseService     service.Course
	planService       service.Plan
	workoutService    service.Workout
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
}

// GetHomePage 獲取商店首頁資料
// @Summary 獲取商店首頁資料 (API已過時，更新為 /v2/store/home_page [GET])
// @Description 獲取商店首頁資料
// @Tags Store_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.StoreHomePage} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/store_home_page [GET]
func (s *Store) GetHomePage(c *gin.Context) {
	homePage, err := s.storeService.GetHomePage(c)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, homePage, "success!")
}