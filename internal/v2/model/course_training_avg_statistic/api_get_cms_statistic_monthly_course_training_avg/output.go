package api_get_cms_statistic_monthly_course_training_avg

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	avgRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course_training_avg_statistic/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/statistic_monthly/course/training_avg [GET]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
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
		avgRequired.RateField
	} `json:"course_training_avg_statistic,omitempty"`
}
