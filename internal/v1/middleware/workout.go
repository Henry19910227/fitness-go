package middleware

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type workout struct {
	Base
	courseRepo repository.Course
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewWorkout(courseRepo repository.Course, jwtTool tool.JWT, errHandler errcode.Handler) Plan {
	return &workout{courseRepo: courseRepo, jwtTool:jwtTool, errHandler: errHandler}
}

func (w *workout) CourseStatusVerify(currentStatus func(c *gin.Context, workoutID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var workoutUri validator.WorkoutIDUri
		if err := c.ShouldBindUri(&workoutUri); err != nil {
			w.JSONErrorResponse(c, w.errHandler.Set(c, "json repo", err))
			c.Abort()
			return
		}
		current, e := currentStatus(c, workoutUri.WorkoutID)
		if e != nil {
			w.JSONErrorResponse(c, w.errHandler.Set(c, "workout repo", errors.New(strconv.Itoa(e.Code()))))
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
