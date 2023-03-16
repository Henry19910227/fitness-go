package api_get_cms_user_subscribe_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_monthly_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/statistic_monthly/user/subscribe [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.YearField
	optional.MonthField
	optional.TotalField
	optional.MaleField
	optional.FemaleField
	optional.Age13to17Field
	optional.Age18to24Field
	optional.Age25to34Field
	optional.Age35to44Field
	optional.Age45to54Field
	optional.Age55to64Field
	optional.Age65UpField
	optional.CreateAtField
	optional.UpdateAtField
}
