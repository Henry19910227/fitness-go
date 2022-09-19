package course_training_avg_statistic

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	avgOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course_training_avg_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
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
	courseOptional.IDField
	courseOptional.NameField
	courseOptional.CourseStatusField
	courseOptional.ScheduleTypeField
	courseOptional.SaleTypeField
	SaleItem *struct {
		saleItemOptional.IDField
		saleItemOptional.NameField
		ProductLabel *struct {
			optional.IDField
			optional.ProductIDField
			optional.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
	CourseTrainingAvgStatistic struct {
		avgOptional.RateField
	} `json:"course_training_avg_statistic,omitempty"`
}
