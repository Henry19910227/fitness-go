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
	input.UserID = util.PointerInt64(uid.(int64))
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateBodyRecord(&input)
	ctx.JSON(http.StatusOK, output)
}
