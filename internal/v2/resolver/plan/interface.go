package plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSPlans(input *model.APIGetCMSPlansInput) interface{}
	APICreatePersonalPlan(tx *gorm.DB, input *model.APICreatePersonalPlanInput) (output model.APICreatePersonalPlanOutput)
	APIDeletePersonalPlan(tx *gorm.DB, input *model.APIDeletePersonalPlanInput) (output model.APIDeletePersonalPlanOutput)
}
