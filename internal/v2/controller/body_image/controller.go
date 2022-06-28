package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
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

// GetBodyImages 獲取體態照片列表
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
func (c *controller) GetBodyImages(ctx *gin.Context) {
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

// CreateBodyImage 新增體態照片
// @Summary 新增體態照片
// @Description 查看體態照片 : {Base URL}/v2/resource/body/image/{Filename}
// @Tags 體態紀錄_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param image formData file true "體態照片"
// @Success 200 {object} body_image.APICreateBodyImageOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/body_image [POST]
func (c *controller) CreateBodyImage(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APICreateBodyImageInput{}
	input.UserID = uid.(int64)
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
	output := c.resolver.APICreateBodyImage(&input)
	ctx.JSON(http.StatusOK, output)
}
