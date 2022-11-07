package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course/api_create_favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/favorite_course"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver favorite_course.Resolver
}

func New(resolver favorite_course.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateFavoriteCourse 收藏課表
// @Summary 收藏課表
// @Description 收藏課表
// @Tags 收藏_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} api_create_favorite_course.Output "Success"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/favorite/course/{course_id} [POST]
func (c *controller) CreateFavoriteCourse(ctx *gin.Context) {
	input := api_create_favorite_course.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateFavoriteCourse(&input)
	ctx.JSON(http.StatusOK, output)
}