package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
)

// APICreateFeedbackInput /v2/feedback [POST] 新增反饋
type APICreateFeedbackInput struct {
	UserIDRequired
	Form  APICreateFeedbackForm
	Files []*file.Input
}
type APICreateFeedbackForm struct {
	VersionOptional
	PlatformOptional
	OSVersionOptional
	PhoneModelOptional
	BodyRequired
}
