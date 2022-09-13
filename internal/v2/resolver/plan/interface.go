package plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSPlans(input *model.APIGetCMSPlansInput) interface{}
	APICreateUserPlan(tx *gorm.DB, input *model.APICreateUserPlanInput) (output model.APICreateUserPlanOutput)
	APIDeleteUserPlan(tx *gorm.DB, input *model.APIDeleteUserPlanInput) (output model.APIDeleteUserPlanOutput)
	APIGetUserPlans(input *model.APIGetUserPlansInput) (output model.APIGetUserPlansOutput)
	APIUpdateUserPlan(input *model.APIUpdateUserPlanInput) (output model.APIUpdateUserPlanOutput)

	APICreateTrainerPlan(tx *gorm.DB, input *model.APICreateTrainerPlanInput) (output model.APICreateTrainerPlanOutput)
}
