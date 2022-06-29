package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	body "github.com/Henry19910227/fitness-go/internal/v2/resolver/body_record"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver body.Resolver
}

func New(resolver body.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateBodyRecord 創建體態紀錄
// @Summary 創建體態紀錄
// @Description 創建體態紀錄
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body body_record.APICreateBodyRecordBody true "輸入參數"
// @Success 200 {object} body_record.APICreateBodyRecordOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_record [POST]
func (c *controller) CreateBodyRecord(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APICreateBodyRecordInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateBodyRecord(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetBodyRecords 獲取體態紀錄列表
// @Summary 獲取體態紀錄列表
// @Description 獲取體態紀錄列表
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param record_type query int true "紀錄類型(1:體重紀錄/2:體脂紀錄/3:胸圍紀錄/4:腰圍紀錄/5:臀圍紀錄/6:身高紀錄/7:臂圍紀錄/8:小臂圍紀錄/9:肩圍紀錄/10:大腿圍紀錄/11:小腿圍紀錄/12:頸圍紀錄"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} body_record.APIGetBodyRecordsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_records [GET]
func (c *controller) GetBodyRecords(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIGetBodyRecordsInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetBodyRecords(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetBodyRecordsLatest 獲取各類型最新體態紀錄列表
// @Summary 獲取各類型最新體態紀錄列表
// @Description 獲取各類型最新體態紀錄列表
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} body_record.APIGetBodyRecordsLatestOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_records/latest [GET]
func (c *controller) GetBodyRecordsLatest(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIGetBodyRecordsLatestInput{}
	input.UserID = uid.(int64)
	output := c.resolver.APIGetBodyRecordsLatest(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateBodyRecord 修改體態紀錄
// @Summary 修改體態紀錄
// @Description 修改體態紀錄
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param body_record_id path int64 true "紀錄id"
// @Param json_body body body_record.APIUpdateBodyRecordBody true "輸入參數"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_record/{body_record_id} [PATCH]
func (c *controller) UpdateBodyRecord(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIUpdateBodyRecordInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateBodyRecord(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteBodyRecord 刪除體態紀錄
// @Summary 刪除體態紀錄
// @Description 刪除體態紀錄
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param body_record_id path int64 true "紀錄id"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_record/{body_record_id} [DELETE]
func (c *controller) DeleteBodyRecord(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIDeleteBodyRecordInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteBodyRecord(&input)
	ctx.JSON(http.StatusOK, output)
}
