package banner

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/banner"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver banner.Resolver
}

func New(resolver banner.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateCMSBanner 新增Banner
// @Summary 新增Banner
// @Description 查看Banner照片 : {Base URL}/v2/resource/banner/image/{Filename}
// @Tags CMS內容管理_Banner_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id formData int64 false "課表id"
// @Param user_id formData int64 false "教練id"
// @Param type formData int true "類型(1:課表/2:教練/3:訂閱)"
// @Param image formData file true "圖片"
// @Success 200 {object} banner.APICreateCMSBannerOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/banner [POST]
func (c *controller) CreateCMSBanner(ctx *gin.Context) {
	input := model.APICreateCMSBannerInput{}
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取動作封面
	file, fileHeader, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if file != nil {
		input.ImageFile = &fileModel.Input{}
		input.ImageFile.Named = fileHeader.Filename
		input.ImageFile.Data = file
	}
	output := c.resolver.APICreateCMSBanner(&input)
	ctx.JSON(http.StatusOK, output)
}
