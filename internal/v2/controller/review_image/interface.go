package review_image

import "github.com/gin-gonic/gin"

type Controller interface {
	DeleteCMSReviewImage(ctx *gin.Context)
}
