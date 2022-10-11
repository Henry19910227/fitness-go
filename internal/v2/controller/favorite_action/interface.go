package favorite_action

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateFavoriteAction(ctx *gin.Context)
}
