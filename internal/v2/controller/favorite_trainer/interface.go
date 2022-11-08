package favorite_trainer

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateFavoriteTrainer(ctx *gin.Context)
	DeleteFavoriteTrainer(ctx *gin.Context)
}
