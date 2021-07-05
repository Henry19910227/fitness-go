package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/gin-gonic/gin"
)

type Course interface {
	CheckTrainerValidByUID(c *gin.Context, token string) errcode.Error
	CheckCourseOwnerByCourseID(c *gin.Context, token string, courseID int64) errcode.Error
	CheckPlanOwnerByPlanID(c *gin.Context, token string, planID int64) errcode.Error
	CheckWorkoutOwnerByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error
	CheckActionOwnerByActionID(c *gin.Context, token string, actionID int64) errcode.Error
	CourseValidationByCourseID(c *gin.Context, token string, courseID int64) errcode.Error
	CourseValidationByPlanID(c *gin.Context, token string, planID int64) errcode.Error
	CourseValidationByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error
	CourseValidationByActionID(c *gin.Context, token string, actionID int64) errcode.Error
}