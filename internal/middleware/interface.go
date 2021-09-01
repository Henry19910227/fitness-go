package middleware

import "github.com/gin-gonic/gin"

type User interface {
	TokenPermission(roles []Role) gin.HandlerFunc
	UserStatusPermission(status []UserStatus) gin.HandlerFunc
	TrainerStatusPermission(status []TrainerStatus) gin.HandlerFunc
}

type Course interface {
	WorkoutSetPermission(status []CourseStatus) gin.HandlerFunc
	CoursePermission(status []CourseStatus) gin.HandlerFunc
}