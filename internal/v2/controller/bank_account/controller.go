package bank_account

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/bank_account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver bank_account.Resolver
}

func New(resolver bank_account.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetTrainerBankAccount 獲取教練個人銀行帳戶
// @Summary 獲取教練個人銀行帳戶
// @Description 獲取教練個人銀行帳戶
// @Tags 教練個人_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} bank_account.APIGetTrainerBankAccountOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/bank_account [GET]
func (c *controller) GetTrainerBankAccount(ctx *gin.Context) {
	input := model.APIGetTrainerBankAccountInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetTrainerBankAccount(&input)
	ctx.JSON(http.StatusOK, output)
}
