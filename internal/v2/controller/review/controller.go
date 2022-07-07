package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/review"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver review.Resolver
}

func New(resolver review.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSReviews 獲取評論列表
// @Summary 獲取評論列表
// @Description 獲取評論列表
// @Tags CMS評論管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "課表名稱"
// @Param nickname query string false "用戶暱稱"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} receipt.APIGetCMSReceiptsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/reviews [GET]
func (c *controller) GetCMSReviews(ctx *gin.Context) {
	input := model.APIGetCMSReviewsInput{}
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSReviews(&input)
	ctx.JSON(http.StatusOK, output)
}

