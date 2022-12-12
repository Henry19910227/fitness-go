package api_get_diet

import (
	dietRequired "github.com/Henry19910227/fitness-go/internal/v2/field/diet/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

type Input struct {
	userRequired.UserIDField
	Query Query
}

type Query struct {
	dietRequired.ScheduleAtField
}
