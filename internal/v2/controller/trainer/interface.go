package trainer

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdateCMSTrainerAvatar(ctx *gin.Context)
}
