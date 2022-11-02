package trainer

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateTrainer(ctx *gin.Context)
	GetTrainerProfile(ctx *gin.Context)
	GetStoreTrainer(ctx *gin.Context)
	GetStoreTrainers(ctx *gin.Context)
	GetFavoriteTrainers(ctx *gin.Context)
	UpdateCMSTrainerAvatar(ctx *gin.Context)
}
