package user_subscribe_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_subscribe_monthly_statistics"
}

// APIGetCMSUserSubscribeStatisticOutput /v2/cms/statistic_monthly/user/subscribe [GET]
type APIGetCMSUserSubscribeStatisticOutput struct {
	base.Output
	Data *APIGetCMSUserSubscribeStatisticData `json:"data,omitempty"`
}
type APIGetCMSUserSubscribeStatisticData struct {
	Table
}
