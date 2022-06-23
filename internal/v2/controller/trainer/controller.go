package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/trainer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver trainer.Resolver
}

func New(resolver trainer.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdateCMSTrainerAvatar 更新教練大頭照
// @Summary 更新教練大頭照
// @Description 查看教練大頭照 : {Base URL}/v2/resource/trainer/avatar/{Filename}
// @Tags CMS會員管理_v2
// @Security fitness_token
// @Accept mpfd
// @Param user_id path int64 true "教練id"
// @Param avatar formData file true "教練大頭照"
// @Produce json
// @Success 200 {object} trainer.APIUpdateCMSTrainerAvatarOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/trainer/{user_id}/avatar [PATCH]
func (c *controller) UpdateCMSTrainerAvatar(ctx *gin.Context) {
	var uri struct {
		model.UserIDRequired
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	file, fileHeader, err := ctx.Request.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIUpdateCMSTrainerAvatarInput{}
	input.UserID = uri.UserID
	input.CoverNamed = fileHeader.Filename
	input.File = file
	output := c.resolver.APIUpdateCMSTrainerAvatar(&input)
	ctx.JSON(http.StatusOK, output)
}
