package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type GenerateInput struct {
	DataAmount int
	UserID     []*base.GenerateSetting
}

type FindInput struct {
	IDOptional
	PreloadInput
}

type ListInput struct {
	IDOptional
	NameOptional
	CourseStatusOptional
	SaleTypeOptional
	IgnoredCourseStatus []int // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	IDs []int64
	PagingInput
	PreloadInput
	OrderByInput
}

type APIGetCMSCoursesInput struct {
	IDOptional
	NameOptional
	CourseStatusOptional
	SaleTypeOptional
	PagingInput
	OrderByInput
}

type APIGetCMSCourseInput struct {
	IDRequired
}

type APIUpdateCMSCoursesStatusInput struct {
	IDs []int64 `json:"ids" binding:"required"` // 課表 id
	CourseStatusRequired
}

type APIUpdateCMSCourseCoverInput struct {
	IDRequired
	CoverNamed string
	File       multipart.File
}
