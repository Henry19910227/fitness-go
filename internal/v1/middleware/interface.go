package middleware

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/gin-gonic/gin"
)

type User interface {
	TokenPermission(roles []global.Role) gin.HandlerFunc
	UserStatusPermission(status []global.UserStatus) gin.HandlerFunc
	TrainerStatusPermission(status []global.TrainerStatus) gin.HandlerFunc
	TrainerAlbumPhotoLimit(currentCount func(c *gin.Context, uid int64) (int, errcode.Error), createCount, deleteCount func(c *gin.Context) int, limitCount int) gin.HandlerFunc
	CertificateLimit(currentCount func(c *gin.Context, uid int64) (int, errcode.Error), createCount, deleteCount func(c *gin.Context) int, limitCount int) gin.HandlerFunc
}

type Course interface {
	WorkoutSetStatusAccessRange(status []global.CourseStatus, ext []global.CourseStatus) gin.HandlerFunc
	CourseCreatorVerify() gin.HandlerFunc
	UserRoleAccessCourseByStatusRange(status []global.CourseStatus) gin.HandlerFunc
	AdminAccessCourseByStatusRange(status []global.CourseStatus) gin.HandlerFunc
	CourseStatusVerify(currentStatus func(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc
}

type Plan interface {
	CourseStatusVerify(currentStatus func(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc
}

type Workout interface {
	CourseStatusVerify(currentStatus func(c *gin.Context, workoutID int64) (global.CourseStatus, errcode.Error), allowStatus []global.CourseStatus) gin.HandlerFunc
}

type Review interface {
	ReviewCreatorVerify(reviewOwner func(c *gin.Context, reviewID int64) (int64, errcode.Error)) gin.HandlerFunc
}