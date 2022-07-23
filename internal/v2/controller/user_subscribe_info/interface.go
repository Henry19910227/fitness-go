package user_subscribe_info

import "github.com/gin-gonic/gin"

type Controller interface {
	GetUserSubscribeInfo(ctx *gin.Context)
}
