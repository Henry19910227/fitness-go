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

type plan struct {
	Base
	courseRepo repository.Course
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewPlan(courseRepo repository.Course, jwtTool tool.JWT, errHandler errcode.Handler) Plan {
	return &plan{courseRepo: courseRepo, jwtTool:jwtTool, errHandler: errHandler}
}

func (p *plan) CourseStatusVerify(currentStatus func(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var planUri validator.PlanIDUri
		if err := c.ShouldBindUri(&planUri); err != nil {
			p.JSONErrorResponse(c, p.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		current, e := currentStatus(c, planUri.PlanID)
		if e != nil {
			p.JSONErrorResponse(c, p.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(e.Code()))))
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
