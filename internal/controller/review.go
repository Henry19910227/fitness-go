package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Review struct {
	Base
	reviewService service.Review
}

func NewReview(baseGroup *gin.RouterGroup, reviewService service.Review, userMiddleware midd.User, courseMiddleware midd.Course)  {
	review := Review{reviewService: reviewService}
	courseGroup := baseGroup.Group("/course")
	courseGroup.POST("/:course_id/submit",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		userMiddleware.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		courseMiddleware.CourseCreatorVerify(),
		courseMiddleware.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		review.SubmitForReview)
}

// SubmitForReview 送審課表
// @Summary 送審課表
// @Description 送審課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /course/{course_id}/submit [POST]
func (r *Review) SubmitForReview(c *gin.Context)  {
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := r.reviewService.CourseSubmit(c, uri.CourseID); err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, nil, "success!")
}




