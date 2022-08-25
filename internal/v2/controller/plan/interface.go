package plan

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetCMSPlans(ctx *gin.Context)
	CreatePersonalPlan(ctx *gin.Context)
	DeletePersonalPlan(ctx *gin.Context)
}
