package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type CourseProduct struct {
	Base
	courseService     service.Course
	planService       service.Plan
	workoutSetService service.WorkoutSet
	actionService     service.Action
	reviewService service.Review
	userMidd      midd.User
	courseMidd    midd.Course
}

func NewCourseProduct(baseGroup *gin.RouterGroup, courseService service.Course, planService service.Plan, workoutSetService service.WorkoutSet, courseMidd midd.Course, userMidd midd.User) {
	course := CourseProduct{
		courseService:     courseService,
		planService:       planService,
		workoutSetService: workoutSetService,
	}
	baseGroup.GET("/course_product/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		course.GetCourseProduct)
	baseGroup.GET("/course_product_structure/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		course.GetCourseProductStructure)
	baseGroup.GET("/course_product/:course_id/plans",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		course.GetPlanProducts)
	baseGroup.GET("/course_products", course.SearchCourseProducts)
}

// GetCourseProduct 獲取課表產品詳細
// @Summary 獲取課表產品詳細
// @Description 獲取課表產品詳細
// @Tags Explore_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=dto.CourseProduct} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_product/{course_id} [GET]
func (p *CourseProduct) GetCourseProduct(c *gin.Context) {
	uid, e := p.GetUID(c)
	if e != nil {
		p.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := p.courseService.GetCourseProductByCourseID(c, uid, uri.CourseID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, course, "success!")
}

// GetCourseProductStructure 獲取課表結構(只限單一訓練課表)
// @Summary 獲取課表結構(只限單一訓練課表)
// @Description 只限單一訓練的課表使用
// @Tags Explore_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=dto.CourseProductStructure} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_product_structure/{course_id} [GET]
func (p *CourseProduct) GetCourseProductStructure(c *gin.Context) {
	uid, e := p.GetUID(c)
	if e != nil {
		p.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := p.courseService.GetCourseProductStructure(c, uid, uri.CourseID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, course, "success!")
}

// SearchCourseProducts 搜尋課表產品列表
// @Summary 搜尋課表產品列表
// @Description 搜尋課表產品列表
// @Tags Explore_v1
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
// @Param sale_type query int false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表)-複選"
// @Param trainer_sex query string false "教練性別(m:男性/f:女性)-複選"
// @Param trainer_skill query int false "教練專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessPagingResult{data=[]dto.CourseProductSummary} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_products [GET]
func (p *CourseProduct) SearchCourseProducts(c *gin.Context) {
	var query validator.SearchCourseProductsQuery
	if err := c.ShouldBind(&query); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	courses, paging, err := p.courseService.GetCourseProductSummaries(c, &dto.GetCourseProductSummariesParam{
		Name:         query.Name,
		OrderType:    query.OrderType,
		Score:        query.Score,
		Level:        query.Level,
		Category:     query.Category,
		Suit:         query.Suit,
		Equipment:    query.Equipment,
		Place:        query.Place,
		TrainTarget:  query.TrainTarget,
		BodyTarget:   query.BodyTarget,
		SaleType:     query.SaleType,
		TrainerSex:   query.TrainerSex,
		TrainerSkill: query.TrainerSkill,
	}, query.Page, query.Size)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessPagingResponse(c, courses, paging, "success!")
}

// GetPlanProducts 獲取課表產品計畫列表
// @Summary 獲取課表產品計畫列表
// @Description 獲取課表產品計畫列表
// @Tags Explore_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=[]dto.Plan} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_product/{course_id}/plans [GET]
func (p *CourseProduct) GetPlanProducts(c *gin.Context) {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	plans, err := p.planService.GetPlansByCourseID(c, uri.CourseID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, plans, "success!")
}
