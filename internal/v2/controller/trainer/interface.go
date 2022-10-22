package trainer

import "github.com/gin-gonic/gin"

type Controller interface {
	GetTrainerProfile(ctx *gin.Context)
	GetTrainer(ctx *gin.Context)
	GetTrainers(ctx *gin.Context)
	GetFavoriteTrainers(ctx *gin.Context)
	UpdateCMSTrainerAvatar(ctx *gin.Context)
}
