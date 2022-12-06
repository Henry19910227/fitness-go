package plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/plan/api_update_trainer_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/plan/api_update_user_plan"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSPlans(input *model.APIGetCMSPlansInput) interface{}

	APICreateUserPlan(tx *gorm.DB, input *model.APICreateUserPlanInput) (output model.APICreateUserPlanOutput)
	APIDeleteUserPlan(tx *gorm.DB, input *model.APIDeleteUserPlanInput) (output model.APIDeleteUserPlanOutput)
	APIGetUserPlans(input *model.APIGetUserPlansInput) (output model.APIGetUserPlansOutput)
	APIUpdateUserPlan(input *api_update_user_plan.Input) (output api_update_user_plan.Output)

	APICreateTrainerPlan(tx *gorm.DB, input *model.APICreateTrainerPlanInput) (output model.APICreateTrainerPlanOutput)
	APIGetTrainerPlans(input *model.APIGetTrainerPlansInput) (output model.APIGetTrainerPlansOutput)
	APIDeleteTrainerPlan(tx *gorm.DB, input *model.APIDeleteTrainerPlanInput) (output model.APIDeleteTrainerPlanOutput)
	APIUpdateTrainerPlan(input *api_update_trainer_plan.Input) (output api_update_trainer_plan.Output)

	APIGetStorePlans(input *model.APIGetStorePlansInput) (output model.APIGetStorePlansOutput)
}
