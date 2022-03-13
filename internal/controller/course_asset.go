package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type CourseAsset struct {
	Base
	courseService service.Course
	planService   service.Plan
	userMidd      midd.User
	courseMidd    midd.Course
}

func NewCourseAsset(baseGroup *gin.RouterGroup, courseService service.Course, planService service.Plan, userMidd midd.User, courseMidd midd.Course) {
	course := CourseAsset{
		courseService: courseService,
		planService:   planService,
		userMidd:      userMidd,
		courseMidd:    courseMidd,
	}
	baseGroup.GET("/course_asset/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		course.GetCourseAsset)
	baseGroup.GET("/course_assets",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		course.GetCourseAssets)
	baseGroup.GET("/course_asset/:course_id/plans",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		course.GetPlanAssets)
}

// GetCourseAsset 獲取課表資源詳細
// @Summary 獲取課表資源詳細
// @Description 獲取課表資源詳細
// @Tags Exercise
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=dto.CourseAsset} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course_asset/{course_id} [GET]
func (a *CourseAsset) GetCourseAsset(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := a.courseService.GetCourseAsset(c, uid, uri.CourseID)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, course, "success!")
}

// GetCourseAssets 獲取課表資源
// @Summary 獲取課表資源
// @Description 獲取課表資源
// @Tags Exercise
// @Accept json
// @Produce json
// @Security fitness_token
// @Param type query int true "搜尋類別(1:進行中課表/2:付費課表)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=[]dto.CourseAssetSummary} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course_assets [GET]
func (a *CourseAsset) GetCourseAssets(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var query validator.GetCourseAssetQuery
	if err := c.ShouldBind(&query); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if query.Type == 1 {
		courses, paging, err := a.courseService.GetProgressCourseAssetSummaries(c, uid, query.Page, query.Size)
		if err != nil {
			a.JSONErrorResponse(c, err)
			return
		}
		a.JSONSuccessPagingResponse(c, courses, paging, "success!")
		return
	}
	courses, paging, err := a.courseService.GetChargeCourseAssetSummaries(c, uid, query.Page, query.Size)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessPagingResponse(c, courses, paging, "success!")
}

// GetPlanAssets 獲取課表資源計畫列表
// @Summary 獲取課表資源計畫列表
// @Description 獲取課表資源計畫列表
// @Tags Exercise
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=[]dto.PlanAsset} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course_asset/{course_id}/plans [GET]
func (a *CourseAsset) GetPlanAssets(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	plans, err := a.planService.GetPlanAssets(c, uid, uri.CourseID)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, plans, "success!")
}
