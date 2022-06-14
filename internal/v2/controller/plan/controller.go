package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/plan"
	"github.com/gin-gonic/gin"
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
	var uri struct {
		model.CourseIDField
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	type pagingInput paging.Input
	type orderByInput orderBy.Input
	var query struct {
		pagingInput
		orderByInput
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSPlansInput{}
	input.CourseID = uri.CourseID
	if err := util.Parser(query, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSPlans(&input)
	ctx.JSON(http.StatusOK, output)
}
