package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	Base
	courseService service.Course
	planService   service.Plan
	actionService service.Action
	userMidd midd.User
	courseMidd midd.Course
}

func NewCourse(baseGroup *gin.RouterGroup,
	courseService service.Course,
	planService service.Plan,
	actionService service.Action,
	userMidd midd.User,
	courseMidd midd.Course) {

	course := &Course{courseService: courseService,
		planService: planService,
		actionService: actionService,
		userMidd: userMidd,
		courseMidd: courseMidd}

	baseGroup.StaticFS("/resource/course/cover", http.Dir("./volumes/storage/course/cover"))

	baseGroup.GET("/courses",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		course.GetCourses)

	baseGroup.POST("/course",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		course.CreateCourse)

	baseGroup.PATCH("/course/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		course.UpdateCourse)

	baseGroup.GET("/course/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		course.GetCourse)

	baseGroup.POST("/course/:course_id/cover",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		course.UploadCourseCover)

	baseGroup.DELETE("/course/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing}),
		course.DeleteCourse)

	baseGroup.POST("/course/:course_id/plan",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		course.CreatePlan)

	baseGroup.GET("/course/:course_id/plans",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		course.GetPlans)

	baseGroup.POST("/course/:course_id/action",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		course.CreateAction)

	baseGroup.GET("/course/:course_id/actions",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		course.SearchActions)
}

// CreateCourse 創建課表
// @Summary 創建課表
// @Description 創建課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param json_body body validator.CreateCourseBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Course} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /course [POST]
func (cc *Course) CreateCourse(c *gin.Context) {
	uid, e := cc.GetUID(c)
	if e != nil {
		cc.JSONValidatorErrorResponse(c, e.Error())
	}
	var body validator.CreateCourseBody
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := cc.courseService.CreateCourse(c, uid, &dto.CreateCourseParam{
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
// @Success 200 {object} model.SuccessResult{data=dto.Course} "更新成功!"
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
	course, err := cc.courseService.UpdateCourse(c, uri.CourseID, &dto.UpdateCourseParam{
		Category: body.Category,
		SaleID: body.SaleID,
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

// GetCourses 獲取我創建的課表
// @Summary 獲取我的課表列表
// @Description 獲取我的課表列表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param status query int false "課表狀態(1:準備中/2:審核中/3:銷售中/4:退審/5:下架)"
// @Success 200 {object} model.SuccessResult{data=[]dto.CourseSummary} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /courses [GET]
func (cc *Course) GetCourses(c *gin.Context) {
	uid, e := cc.GetUID(c)
	if e != nil {
		cc.JSONValidatorErrorResponse(c, e.Error())
	}
	var query validator.CourseStatusQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	courses, err := cc.courseService.GetCourseSummariesByUID(c, uid, query.Status)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, courses, "success")
}

// GetCourse 獲取課表詳細
// @Summary 獲取課表詳細
// @Description 獲取課表詳細
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=dto.Course} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course/{course_id} [GET]
func (cc *Course) GetCourse(c *gin.Context) {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	course, err := cc.courseService.GetCourseDetailByCourseID(c, uri.CourseID)
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
// @Success 200 {object} model.SuccessResult{data=dto.CourseCover} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /course/{course_id}/cover [POST]
func (cc *Course) UploadCourseCover(c *gin.Context) {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("cover")
	if err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := cc.courseService.UploadCourseCoverByID(c, uri.CourseID, &dto.UploadCourseCoverParam{
		File:       file,
		CoverNamed: fileHeader.Filename,
	})
	if e != nil {
		cc.JSONErrorResponse(c, e)
		return
	}
	cc.JSONSuccessResponse(c, result, "success upload")
}

// DeleteCourse 刪除課表
// @Summary 刪除課表
// @Description 刪除課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=dto.CourseID} "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /course/{course_id} [DELETE]
func (cc *Course) DeleteCourse(c *gin.Context) {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := cc.courseService.DeleteCourse(c, uri.CourseID)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, result, "delete success!")
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
// @Success 200 {object} model.SuccessResult{data=plandto.Plan} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /course/{course_id}/plan [POST]
func (cc *Course) CreatePlan(c *gin.Context) {
	var uri validator.CourseIDUri
	var body validator.CreatePlanBody
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	plan, err := cc.planService.CreatePlan(c, uri.CourseID, body.Name)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, plan, "success create plan!")
}

// GetPlans 取得課表內的計畫列表
// @Summary  取得課表內的計畫列表
// @Description  取得課表內的計畫列表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult{data=[]plandto.Plan} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course/{course_id}/plans [GET]
func (cc *Course) GetPlans(c *gin.Context) {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	plans, err := cc.planService.GetPlansByCourseID(c, uri.CourseID)
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, plans, "success!")
}

// CreateAction 創建動作
// @Summary 創建動作
// @Description 創建動作
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Param json_body body validator.CreateActionBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Action} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /course/{course_id}/action [POST]
func (cc *Course) CreateAction(c *gin.Context) {
	var uri validator.CourseIDUri
	var body validator.CreateActionBody
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	action, err := cc.actionService.CreateAction(c, uri.CourseID, &dto.CreateActionParam{
		Name: body.Name,
		Type: body.Type,
		Category: body.Category,
		Body: body.Body,
		Equipment: body.Equipment,
		Intro: body.Intro,
	})
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, action, "success create!")
}

// SearchActions 搜尋動作列表
// @Summary 搜尋動作列表
// @Description 搜尋動作列表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Param name query string false "動作名稱"
// @Param source query string false "動作來源(1:平台動作/2:教練動作)"
// @Param category query string false "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"
// @Param body query string false "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"
// @Param equipment query string false "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:其他)"
// @Success 200 {object} model.SuccessResult{data=[]dto.Action} "查詢成功!"
// @Failure 400 {object} model.ErrorResult "查詢失敗"
// @Router /course/{course_id}/actions [GET]
func (cc *Course) SearchActions(c *gin.Context) {
	var uri validator.CourseIDUri
	var query validator.SearchActionsQuery
	if err := c.ShouldBindUri(&uri); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	actions, err := cc.actionService.SearchActions(c, uri.CourseID, &dto.FindActionsParam{
		Name: query.Name,
		Source: query.Source,
		Category: query.Category,
		Body: query.Body,
		Equipment: query.Equipment,
	})
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, actions, "success!")
}
