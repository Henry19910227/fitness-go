package plan

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetCMSPlans(ctx *gin.Context)
	CreateUserPlan(ctx *gin.Context)
	DeleteUserPlan(ctx *gin.Context)
	GetUserPlans(ctx *gin.Context)
	UpdateUserPlan(ctx *gin.Context)
}
