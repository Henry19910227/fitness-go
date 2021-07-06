package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/gin-gonic/gin"
)

type Course interface {
	CheckCreateAllow(c *gin.Context, token string) errcode.Error
	CheckEditAllowByCourseID(c *gin.Context, token string, courseID int64) errcode.Error
	CheckEditAllowByPlanID(c *gin.Context, token string, planID int64) errcode.Error
	CheckEditAllowByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error
	CheckEditAllowByActionID(c *gin.Context, token string, actionID int64) errcode.Error
}