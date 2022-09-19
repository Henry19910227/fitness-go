package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course"
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course_training_avg_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	productLabel "github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
	saleItem "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
)

type Output struct {
	course_training_avg_statistic.Table
}

func (Output) TableName() string {
	return "course_training_avg_statistics"
}

// APIGetCMSCourseTrainingAvgStatisticOutput /v2/cms/statistic_monthly/course/training_avg [GET]
type APIGetCMSCourseTrainingAvgStatisticOutput struct {
	base.Output
	Data   *APIGetCMSCourseTrainingAvgStatisticData `json:"data,omitempty"`
	Paging *paging.Output                           `json:"paging,omitempty"`
}
type APIGetCMSCourseTrainingAvgStatisticData []*struct {
	course.IDField
	course.NameField
	course.CourseStatusField
	course.ScheduleTypeField
	course.SaleTypeField
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
	CourseTrainingAvgStatistic struct {
		course_training_avg_statistic.RateRequired
	} `json:"course_training_avg_statistic,omitempty"`
}
