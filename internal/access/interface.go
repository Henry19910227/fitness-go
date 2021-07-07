package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/gin-gonic/gin"
)

type Trainer interface {
	StatusVerify(c *gin.Context, token string) errcode.Error
}

type Course interface {
	CreateVerify(c *gin.Context, token string) errcode.Error
	UpdateVerifyByCourseID(c *gin.Context, token string, courseID int64) errcode.Error
}

type Plan interface {
	CreateVerifyByCourseID(c *gin.Context, token string, courseID int64) errcode.Error
	UpdateVerifyByPlanID(c *gin.Context, token string, planID int64) errcode.Error
}

type Workout interface {
	CreateVerifyByPlanID(c *gin.Context, token string, planID int64) errcode.Error
	UpdateVerifyByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error
}

type WorkoutSet interface {
	CreateVerifyByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error
}

type Action interface {
	CreateVerifyByCourseID(c *gin.Context, token string, courseID int64) errcode.Error
	UpdateVerifyByActionID(c *gin.Context, token string, actionID int64) errcode.Error
}

