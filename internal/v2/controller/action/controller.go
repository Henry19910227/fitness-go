package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/action"
	"github.com/gin-gonic/gin"
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
	input.CoverFile = &fileModel.Input{}
	input.CoverFile.Named = fileHeader.Filename
	input.CoverFile.Data = file

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
