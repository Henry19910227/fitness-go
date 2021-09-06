package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type set struct {
	courseRepo repository.Course
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewWorkoutSet(courseRepo repository.Course,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) WorkoutSet {
	return &set{courseRepo: courseRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (s *set) CreateVerifyByWorkoutID(c *gin.Context, uid int64, workoutID int64) errcode.Error {
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := s.courseRepo.FindCourseByWorkoutID(workoutID, &course); err != nil {
		return s.errHandler.Set(c, "access", err)
	}
	if course.UserID != uid {
		return s.errHandler.PermissionDenied()
	}
	if !(course.Status == 1 || course.Status == 4) {
		return s.errHandler.PermissionDenied()
	}
	return nil
}

func (s *set) UpdateVerifyByWorkoutSetID(c *gin.Context, token string, setID int64) errcode.Error {
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return s.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := s.courseRepo.FindCourseByWorkoutSetID(setID, &course); err != nil {
		s.logger.Set(c, handler.Error, "CourseRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	if course.UserID != uid {
		return s.errHandler.PermissionDenied()
	}
	if !(course.Status == 1 || course.Status == 4) {
		return s.errHandler.PermissionDenied()
	}
	return nil
}
