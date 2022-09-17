package user_promote_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_promote_monthly_statistics"
}

// APIGetCMSUserPromoteStatisticOutput /v2/cms/statistic_monthly/user/promote [GET]
type APIGetCMSUserPromoteStatisticOutput struct {
	base.Output
	Data *APIGetCMSUserPromoteStatisticData `json:"data,omitempty"`
}
type APIGetCMSUserPromoteStatisticData struct {
	Table
}
