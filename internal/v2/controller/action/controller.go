package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/action"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver action.Resolver
}

func New(resolver action.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSActions 獲取動作列表
// @Summary 獲取動作列表
// @Description 獲取動作列表
// @Tags CMS內容管理_動作庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} action.APIGetCMSActionsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/actions [GET]
func (c *controller) GetCMSActions(ctx *gin.Context) {
	var query struct {
		paging.Input
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSActionsInput{}
	if err := util.Parser(query, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSActions(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateCMSAction 創建動作
// @Summary 創建動作
// @Description 創建動作
// @Tags CMS內容管理_動作庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name formData string true "動作名稱(1~20字元)"`
// @Param type formData int true "紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)"`
// @Param category formData int true "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"`
// @Param body formData int true "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"`
// @Param equipment formData int true "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"`
// @Param intro formData string true "動作介紹(1~400字元)"`
// @Param cover formData file true "課表封面照"
// @Param video formData file false "影片檔"
// @Success 200 {object} action.APICreateCMSActionOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/action [POST]
func (c *controller) CreateCMSAction(ctx *gin.Context) {
	input := model.APICreateCMSActionInput{}
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取動作封面
	file, fileHeader, err := ctx.Request.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if file != nil {
		input.CoverFile = &fileModel.Input{}
		input.CoverFile.Named = fileHeader.Filename
		input.CoverFile.Data = file
	}
	//獲取動作影片
	file, fileHeader, _ = ctx.Request.FormFile("video")
	if file != nil {
		input.VideoFile = &fileModel.Input{}
		input.VideoFile.Named = fileHeader.Filename
		input.VideoFile.Data = file
	}
	output := c.resolver.APICreateCMSAction(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateCMSAction 修改動作
// @Summary 修改動作
// @Description 查看封面照 : {Base URL}/v2/resource/action/cover/{Filename} 查看影片 : {Base URL}/v2/resource/action/video/{Filename}
// @Tags CMS內容管理_動作庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Param name formData string false "動作名稱(1~20字元)"`
// @Param intro formData string false "動作介紹(1~400字元)"`
// @Param status formData int false "動作狀態(0:下架/1:上架)"`
// @Param cover formData file false "課表封面照"
// @Param video formData file false "影片檔"
// @Success 200 {object} base.Output "更新成功!"
// @Failure 400 {object} base.Output "更新失敗"
// @Router /v2/cms/action/{action_id} [PATCH]
func (c *controller) UpdateCMSAction(ctx *gin.Context) {
	input := model.APIUpdateCMSActionInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取動作封面
	file, fileHeader, _ := ctx.Request.FormFile("cover")
	if file != nil {
		input.CoverFile = &fileModel.Input{}
		input.CoverFile.Named = fileHeader.Filename
		input.CoverFile.Data = file
	}
	//獲取動作影片
	file, fileHeader, _ = ctx.Request.FormFile("video")
	if file != nil {
		input.VideoFile = &fileModel.Input{}
		input.VideoFile.Named = fileHeader.Filename
		input.VideoFile.Data = file
	}
	output := c.resolver.APIUpdateCMSAction(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateUserAction 新增個人動作
// @Summary 新增個人動作
// @Description 查看封面照 : {Base URL}/v2/resource/action/cover/{Filename} 查看影片 : {Base URL}/v2/resource/action/video/{Filename}
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name formData string true "動作名稱(1~20字元)"`
// @Param type formData int true "紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)"`
// @Param category formData int true "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"`
// @Param body formData int true "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"`
// @Param equipment formData int true "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"`
// @Param intro formData string true "動作介紹(1~400字元)"`
// @Param cover formData file true "課表封面照"
// @Param video formData file false "影片檔"
// @Success 200 {object} action.APICreateUserActionOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/action [POST]
func (c *controller) CreateUserAction(ctx *gin.Context) {
	input := model.APICreateUserActionInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取封面圖
	file, fileHeader, err := ctx.Request.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if file != nil {
		input.Cover = &fileModel.Input{}
		input.Cover.Named = fileHeader.Filename
		input.Cover.Data = file
	}
	//獲取訓練影片
	file, fileHeader, _ = ctx.Request.FormFile("video")
	if file != nil {
		input.Video = &fileModel.Input{}
		input.Video.Named = fileHeader.Filename
		input.Video.Data = file
	}
	output := c.resolver.APICreateUserAction(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateUserAction 更新個人動作
// @Summary 更新個人動作
// @Description 更新個人動作
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作 id"`
// @Param name formData string false "動作名稱(1~20字元)"`
// @Param category formData int false "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"`
// @Param body formData int false "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"`
// @Param equipment formData int false "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"`
// @Param intro formData string false "動作介紹(1~400字元)"`
// @Param cover formData file false "課表封面照"
// @Param video formData file false "影片檔"
// @Success 200 {object} action.APIUpdateUserActionOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/action/{action_id} [PATCH]
func (c *controller) UpdateUserAction(ctx *gin.Context) {
	input := model.APIUpdateUserActionInput{}
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
	//獲取訓練影片
	file, fileHeader, _ = ctx.Request.FormFile("video")
	if file != nil {
		input.Video = &fileModel.Input{}
		input.Video.Named = fileHeader.Filename
		input.Video.Data = file
	}
	output := c.resolver.APIUpdateUserAction(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserActions 獲取個人動作列表
// @Summary 獲取個人動作列表
// @Description 獲取個人動作列表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "動作名稱"
// @Param source query string false "動作來源(1:平台動作/3:個人動作)"
// @Param category query string false "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"
// @Param body query string false "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"
// @Param equipment query string false "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} action.APIGetUserActionsOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/actions [GET]
func (c *controller) GetUserActions(ctx *gin.Context) {
	var input model.APIGetUserActionsInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserActions(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteUserAction 刪除個人動作
// @Summary 刪除個人動作
// @Description 刪除個人動作
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} action.APIDeleteUserActionOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/action/{action_id} [DELETE]
func (c *controller) DeleteUserAction(ctx *gin.Context) {
	var input model.APIDeleteUserActionInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserAction(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteUserActionVideo 刪除個人動作影片
// @Summary 刪除個人動作影片
// @Description 刪除個人動作影片
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} action.APIDeleteUserActionVideoOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/action/{action_id}/video [DELETE]
func (c *controller) DeleteUserActionVideo(ctx *gin.Context) {
	var input model.APIDeleteUserActionVideoInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserActionVideo(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateTrainerAction 新增教練動作
// @Summary 新增教練動作
// @Description 查看封面照 : {Base URL}/v2/resource/action/cover/{Filename} 查看影片 : {Base URL}/v2/resource/action/video/{Filename}
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name formData string true "動作名稱(1~20字元)"`
// @Param type formData int true "紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)"`
// @Param category formData int true "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"`
// @Param body formData int true "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"`
// @Param equipment formData int true "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"`
// @Param intro formData string true "動作介紹(1~400字元)"`
// @Param cover formData file true "課表封面照"
// @Param video formData file false "影片檔"
// @Success 200 {object} action.APICreateTrainerActionOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/action [POST]
func (c *controller) CreateTrainerAction(ctx *gin.Context) {
	input := model.APICreateTrainerActionInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取封面圖
	file, fileHeader, err := ctx.Request.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if file != nil {
		input.Cover = &fileModel.Input{}
		input.Cover.Named = fileHeader.Filename
		input.Cover.Data = file
	}
	//獲取訓練影片
	file, fileHeader, _ = ctx.Request.FormFile("video")
	if file != nil {
		input.Video = &fileModel.Input{}
		input.Video.Named = fileHeader.Filename
		input.Video.Data = file
	}
	output := c.resolver.APICreateTrainerAction(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetTrainerActions 獲取教練動作列表
// @Summary 獲取教練動作列表
// @Description 獲取教練動作列表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "動作名稱"
// @Param source query string false "動作來源(1:平台動作/2:教練動作)"
// @Param category query string false "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"
// @Param body query string false "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"
// @Param equipment query string false "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} action.APIGetTrainerActionsOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/actions [GET]
func (c *controller) GetTrainerActions(ctx *gin.Context) {
	var input model.APIGetTrainerActionsInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetTrainerActions(&input)
	ctx.JSON(http.StatusOK, output)
}
