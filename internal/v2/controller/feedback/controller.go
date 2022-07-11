package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/feedback"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver feedback.Resolver
}

func New(resolver feedback.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateFeedback 創建反饋
// @Summary 創建反饋
// @Description 查看反饋圖片 : {Base URL}/v2/resource/feedback/image/{Filename}
// @Tags 意見反饋_v2
// @Accept mpfd
// @Produce json
// @Security fitness_token
// @Param version formData string false "軟體版本"
// @Param platform formData string false "平台(ios/android)"
// @Param os_version formData string false "OS版本"
// @Param phone_model formData string false "手機型號"
// @Param body formData string true "內文"
// @Param feedback_image[] formData file false "反饋圖片"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/feedback [POST]
func (c *controller) CreateFeedback(ctx *gin.Context) {
	input := model.APICreateFeedbackInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	fileDatas := ctx.Request.MultipartForm.File["feedback_image[]"]
	if len(fileDatas) > 5 {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString("超過圖片上傳上限")))
		return
	}
	files := make([]*file.Input, 0)
	for _, fileData := range fileDatas {
		data, _ := fileData.Open()
		f := file.Input{}
		f.Named = fileData.Filename
		f.Data = data
		files = append(files, &f)
	}
	input.Files = files
	output := c.resolver.APICreateFeedback(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSFeedbacks 獲取反饋列表
// @Summary 獲取反饋列表
// @Description 獲取反饋列表
// @Tags CMS平台管理_意見反饋_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param platform query string false "平台(ios/android)"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} feedback.APIGetCMSFeedbacksOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/feedbacks [GET]
func (c *controller) GetCMSFeedbacks(ctx *gin.Context) {
	input := model.APIGetCMSFeedbacksInput{}
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSFeedbacks(&input)
	ctx.JSON(http.StatusOK, output)
}
