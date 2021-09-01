package middleware

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CourseStatus int
const (
	Preparing CourseStatus = 1
	Reviewing = 2
	Sale = 3
	Reject = 4
	Remove = 5
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

func (cm *course) CourseOwnerVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, isExists := c.Get("uid")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "gin", errors.New(strconv.Itoa(errcode.DataNotFound))))
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
			return
		}
		if course.UserID != uid {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (cm *course) WorkoutSetPermission(status []CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, isExists := c.Get("uid")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "gin", errors.New(strconv.Itoa(errcode.DataNotFound))))
			c.Abort()
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
		if !containCourseStatus(status, CourseStatus(course.Status)) {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "permission", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (cm *course) CourseStatusPermission(status []CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		if !containCourseStatus(status, CourseStatus(course.Status)) {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "permission", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
		}
	}
}

func containCourseStatus(items []CourseStatus, target CourseStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}
