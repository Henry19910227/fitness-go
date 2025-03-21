package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type CourseStatistic struct {
	Base
	courseService service.Course
	userMidd      midd.User
}

func NewCourseStatistic(baseGroup *gin.RouterGroup, courseService service.Course, userMidd midd.User) {
	course := CourseStatistic{
		courseService: courseService,
	}
	baseGroup.GET("/course_statistics",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		course.GetCourseStatistics)
	baseGroup.GET("/course_statistic/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		course.GetCourseStatistic)
}

// GetCourseStatistics 獲取個人課表數據統計列表
// @Summary 獲取個人課表數據統計列表 (API已過時，更新為 /v2/trainer/course/statistics [GET])
// @Description 獲取個人課表數據統計列表
// @Tags Statistic_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessPagingResult{data=[]dto.CourseStatisticSummary} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_statistics [GET]
func (a *CourseStatistic) GetCourseStatistics(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var pagingQuery validator.PagingQuery
	if err := c.ShouldBind(&pagingQuery); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	courses, paging, err := a.courseService.GetCourseStatisticSummaries(c, uid, &dto.PagingParam{
		Page: pagingQuery.Page,
		Size: pagingQuery.Size,
	})
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessPagingResponse(c, courses, paging, "success!")
}

// GetCourseStatistic 獲取個人課表數據詳細
// @Summary 獲取個人課表數據詳細 (API已過時，更新為 /v2/trainer/course/{course_id}/statistic [GET])
// @Description 獲取個人課表數據詳細
// @Tags Statistic_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=dto.CourseStatistic} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_statistic/{course_id} [GET]
func (a *CourseStatistic) GetCourseStatistic(c *gin.Context) {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := a.courseService.GetCourseStatistic(c, uri.CourseID)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, course, "success!")
}
