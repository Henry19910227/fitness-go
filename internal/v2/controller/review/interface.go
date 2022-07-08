package review

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSReviews(ctx *gin.Context)
	UpdateCMSReview(ctx *gin.Context)
	DeleteCMSReview(ctx *gin.Context)
}
