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
	UploadCourseCover(file io.Reader, imageNamed string) (string, error)
	UploadActionCover(file io.Reader, imageNamed string) (string, error)
	UploadTrainerAvatar(file io.Reader, imageNamed string) (string, error)
	UploadUserAvatar(file io.Reader, imageNamed string) (string, error)
	UploadWorkoutStartAudio(file io.Reader, audioNamed string) (string, error)
	UploadWorkoutEndAudio(file io.Reader, audioNamed string) (string, error)
	UploadWorkoutSetStartAudio(file io.Reader, audioNamed string) (string, error)
	UploadWorkoutSetProgressAudio(file io.Reader, audioNamed string) (string, error)
	UploadActionVideo(file io.Reader, videoNamed string) (string, error)
	UploadCardFrontImage(file io.Reader, imageNamed string) (string, error)
}

type Resource interface {
	DeleteCourseCover(imageNamed string) error
	DeleteTrainerAvatar(imageNamed string) error
	DeleteUserAvatar(imageNamed string) error
	DeleteWorkoutSetStartAudio(audioNamed string) error
	DeleteWorkoutSetProgressAudio(audioNamed string) error
	DeleteWorkoutStartAudio(audioNamed string) error
	DeleteWorkoutEndAudio(audioNamed string) error
	DeleteCardFrontImage(imageNamed string) error
}