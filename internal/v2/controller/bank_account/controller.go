package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
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
// @Description 查看銀行帳戶圖片 : {Base URL}/v2/resource/trainer/account_image/{Filename}
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

// UpdateTrainerBankAccount 更新教練個人銀行帳戶
// @Summary 更新教練個人銀行帳戶
// @Description 更新教練個人銀行帳戶
// @Tags 教練個人_v2
// @Accept mpfd
// @Produce json
// @Security fitness_token
// @Param account_name formData string false "戶名"
// @Param bank_code formData string false "銀行代號"
// @Param branch formData string false "分行"
// @Param account formData string false "帳號"
// @Param account_image formData file false "帳戶照片"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/bank_account [PATCH]
func (c *controller) UpdateTrainerBankAccount(ctx *gin.Context) {
	input := model.APIUpdateTrainerBankAccountInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取動作封面
	file, fileHeader, _ := ctx.Request.FormFile("account_image")
	if file != nil {
		input.Form.AccountImageFile = &fileModel.Input{}
		input.Form.AccountImageFile.Named = fileHeader.Filename
		input.Form.AccountImageFile.Data = file
	}
	output := c.resolver.APIUpdateTrainerBankAccount(&input)
	ctx.JSON(http.StatusOK, output)
}
