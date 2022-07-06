package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	PlatformOptional
	PagingInput
	PreloadInput
	OrderByInput
}

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

// APIGetCMSFeedbacksInput /v2/cms/feedbacks [GET] 獲取反饋列表
type APIGetCMSFeedbacksInput struct {
	Form APIGetCMSFeedbacksForm
}
type APIGetCMSFeedbacksForm struct {
	PlatformOptional
	OrderByInput
	PagingInput
}
