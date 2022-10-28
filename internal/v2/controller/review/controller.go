package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/review"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
// @Param score query int false "評分(1~5分)"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} receipt.APIGetCMSOrderReceiptsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/reviews [GET]
func (c *controller) GetCMSReviews(ctx *gin.Context) {
	input := model.APIGetCMSReviewsInput{}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSReviews(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateCMSReview 修改評論
// @Summary 修改評論
// @Description 修改評論
// @Tags CMS評論管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param review_id path int64 true "評論id"
// @Param json_body body review.APIUpdateCMSReviewBody true "輸入參數"
// @Success 200 {object} review.APIUpdateCMSReviewOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/review/{review_id} [PATCH]
func (c *controller) UpdateCMSReview(ctx *gin.Context) {
	input := model.APIUpdateCMSReviewInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateCMSReview(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteCMSReview 刪除評論
// @Summary 刪除評論
// @Description 刪除評論
// @Tags CMS評論管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param review_id path int64 true "評論id"
// @Success 200 {object} review.APIDeleteCMSReviewOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/review/{review_id} [DELETE]
func (c *controller) DeleteCMSReview(ctx *gin.Context) {
	input := model.APIDeleteCMSReviewInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteCMSReview(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetStoreCourseReviews 獲取商店課表評論列表
// @Summary 獲取商店課表評論列表
// @Description 查看評論圖 https://www.fitopia-hub.tk/api/v2/resource/course/review/{圖片名}
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Param filter_type query int false "篩選類型(1:全部/2:有照片)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} review.APIGetStoreCourseReviewsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/course/{course_id}/reviews [GET]
func (c *controller) GetStoreCourseReviews(ctx *gin.Context) {
	input := model.APIGetStoreCourseReviewsInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStoreCourseReviews(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateStoreCourseReview 創建商店課表評論
// @Summary 創建商店課表評論
// @Description 查看評論圖 https://www.fitopia-hub.tk/api/v2/resource/course/review/{圖片名}
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Param score formData int true "評分"
// @Param body formData string false "評論內文"
// @Param review_image[] formData file false "評論照片(多張)"
// @Success 200 {object} review.APICreateStoreCourseReviewOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/course/{course_id}/review [POST]
func (c *controller) CreateStoreCourseReview(ctx *gin.Context) {
	input := model.APICreateStoreCourseReviewInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	fileDatas := ctx.Request.MultipartForm.File["review_image[]"]
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
	output := c.resolver.APICreateStoreCourseReview(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteStoreCourseReview 刪除商店課表評論
// @Summary 刪除商店課表評論
// @Description 刪除商店課表評論
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param review_id path int64 true "評論id"
// @Success 200 {object} review.APIDeleteStoreCourseReviewOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/course/review/{review_id} [DELETE]
func (c *controller) DeleteStoreCourseReview(ctx *gin.Context) {
	input := model.APIDeleteStoreCourseReviewInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteStoreCourseReview(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}
