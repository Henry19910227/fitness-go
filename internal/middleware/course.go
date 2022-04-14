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
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course, jwtTool tool.JWT, errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo, jwtTool: jwtTool, errHandler: errHandler}
}

func (cm *course) WorkoutSetStatusAccessRange(status []global.CourseStatus, ext []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		var uri validator.WorkoutSetIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			cm.JSONValidatorErrorResponse(c, err)
			c.Abort()
			return
		}
		course := struct {
			Status int `gorm:"column:course_status"`
		}{}
		if err := cm.courseRepo.FindCourseByWorkoutSetID(uri.WorkoutSetID, &course); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		s := status
		if global.Role(role.(int)) == global.AdminRole && ext != nil {
			s = ext
		}
		if !containCourseStatus(s, global.CourseStatus(course.Status)) {
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
			c.Abort()
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		course := struct {
			UserID int64 `gorm:"column:user_id"`
		}{}
		if err := cm.findCourse(c, &course); err != nil {
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

func (cm *course) UserRoleAccessCourseByStatusRange(status []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		course := struct {
			Status int `gorm:"column:course_status"`
		}{}
		if err := cm.findCourse(c, &course); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
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

func (cm *course) AdminAccessCourseByStatusRange(status []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		if global.Role(role.(int)) == global.UserRole {
			return
		}
		course := struct {
			Status int `gorm:"column:course_status"`
		}{}
		if err := cm.findCourse(c, &course); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
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

func (cm *course) CourseStatusVerify(currentStatus func(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var courseUri validator.CourseIDUri
		if err := c.ShouldBindUri(&courseUri); err != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		current, e := currentStatus(c, courseUri.CourseID)
		if e != nil {
			cm.JSONErrorResponse(c, cm.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(e.Code()))))
			c.Abort()
			return
		}
		if !containCourseStatus(allowStatus, current) {
			cm.JSONErrorResponse(c, cm.errHandler.Custom(8999, errors.New("此課表尚未販售")))
			c.Abort()
			return
		}
	}
}

func (cm *course) findCourse(c *gin.Context, course interface{}) error {
	var courseUri validator.CourseIDUri
	var err error
	if err = c.ShouldBindUri(&courseUri); err == nil {
		if err = cm.courseRepo.FindCourseByID(nil, courseUri.CourseID, course); err != nil {
			return err
		}
		return nil
	}
	var planUri validator.PlanIDUri
	if err = c.ShouldBindUri(&planUri); err == nil {
		if err = cm.courseRepo.FindCourseByPlanID(planUri.PlanID, course); err != nil {
			return err
		}
		return nil
	}
	var workoutUri validator.WorkoutIDUri
	if err = c.ShouldBindUri(&workoutUri); err == nil {
		if err = cm.courseRepo.FindCourseByWorkoutID(workoutUri.WorkoutID, course); err != nil {
			return err
		}
		return nil
	}
	var workoutSetUri validator.WorkoutSetIDUri
	if err = c.ShouldBindUri(&workoutSetUri); err == nil {
		if err = cm.courseRepo.FindCourseByWorkoutSetID(workoutSetUri.WorkoutSetID, course); err != nil {
			return err
		}
		return nil
	}
	var actionUri validator.ActionIDUri
	if err = c.ShouldBindUri(&actionUri); err == nil {
		if err = cm.courseRepo.FindCourseByActionID(actionUri.ActionID, course); err != nil {
			return err
		}
		return nil
	}
	return err
}

func containCourseStatus(items []global.CourseStatus, target global.CourseStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}
