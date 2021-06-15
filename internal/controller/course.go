package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto/coursedto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	Base
	courseService service.Course
	planService service.Plan
}

func NewCourse(baseGroup *gin.RouterGroup, courseService service.Course, planService service.Plan, userMiddleware gin.HandlerFunc) {

	course := &Course{courseService: courseService, planService: planService}

	baseGroup.StaticFS("/resource/course/cover", http.Dir("./volumes/storage/course/cover"))
	courseGroup := baseGroup.Group("/course")
	courseGroup.Use(userMiddleware)
	courseGroup.POST("", course.CreateCourse)
	courseGroup.PATCH("/:course_id", course.UpdateCourse)
	courseGroup.GET("/list", course.GetCourseList)
	courseGroup.GET("/:course_id", course.GetCourse)
	courseGroup.POST("/:course_id/cover", course.UploadCourseCover)
	courseGroup.POST("/:course_id/plan", course.CreatePlan)
}

// CreateCourse 創建課表
// @Summary 創建課表
// @Description 創建課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param json_body body validator.CreateCourseBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=coursedto.CreateResult} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /course [POST]
func (cc *Course) CreateCourse(c *gin.Context) {
	var header validator.TokenHeader
	var body validator.CreateCourseBody
	if err := c.ShouldBindHeader(&header); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := cc.courseService.CreateCourseByToken(c, header.Token, &coursedto.CreateCourseParam{
		Name: body.Name,
		Level: body.Level,
		Category: body.Category,
		ScheduleType: body.ScheduleType,
	})
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, result, "創建成功!")
}

// UpdateCourse 更新課表
// @Summary 更新課表
// @Description 更新課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Param json_body body validator.UpdateCourseBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=coursedto.Course} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /course/{course_id} [PATCH]
func (cc *Course) UpdateCourse(c *gin.Context) {
	var uri validator.CourseIDUri
	var body validator.UpdateCourseBody
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := cc.courseService.UpdateCourse(c, uri.CourseID, &coursedto.UpdateCourseParam{
		Category: body.Category,
		ScheduleType: body.ScheduleType,
		SaleType: body.SaleType,
		Price: body.Price,
		Name: body.Name,
		Intro: body.Intro,
		Food: body.Food,
		Level: body.Level,
		Suit: body.Suit,
		Equipment: body.Equipment,
		Place: body.Place,
		TrainTarget: body.TrainTarget,
		BodyTarget: body.BodyTarget,
		Notice: body.Notice,
	})
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, course, "更新成功!")
}

// GetCourseList 獲取我的課表列表
// @Summary 獲取我的課表列表
// @Description 獲取我的課表列表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Success 200 {object} model.SuccessResult{data=[]coursedto.Course} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course/list [GET]
func (cc *Course) GetCourseList(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	courses, err := cc.courseService.GetCoursesByToken(c, header.Token)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, courses, "success")
}

// GetCourse 以id獲取課表
// @Summary 以id獲取課表
// @Description 以id獲取課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=coursedto.Course} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course/{course_id} [GET]
func (cc *Course) GetCourse(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.CourseIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := cc.courseService.GetCourseByTokenAndCourseID(c, header.Token, uri.CourseID)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, course, "獲取成功")
}

// UploadCourseCover 上傳課表封面照
// @Summary 上傳課表封面照
// @Description 查看封面照 : https://www.fitness-app.tk/api/v1/resource/course/cover/{圖片名}
// @Tags Course
// @Security fitness_user_token
// @Accept mpfd
// @Param course_id path int64 true "課表id"
// @Param cover formData file true "課表封面照"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=coursedto.CourseCover} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /course/{course_id}/cover [POST]
func (cc *Course) UploadCourseCover(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.CourseIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("cover")
	if err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := cc.courseService.UploadCourseCoverByID(c, uri.CourseID, &coursedto.UploadCourseCoverParam{
		File:       file,
		CoverNamed: fileHeader.Filename,
	})
	if e != nil {
		cc.JSONErrorResponse(c, e)
		return
	}
	cc.JSONSuccessResponse(c, result, "success upload")
}

// CreatePlan 創建計畫
// @Summary 創建計畫
// @Description 創建計畫
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Param json_body body validator.CreatePlanBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=plandto.PlanID} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /course/{course_id}/plan [POST]
func (cc *Course) CreatePlan(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.CourseIDUri
	var body validator.CreatePlanBody
	if err := c.ShouldBindHeader(&header); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := cc.planService.CreatePlanByToken(c, header.Token, uri.CourseID, body.Name)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, result, "success create plan!")
}