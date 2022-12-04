package banner

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_get_cms_banners"
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

// GetBanners 獲取banner列表
// @Summary 獲取banner列表
// @Description 獲取banner列表
// @Tags Banner_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} banner.APIGetBannersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/banners [GET]
func (c *controller) GetBanners(ctx *gin.Context) {
	input := model.APIGetBannersInput{}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetBanners(&input)
	ctx.JSON(http.StatusOK, output)
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

// GetCMSBanners 獲取banner列表
// @Summary 獲取banner列表
// @Description 獲取banner列表
// @Tags CMS內容管理_Banner_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_field query string true "排序欄位 (create_at:創建時間, seq:手動排序順序)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int false "頁數(從第一頁開始)"
// @Param size query int false "筆數"
// @Success 200 {object} api_get_cms_banners.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/banners [GET]
func (c *controller) GetCMSBanners(ctx *gin.Context) {
	input := api_get_cms_banners.Input{}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSBanners(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteCMSBanner 刪除banner
// @Summary 刪除banner
// @Description 刪除banner
// @Tags CMS內容管理_Banner_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param banner_id path int64 true "banner id"
// @Success 200 {object} banner.APIDeleteCMSBannerOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/banner/{banner_id} [DELETE]
func (c *controller) DeleteCMSBanner(ctx *gin.Context) {
	input := model.APIDeleteCMSBannerInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteCMSBanner(&input)
	ctx.JSON(http.StatusOK, output)
}
