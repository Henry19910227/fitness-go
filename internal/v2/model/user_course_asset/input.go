package user_course_asset

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_course_assets"
}

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	UserIDOptional
	CourseIDOptional
	AvailableOptional
	PagingInput
	PreloadInput
	OrderByInput
}