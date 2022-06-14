package plan

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSPlans(ctx *gin.Context)
}
