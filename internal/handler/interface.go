package handler

import (
	"github.com/gin-gonic/gin"
	"time"
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
	GenerateTrainerToken(uid int64) (string, error)
	GenerateAdminToken(uid int64, lv int) (string, error)
	VerifyUserToken(token string) error
    VerifyTrainerToken(token string) error
	VerifyLV1AdminToken(token string) error
	VerifyLV2AdminToken(token string) error
	ResignAdminToken(token string) error
	ResignAdminTokenWithUID(uid int64) error
	ResignUserToken(token string) error
	ResignUserTokenWithUID(uid int64) error

	RenewOnlineStatus(token string) error
	SetOfflineStatus(token string) error
	SetOfflineStatusWithUID(uid int64, role int) error
	GetOnlineDateTime(uid int64) (*time.Time, error)
}

type Logger interface {
	Set(c *gin.Context, level LogLevel, tag string, code int, msg string)
}