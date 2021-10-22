package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
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
	baseGroup.GET("/course_products",store.SearchCourseProducts)
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

// SearchCourseProducts 獲取課表產品
// @Summary 獲取課表產品
// @Description 獲取課表產品
// @Tags Store
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "課表名稱(1~20字元)"
// @Param order_type query string false "排序類型(latest:最新/popular:熱門)"
// @Param score query int false "評價(1~5分)-單選"
// @Param level query int false "強度(1:初級/2:中級/3:中高級/4:高級)-複選"
// @Param category query int false "課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選"
// @Param suit query int false "適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選"
// @Param equipment query int false "所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選"
// @Param place query int false "適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選"
// @Param train_target query int false "訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選"
// @Param body_target query int false "體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選"
// @Param sale_type query int false "銷售類型(1:免費課表/2:付費課表/3:訂閱課表)-複選"
// @Param trainer_sex query string false "教練性別(m:男性/f:女性)-複選"
// @Param trainer_skill query int false "教練專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=[]dto.CourseProductSummary} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course_products [GET]
func (s *Store) SearchCourseProducts(c *gin.Context) {
	var query validator.SearchCourseProductsQuery
	if err := c.ShouldBind(&query); err != nil {
		s.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	courses, err := s.storeService.GetCourseProductSummaries(c, &dto.GetCourseProductSummariesParam{
		Name: query.Name,
		OrderType: query.OrderType,
		Score: query.Score,
		Level: query.Level,
		Category: query.Category,
		Suit: query.Suit,
		Equipment: query.Equipment,
		Place: query.Place,
		TrainTarget: query.TrainTarget,
		BodyTarget: query.BodyTarget,
		SaleType: query.SaleType,
		TrainerSex: query.TrainerSex,
		TrainerSkill: query.TrainerSkill,
	}, query.Page, query.Size)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, courses, "success!")
}