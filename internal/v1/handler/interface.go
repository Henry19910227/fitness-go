package handler

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
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
	GenerateNewImageName(original string) (string, error)
	GenerateNewVideoName(original string) (string, error)
	UploadCourseCover(file io.Reader, imageNamed string) (string, error)
	UploadActionCover(file io.Reader, imageNamed string) error
	UploadTrainerAvatar(file io.Reader, imageNamed string) error
	UploadUserAvatar(file io.Reader, imageNamed string) (string, error)
	UploadWorkoutStartAudio(file io.Reader, audioNamed string) (string, error)
	UploadWorkoutEndAudio(file io.Reader, audioNamed string) (string, error)
	UploadWorkoutSetStartAudio(file io.Reader, audioNamed string) (string, error)
	UploadWorkoutSetProgressAudio(file io.Reader, audioNamed string) (string, error)
	UploadActionVideo(file io.Reader, videoNamed string) error
	UploadCardFrontImage(file io.Reader, imageNamed string) error
	UploadCardBackImage(file io.Reader, imageNamed string) error
	UploadTrainerAlbumPhoto(file io.Reader, imageNamed string) error
	UploadCertificateImage(file io.Reader, imageNamed string) error
	UploadAccountImage(file io.Reader, imageNamed string) error
	UploadReviewImage(file io.Reader, imageNamed string) error
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
	DeleteCardBackImage(imageNamed string) error
	DeleteTrainerAlbumPhoto(imageNamed string) error
	DeleteActionCover(coverNamed string) error
	DeleteActionVideo(videoNamed string) error
	DeleteCertificateImage(imageNamed string) error
	DeleteReviewImage(imageNamed string) error
}

type IAP interface {
	SandboxURL() string
	ProductURL() string
	Password() string
	ParserIAPNotificationType(notificationType string, subtype string) global.SubscribeLogType
	ParserAppleReceipt(dict map[string]interface{}, receipt *dto.IAPVerifyReceiptResponse) error
	GetAppleStoreAPIAccessToken() (string, error)
	VerifyAppleReceiptAPI(receiptData string) (*dto.IAPVerifyReceiptResponse, error)
	GetSubscribeAPI(originalTransactionId string) (*dto.IAPSubscribeAPIResponse, error)
	GetHistoryAPI(originalTransactionId string) (*dto.IAPHistoryAPIResponse, error)
}

type IAB interface {
	ParserIABNotificationType(notificationType int) global.SubscribeLogType
	ParserIABNotificationMsg(notificationType int) string
	GetGooglePlayApiAccessToken() (string, error)
	GetProductsAPI(productID string, purchaseToken string) (*dto.IABProductAPIResponse, error)
	GetSubscriptionAPI(productID string, purchaseToken string) (*dto.IABSubscriptionAPIResponse, error)
}
