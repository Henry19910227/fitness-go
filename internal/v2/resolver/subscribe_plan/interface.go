package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan/api_get_subscribe_plans"
)

type Resolver interface {
	APIGetSubscribePlans(input *api_get_subscribe_plans.Input) (output api_get_subscribe_plans.Output)
}
