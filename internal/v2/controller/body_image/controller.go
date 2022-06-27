package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	bodyImage "github.com/Henry19910227/fitness-go/internal/v2/resolver/body_image"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver bodyImage.Resolver
}

func New(resolver bodyImage.Resolver) Controller {
	return &controller{resolver: resolver}
}

// APIGetBodyImages 獲取體態照片列表
// @Summary 獲取體態照片列表
// @Description 獲取體態照片列表
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} body_image.APIGetBodyImagesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_images [GET]
func (c *controller) APIGetBodyImages(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIGetBodyImagesInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetBodyImages(&input)
	ctx.JSON(http.StatusOK, output)
}
