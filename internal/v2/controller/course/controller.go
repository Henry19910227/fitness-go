package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver course.Resolver
}

func New(resolver course.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetFavoriteCourses 獲取課表收藏列表
// @Summary 獲取課表收藏列表
// @Description 獲取課表收藏列表
// @Tags 收藏_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} course.APIGetFavoriteCoursesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/favorite/courses [GET]
func (c *controller) GetFavoriteCourses(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIGetFavoriteCoursesInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetFavoriteCourses(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSCourses 獲取課表列表
// @Summary 獲取課表列表
// @Description 獲取課表列表
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id query int64 false "課表ID"
// @Param name query string false "課表名稱"
// @Param course_status query int false "課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)"
// @Param sale_type query int false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表)"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} course.APIGetCMSCoursesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/courses [GET]
func (c *controller) GetCMSCourses(ctx *gin.Context) {
	type pagingInput paging.Input
	type orderByInput orderBy.Input
	var query struct {
		model.IDField
		model.NameField
		model.CourseStatusField
		model.SaleTypeField
		pagingInput
		orderByInput
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSCoursesInput{}
	if err := util.Parser(query, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCourses(ctx, &input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSCourse 獲取課表詳細
// @Summary 獲取課表詳細
// @Description 獲取課表詳細
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表ID"
// @Success 200 {object} course.APIGetCMSCourseOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/course/{course_id} [GET]
func (c *controller) GetCMSCourse(ctx *gin.Context) {
	var uri struct {
		model.IDField
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSCourseInput{}
	if err := util.Parser(uri, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCourse(ctx, &input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateCMSCoursesStatus 批量修改課表審核狀態
// @Summary 批量修改課表審核狀態
// @Description 批量修改課表審核狀態
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body course.APIUpdateCMSCoursesStatusInput true "輸入參數"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/courses/course_status [PATCH]
func (c *controller) UpdateCMSCoursesStatus(ctx *gin.Context) {
	var input model.APIUpdateCMSCoursesStatusInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateCMSCoursesStatus(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateCMSCoursesCover 更新課表封面照
// @Summary 更新課表封面照
// @Description 查看封面照 : {Base URL}/v2/resource/course/cover/{Filename}
// @Tags CMS課表管理_v2
// @Security fitness_token
// @Accept mpfd
// @Param course_id path int64 true "課表id"
// @Param cover formData file true "課表封面照"
// @Produce json
// @Success 200 {object} course.APIUpdateCMSCourseCoverOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/course/{course_id}/cover [PATCH]
func (c *controller) UpdateCMSCoursesCover(ctx *gin.Context) {
	var uri struct {
		model.IDRequired
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	file, fileHeader, err := ctx.Request.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIUpdateCMSCourseCoverInput{}
	input.ID = uri.ID
	input.CoverNamed = fileHeader.Filename
	input.File = file
	output := c.resolver.APIUpdateCMSCourseCover(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateUserCourse 創建個人課表
// @Summary 創建個人課表
// @Description 創建個人課表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body course.APICreateUserCourseBody true "輸入參數"
// @Success 200 {object} course.APICreateUserCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied(需訂閱權限)"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course [POST]
func (c *controller) CreateUserCourse(ctx *gin.Context) {
	var input model.APICreateUserCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if input.Body.ScheduleType == model.SingleWorkout {
		output := c.resolver.APICreateUserSingleWorkoutCourse(ctx.MustGet("tx").(*gorm.DB), &input)
		ctx.JSON(http.StatusOK, output)
		return
	}
	output := c.resolver.APICreateUserCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserCourses 獲取用戶個人課表列表
// @Summary 獲取用戶個人課表列表
// @Description 獲取用戶個人課表列表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param type query int false "搜尋類別(1:進行中課表/2:付費課表/3:個人課表)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} course.APIGetUserCoursesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/courses [GET]
func (c *controller) GetUserCourses(ctx *gin.Context) {
	input := model.APIGetUserCoursesInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if input.Query.Type == 1 {
		output := c.resolver.APIGetUserProgressCourses(&input)
		ctx.JSON(http.StatusOK, output)
		return
	}
	if input.Query.Type == 2 {
		output := c.resolver.APIGetUserChargeCourses(&input)
		ctx.JSON(http.StatusOK, output)
		return
	}
	output := c.resolver.APIGetUserPersonalCourses(&input)
	ctx.JSON(http.StatusOK, output)
	return
}

// DeleteUserCourse 刪除個人課表
// @Summary 刪除個人課表
// @Description 刪除個人課表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIDeleteUserCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course/{course_id} [DELETE]
func (c *controller) DeleteUserCourse(ctx *gin.Context) {
	var input model.APIDeleteUserCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateUserCourse 更新個人課表
// @Summary 更新個人課表
// @Description 更新個人課表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body course.APIUpdateUserCourseBody true "輸入參數"
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIUpdateUserCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course/{course_id} [PATCH]
func (c *controller) UpdateUserCourse(ctx *gin.Context) {
	var input model.APIUpdateUserCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateUserCourse(&input)
	ctx.JSON(http.StatusOK, output)
}
