package review

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSReviews(ctx *gin.Context)
	UpdateCMSReview(ctx *gin.Context)
	DeleteCMSReview(ctx *gin.Context)
	GetStoreCourseReviews(ctx *gin.Context)
	CreateStoreCourseReview(ctx *gin.Context)
	DeleteStoreCourseReview(ctx *gin.Context)
}
