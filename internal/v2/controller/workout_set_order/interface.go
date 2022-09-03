package workout_set_order

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdateUserWorkoutSetOrders(ctx *gin.Context)
}
