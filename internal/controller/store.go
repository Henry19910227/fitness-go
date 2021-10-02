package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
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