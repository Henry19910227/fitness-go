package logger

import (
	"github.com/gin-gonic/gin"
)

type Tool interface {
	Trace(ctx *gin.Context, msg string)
	Debug(ctx *gin.Context, msg string)
	Info(ctx *gin.Context, msg string)
	Warn(ctx *gin.Context, msg string)
	Error(ctx *gin.Context, msg string)
	Fatal(ctx *gin.Context, msg string)
	Panic(ctx *gin.Context, msg string)
}
