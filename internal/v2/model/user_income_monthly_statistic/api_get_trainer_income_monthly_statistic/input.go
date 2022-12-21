package api_get_trainer_income_monthly_statistic

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/trainer/income_monthly_statistic [GET]
type Input struct {
	userRequired.UserIDField
}
