package api_get_cms_statistic_monthly_user_promote

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_promote_monthly_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/statistic_monthly/user/promote [GET]
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
	optional.Exp1to3Field
	optional.Exp4to6Field
	optional.Exp7to10Field
	optional.Exp11to15Field
	optional.Exp16to19Field
	optional.Exp20upField
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
