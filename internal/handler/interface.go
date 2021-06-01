package handler

import (
	"github.com/gin-gonic/gin"
	"io"
)

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Panic
)

type SSO interface {
	GenerateUserToken(uid int64) (string, error)
	GenerateAdminToken(uid int64, lv int) (string, error)
	VerifyUserToken(token string) error
	VerifyLV1AdminToken(token string) error
	VerifyLV2AdminToken(token string) error
	ResignAdminToken(token string) error
	ResignAdminTokenWithUID(uid int64) error
	ResignUserToken(token string) error
	ResignUserTokenWithUID(uid int64) error
}

type Logger interface {
	Set(c *gin.Context, level LogLevel, tag string, code int, msg string)
}

type Uploader interface {
	UploadCourseCover(file io.Reader, imageNamed string, courseID int64) (string, error)
	UploadTrainerAvatar(file io.Reader, imageNamed string, uid int64) (string, error)
}