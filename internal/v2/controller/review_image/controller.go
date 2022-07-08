package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/review_image"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver review_image.Resolver
}

func New(resolver review_image.Resolver) Controller {
	return &controller{resolver: resolver}
}

// DeleteCMSReviewImage 刪除評論照片
// @Summary 刪除評論照片
// @Description 刪除評論照片
// @Tags CMS評論管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param review_image_id path int64 true "評論照片id"
// @Success 200 {object} review_image.APIDeleteCMSReviewImageOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/review_image/{review_image_id} [DELETE]
func (c *controller) DeleteCMSReviewImage(ctx *gin.Context) {
	input := model.APIDeleteCMSReviewImageInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteCMSReviewImage(&input)
	ctx.JSON(http.StatusOK, output)
}
