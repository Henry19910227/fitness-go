package course_category_training_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_category_training_monthly_statistics"
}

// APIGetCMSCategoryTrainingStatisticOutput /v2/cms/statistic_monthly/course_category/training [GET]
type APIGetCMSCategoryTrainingStatisticOutput struct {
	base.Output
	Data *APIGetCMSCategoryTrainingStatisticData `json:"data,omitempty"`
}
type APIGetCMSCategoryTrainingStatisticData struct {
	Table
}
