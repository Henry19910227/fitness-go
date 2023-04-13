package api_get_cms_statistic_monthly_course_training_avg

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course_training_avg_statistic/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

// Input /v2/cms/statistic_monthly/course/training_avg [GET]
type Input struct {
	Query Query
}
type Query struct {
	optional.CourseIDField
	courseOptional.NameField
	courseOptional.CourseStatusField
	courseOptional.SaleTypeField
	PagingInput
}
