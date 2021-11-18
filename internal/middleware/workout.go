package middleware

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type workout struct {
	Base
	courseRepo repository.Course
	jwtTool tool.JWT
	errHandler errcode.Handler
}

func NewWorkout(courseRepo repository.Course, jwtTool tool.JWT, errHandler errcode.Handler) Plan {
	return &workout{courseRepo:courseRepo, jwtTool:jwtTool, errHandler: errHandler}
}

func (w *workout) CourseStatusVerify(currentStatus func(c *gin.Context, workoutID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var workoutUri validator.WorkoutIDUri
		var err error
		if err = c.ShouldBindUri(&workoutUri); err != nil {
			w.JSONErrorResponse(c, w.errHandler.Set(c, "json repo", err))
			c.Abort()
			return
		}
		current, e := currentStatus(c, workoutUri.WorkoutID)
		if e != nil {
			w.JSONErrorResponse(c, w.errHandler.Set(c, "workout repo", err))
			c.Abort()
			return
		}
		if !containCourseStatus(allowStatus, current) {
			w.JSONErrorResponse(c, w.errHandler.Set(c, "permission", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}
