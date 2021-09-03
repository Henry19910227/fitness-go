package middleware

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/gin-gonic/gin"
)

type User interface {
	TokenPermission(roles []global.Role) gin.HandlerFunc
	UserStatusPermission(status []global.UserStatus) gin.HandlerFunc
	TrainerStatusPermission(status []global.TrainerStatus) gin.HandlerFunc
}

type Course interface {
	WorkoutSetPermission(status []global.CourseStatus) gin.HandlerFunc
	CourseCreatorVerify() gin.HandlerFunc
	CourseStatusAccessRange(status []global.CourseStatus, ext []global.CourseStatus) gin.HandlerFunc
}