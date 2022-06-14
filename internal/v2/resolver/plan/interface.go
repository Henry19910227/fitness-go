package plan

import model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"

type Resolver interface {
	APIGetCMSPlans(input *model.APIGetCMSPlansInput) interface{}
}
