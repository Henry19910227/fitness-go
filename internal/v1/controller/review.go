package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Review struct {
	Base
	courseService service.Course
	reviewService service.Review
}

func NewReview(baseGroup *gin.RouterGroup,
	courseService service.Course,
	reviewService service.Review,
	userMidd midd.User,
	courseMidd midd.Course,
	reviewMidd midd.Review) {

	review := &Review{courseService: courseService, reviewService: reviewService}

	baseGroup.StaticFS("/resource/course/review", http.Dir("./volumes/storage/course/review"))
	baseGroup.POST("/course_product/:course_id/review",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		review.CreateReview)

	baseGroup.GET("/review/:review_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		review.GetReview)

	baseGroup.GET("/course_product/:course_id/reviews",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		review.GetReviews)

	baseGroup.DELETE("/review/:review_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		reviewMidd.ReviewCreatorVerify(reviewService.GetReviewOwner),
		review.DeleteReview)
}

// CreateReview 創建評論
// @Summary 創建評論 (API已經過時，更新為 /v2/store/course/{course_id}/review [POST])
// @Description 查看評論圖 https://www.fitopia-hub.tk/api/v1/resource/course/review/{圖片名}
// @Tags Review_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Param score formData int true "評分"
// @Param body formData string false "評論內文"
// @Param review_images formData file false "評論照片(多張)"
// @Success 200 {object} model.SuccessResult{data=dto.Review} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/course_product/{course_id}/review [POST]
func (r *Review) CreateReview(c *gin.Context) {
	var uri validator.CourseIDUri
	var form validator.CreateReviewForm
	uid, e := r.GetUID(c)
	if e != nil {
		r.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBind(&form); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//獲取評論照片
	files := c.Request.MultipartForm.File["review_images"]
	var reviewImages []*dto.File
	for _, f := range files {
		data, _ := f.Open()
		file := &dto.File{
			FileNamed: f.Filename,
			Data: data,
		}
		reviewImages = append(reviewImages, file)
	}
	review, err := r.reviewService.CreateReview(c, &dto.CreateReviewParam{
		CourseID: uri.CourseID,
		UserID: uid,
		Score: form.Score,
		Body: form.Body,
		Images: reviewImages,
	})
	if err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, review, "success!")
}

// GetReviews 獲取評論列表
// @Summary 獲取評論列表 (API已經過時，更新為 /v2/store/course/{course_id}/reviews [GET])
// @Description 查看評論圖 https://www.fitopia-hub.tk/api/v1/resource/course/review/{圖片名}
// @Tags Review_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Param filter_type query int false "篩選類型(1:全部/2:有照片)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessPagingResult{data=[]dto.Review} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/course_product/{course_id}/reviews [GET]
func (r *Review) GetReviews(c *gin.Context) {
	var uri validator.CourseIDUri
	var reviewQuery validator.GetReviewsQuery
	var pageQuery validator.PagingQuery
	uid, e := r.GetUID(c)
	if e != nil {
		r.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBind(&reviewQuery); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBind(&pageQuery); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	reviews, paging, err := r.reviewService.GetReviews(c, uid, &dto.GetReviewsParam{
		CourseID:   uri.CourseID,
		FilterType: global.ReviewFilterType(reviewQuery.FilterType),
	}, pageQuery.Page, pageQuery.Size)
	if err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessPagingResponse(c, reviews, paging, "success!")
}

// GetReview 獲取評論
// @Summary 獲取評論 (API已經過時，更新為 /v2/store/course/review/{review_id} [GET])
// @Description 查看評論圖 https://www.fitopia-hub.tk/api/v1/resource/course/review/{圖片名}
// @Tags Review_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param review_id path int64 true "評論id"
// @Success 200 {object} model.SuccessResult{data=dto.Review} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/review/{review_id} [GET]
func (r *Review) GetReview(c *gin.Context) {
	var uri validator.ReviewIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	review, err := r.reviewService.GetReview(c, uri.ReviewID)
	if err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, review, "success!")
}

// DeleteReview 刪除評論
// @Summary 刪除評論 (API已經過時，更新為 /v2/store/course/review/{review_id} [DELETE])
// @Description 刪除評論
// @Tags Review_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param review_id path int64 true "評論id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/review/{review_id} [DELETE]
func (r *Review) DeleteReview(c *gin.Context) {
	var uri validator.ReviewIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := r.reviewService.DeleteReview(c, uri.ReviewID); err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, nil, "delete success!")
}