package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user.Resolver
}

func New(resolver user.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdatePassword 修改密碼
// @Summary 修改密碼
// @Description 修改密碼
// @Tags 用戶_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body user.APIUpdatePasswordBody true "輸入參數"
// @Success 200 {object} user.APIUpdatePasswordOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/password [PATCH]
func (c *controller) UpdatePassword(ctx *gin.Context) {
	var input model.APIUpdatePasswordInput
	input.ID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdatePassword(&input)
	ctx.JSON(http.StatusOK, output)
}
