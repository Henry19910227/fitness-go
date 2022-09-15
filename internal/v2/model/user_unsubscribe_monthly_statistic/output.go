package user_unsubscribe_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_unsubscribe_monthly_statistics"
}

// APIGetCMSUserUnsubscribeStatisticOutput /v2/cms/statistic_monthly/user/unsubscribe [GET]
type APIGetCMSUserUnsubscribeStatisticOutput struct {
	base.Output
	Data *APIGetCMSUserUnsubscribeStatisticData `json:"data,omitempty"`
}
type APIGetCMSUserUnsubscribeStatisticData struct {
	Table
}
