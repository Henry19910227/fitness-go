package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
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
		courseOptional.IDField
		courseOptional.NameField
		courseOptional.CourseStatusField
		courseOptional.SaleTypeField
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
		courseOptional.IDField
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
		courseRequired.IDField
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
// @Param course_id path int64 true "課表id"
// @Param json_body body course.APIUpdateUserCourseBody true "輸入參數"
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

// GetUserCourse 獲取個人課表詳細
// @Summary 獲取個人課表詳細
// @Description 獲取個人課表詳細
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIGetUserCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course/{course_id} [GET]
func (c *controller) GetUserCourse(ctx *gin.Context) {
	var input model.APIGetUserCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserCourseStructure 獲取用戶個人課表結構
// @Summary 獲取用戶個人課表結構
// @Description 獲取用戶個人課表結構
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIGetUserCourseStructureOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course/{course_id}/structure [GET]
func (c *controller) GetUserCourseStructure(ctx *gin.Context) {
	input := model.APIGetUserCourseStructureInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserCourseStructure(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetTrainerCourses 獲取教練課表列表
// @Summary 獲取教練課表列表
// @Description 獲取教練課表列表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_status query int false "課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} course.APIGetTrainerCoursesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/courses [GET]
func (c *controller) GetTrainerCourses(ctx *gin.Context) {
	var input model.APIGetTrainerCoursesInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetTrainerCourses(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateTrainerCourse 創建教練課表
// @Summary 創建教練課表
// @Description 創建教練課表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body course.APICreateTrainerCourseBody true "輸入參數"
// @Success 200 {object} course.APICreateTrainerCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied(需訂閱權限)"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course [POST]
func (c *controller) CreateTrainerCourse(ctx *gin.Context) {
	var input model.APICreateTrainerCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if input.Body.ScheduleType == model.SingleWorkout {
		output := c.resolver.APICreateTrainerSingleWorkoutCourse(ctx.MustGet("tx").(*gorm.DB), &input)
		ctx.JSON(http.StatusOK, output)
		return
	}
	output := c.resolver.APICreateTrainerCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetTrainerCourse 獲取教練課表詳細
// @Summary 獲取教練課表詳細
// @Description 獲取教練課表詳細
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIGetUserCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course/{course_id} [GET]
func (c *controller) GetTrainerCourse(ctx *gin.Context) {
	var input model.APIGetTrainerCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetTrainerCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateTrainerCourse 更新教練課表
// @Summary 更新教練課表
// @Description 查看封面照 : {Base URL}/v2/resource/course/cover/{Filename}
// @Tags 教練課表_v2
// @Security fitness_token
// @Accept mpfd
// @Param course_id path int64 true "課表id"
// @Param sale_type formData int false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)"
// @Param sale_id formData int64 false "銷售 id"
// @Param category formData int false "課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)"
// @Param name formData string false "課表名稱"
// @Param cover formData file false "課表封面照"
// @Param intro formData string false "課表介紹"
// @Param food formData string false "飲食建議"
// @Param level formData int false "強度(1:初級/2:中級/3:中高級/4:高級)"
// @Param suit formData string false "適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)"
// @Param equipment formData string false "所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"
// @Param place formData string false "適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)"
// @Param train_target formData string false "訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)"
// @Param body_target formData string false "體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)"
// @Param notice formData string false "注意事項"
// @Produce json
// @Success 200 {object} course.APIUpdateCMSCourseCoverOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course/{course_id} [PATCH]
func (c *controller) UpdateTrainerCourse(ctx *gin.Context) {
	var input model.APIUpdateTrainerCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取封面圖
	file, fileHeader, _ := ctx.Request.FormFile("cover")
	if file != nil {
		input.Cover = &fileModel.Input{}
		input.Cover.Named = fileHeader.Filename
		input.Cover.Data = file
	}
	output := c.resolver.APIUpdateTrainerCourse(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteTrainerCourse 刪除教練課表
// @Summary 刪除教練課表
// @Description 刪除教練課表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIDeleteTrainerCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course/{course_id} [DELETE]
func (c *controller) DeleteTrainerCourse(ctx *gin.Context) {
	var input model.APIDeleteTrainerCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteTrainerCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// SubmitTrainerCourse 送審教練課表
// @Summary 送審教練課表
// @Description 送審教練課表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APISubmitTrainerCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course/{course_id}/submit [POST]
func (c *controller) SubmitTrainerCourse(ctx *gin.Context) {
	var input model.APISubmitTrainerCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APISubmitTrainerCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetStoreCourse 獲取商店課表詳細
// @Summary 獲取商店課表詳細
// @Description 獲取商店課表詳細
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIGetUserCourseOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/course/{course_id} [GET]
func (c *controller) GetStoreCourse(ctx *gin.Context) {
	var input model.APIGetStoreCourseInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStoreCourse(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetStoreCourses 獲取商店課表列表
// @Summary 獲取商店課表列表
// @Description 獲取商店課表列表
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "課表名稱(1~40字元)"
// @Param score query int false "評價(1~5分)-單選"
// @Param level query string false "強度(1:初級/2:中級/3:中高級/4:高級)-複選"
// @Param category query string false "課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選"
// @Param suit query string false "適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選"
// @Param equipment query string false "所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選"
// @Param place query string false "適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選"
// @Param train_target query string false "訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選"
// @Param body_target query string false "體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選"
// @Param sale_type query string false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表)-複選"
// @Param trainer_sex query string false "教練性別(m:男性/f:女性)-複選"
// @Param trainer_skill query string false "教練專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
// @Param order_field query string false "排序欄位 (latest:最新/popular:熱門)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} course.APIGetStoreCoursesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/courses [GET]
func (c *controller) GetStoreCourses(ctx *gin.Context) {
	var input model.APIGetStoreCoursesInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStoreCourses(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetStoreCourseStructure 獲取商店課表結構
// @Summary 獲取商店課表結構
// @Description 獲取商店課表結構
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} course.APIGetStoreCourseStructureOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/course/{course_id}/structure [GET]
func (c *controller) GetStoreCourseStructure(ctx *gin.Context) {
	input := model.APIGetStoreCourseStructureInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStoreCourseStructure(&input)
	ctx.JSON(http.StatusOK, output)
}
