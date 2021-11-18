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

type plan struct {
	Base
	courseRepo repository.Course
	jwtTool tool.JWT
	errHandler errcode.Handler
}

func NewPlan(courseRepo repository.Course, jwtTool tool.JWT, errHandler errcode.Handler) Plan {
	return &plan{courseRepo:courseRepo, jwtTool:jwtTool, errHandler: errHandler}
}

func (p *plan) CourseStatusVerify(currentStatus func(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var planUri validator.PlanIDUri
		var err error
		if err = c.ShouldBindUri(&planUri); err != nil {
			p.JSONErrorResponse(c, p.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		current, e := currentStatus(c, planUri.PlanID)
		if e != nil {
			p.JSONErrorResponse(c, p.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		if !containCourseStatus(allowStatus, current) {
			p.JSONErrorResponse(c, p.errHandler.Set(c, "permission", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}
