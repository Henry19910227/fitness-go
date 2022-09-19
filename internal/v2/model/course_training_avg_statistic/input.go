package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course"
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course_training_avg_statistic"
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
	course_training_avg_statistic.CourseIDOptional
	course.NameField
	course.CourseStatusField
	course.SaleTypeField
	PagingInput
	PreloadInput
}
