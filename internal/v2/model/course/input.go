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
	PagingInput
	PreloadInput
	OrderByInput
	IDs []int64
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
