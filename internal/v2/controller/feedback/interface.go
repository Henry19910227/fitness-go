package feedback

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateFeedback(ctx *gin.Context)
	GetCMSFeedbacks(ctx *gin.Context)
}
