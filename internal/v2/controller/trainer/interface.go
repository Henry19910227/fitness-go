package trainer

import "github.com/gin-gonic/gin"

type Controller interface {
	GetFavoriteTrainers(ctx *gin.Context)
	UpdateCMSTrainerAvatar(ctx *gin.Context)
}
