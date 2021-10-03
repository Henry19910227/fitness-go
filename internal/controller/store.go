package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Store struct {
	Base
	storeService service.Store
}

func NewStore(baseGroup *gin.RouterGroup, storeService service.Store) {
	store := Store{
		storeService: storeService,
	}
	baseGroup.GET("/store_home_page",store.GetHomePage)
	baseGroup.GET("/course_products",store.GetCourseProducts)
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

// GetCourseProducts 獲取課表產品
// @Summary 獲取課表產品
// @Description 獲取課表產品
// @Tags Store
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=dto.StoreHomePage} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course_products [GET]
func (s *Store) GetCourseProducts(c *gin.Context) {
	var query validator.GetLatestCoursesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		s.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	courses, err := s.storeService.GetCourseProduct(c, query.Page, query.Size)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, courses, "success!")
}