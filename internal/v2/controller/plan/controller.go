package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/plan"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver plan.Resolver
}

func New(resolver plan.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSPlans 獲取計畫列表
// @Summary 獲取計畫列表
// @Description 獲取計畫列表
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表ID"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} plan.APIGetCMSPlansOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/course/{course_id}/plans [GET]
func (c *controller) GetCMSPlans(ctx *gin.Context) {
	input := model.APIGetCMSPlansInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSPlans(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateUserPlan 創建個人計畫
// @Summary 創建個人課表計畫
// @Description 創建個人課表計畫
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Param json_body body plan.APICreateUserPlanBody true "輸入參數"
// @Success 200 {object} plan.APICreateUserPlanOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course/{course_id}/plan [POST]
func (c *controller) CreateUserPlan(ctx *gin.Context) {
	var input model.APICreateUserPlanInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateUserPlan(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteUserPlan 刪除個人計畫
// @Summary 刪除個人課表計畫
// @Description 刪除個人課表計畫
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} plan.APIDeleteUserPlanOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/plan/{plan_id} [DELETE]
func (c *controller) DeleteUserPlan(ctx *gin.Context) {
	var input model.APIDeleteUserPlanInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserPlan(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserPlans 獲取用戶個人計畫列表
// @Summary 獲取用戶個人計畫列表
// @Description 獲取用戶個人計畫列表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} plan.APIGetUserPlansOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/course/{course_id}/plans [GET]
func (c *controller) GetUserPlans(ctx *gin.Context) {
	input := model.APIGetUserPlansInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserPlans(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateUserPlan 更新個人計畫
// @Summary 更新個人計畫
// @Description 更新個人計畫
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body plan.APIUpdateUserPlanBody true "輸入參數"
// @Success 200 {object} plan.APIUpdateUserPlanOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/plan/{plan_id} [PATCH]
func (c *controller) UpdateUserPlan(ctx *gin.Context) {
	input := model.APIUpdateUserPlanInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateUserPlan(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateTrainerPlan 創建教練計畫
// @Summary 創建個人教練計畫
// @Description 創建個人教練計畫
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Param json_body body plan.APICreateTrainerPlanBody true "輸入參數"
// @Success 200 {object} plan.APICreateTrainerPlanOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course/{course_id}/plan [POST]
func (c *controller) CreateTrainerPlan(ctx *gin.Context) {
	var input model.APICreateTrainerPlanInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateTrainerPlan(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetTrainerPlans 獲取教練課表計畫列表
// @Summary 獲取教練課表計畫列表
// @Description 獲取教練課表計畫列表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} plan.APIGetTrainerPlansOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/course/{course_id}/plans [GET]
func (c *controller) GetTrainerPlans(ctx *gin.Context) {
	input := model.APIGetTrainerPlansInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetTrainerPlans(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteTrainerPlan 刪除教練課表計畫
// @Summary 刪除教練課表計畫
// @Description 刪除教練課表計畫
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} plan.APIDeleteTrainerPlanOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/plan/{plan_id} [DELETE]
func (c *controller) DeleteTrainerPlan(ctx *gin.Context) {
	var input model.APIDeleteTrainerPlanInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteTrainerPlan(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetStorePlans 獲取商店課表計畫列表
// @Summary 獲取商店課表計畫列表
// @Description 獲取商店課表計畫列表
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} plan.APIGetStorePlansOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/course/{course_id}/plans [GET]
func (c *controller) GetStorePlans(ctx *gin.Context) {
	input := model.APIGetStorePlansInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStorePlans(&input)
	ctx.JSON(http.StatusOK, output)
}
