package subscribe_plan

import "github.com/gin-gonic/gin"

type Controller interface {
	GetSubscribePlans(ctx *gin.Context)
}
