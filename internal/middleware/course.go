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

type course struct {
	Base
	courseRepo repository.Course
	jwtTool tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course, jwtTool tool.JWT, errHandler errcode.Handler) Course {
	return &course{courseRepo:courseRepo, jwtTool:jwtTool, errHandler: errHandler}
}

func (cm *course) WorkoutSetPermission(status []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			return
		}
		var uri validator.WorkoutSetIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			cm.JSONValidatorErrorResponse(c, err)
			c.Abort()
			return
		}
		course := struct {
			UserID int64 `gorm:"column:user_id"`
			Status int `gorm:"column:course_status"`
		}{}
		if err := cm.courseRepo.FindCourseByWorkoutSetID(uri.WorkoutSetID, &course); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		if course.UserID != uid {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
		if !containCourseStatus(status, global.CourseStatus(course.Status)) {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "permission", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (cm *course) CourseCreatorVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			return
		}
		var uri validator.CourseIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			cm.JSONValidatorErrorResponse(c, err)
			return
		}
		course := struct {
			UserID int64 `gorm:"column:user_id"`
		}{}
		if err := cm.courseRepo.FindCourseByID(uri.CourseID, &course); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
			c.Abort()
		}
		if course.UserID != uid {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (cm *course) CourseStatusAccessRange(status []global.CourseStatus, ext []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			return
		}
		var uri validator.CourseIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			cm.JSONValidatorErrorResponse(c, err)
			return
		}
		course := struct {
			Status int `gorm:"column:course_status"`
		}{}
		if err := cm.courseRepo.FindCourseByID(uri.CourseID, &course); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
			c.Abort()
		}
		s := status
		if global.Role(role.(int)) == global.AdminRole && ext != nil{
			s = ext
		}
		if !containCourseStatus(s, global.CourseStatus(course.Status)) {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "permission", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
		}
	}
}

func containCourseStatus(items []global.CourseStatus, target global.CourseStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}
