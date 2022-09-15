package user_register_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_register_monthly_statistics"
}

// APIGetCMSUserRegisterStatisticOutput /v2/cms/statistic_monthly/user/register [GET]
type APIGetCMSUserRegisterStatisticOutput struct {
	base.Output
	Data *APIGetCMSUserRegisterStatisticData `json:"data,omitempty"`
}
type APIGetCMSUserRegisterStatisticData struct {
	Table
}
