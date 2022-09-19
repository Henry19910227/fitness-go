package course_training_avg_statistic

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	avgOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course_training_avg_statistic/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

// APIGetCMSCourseTrainingAvgStatisticInput /v2/cms/statistic_monthly/course/training_avg [GET]
type APIGetCMSCourseTrainingAvgStatisticInput struct {
	Query APIGetCMSCourseReleaseStatisticQuery
}
type APIGetCMSCourseReleaseStatisticQuery struct {
	avgOptional.CourseIDField
	courseOptional.NameField
	courseOptional.CourseStatusField
	courseOptional.SaleTypeField
	PagingInput
	PreloadInput
}
