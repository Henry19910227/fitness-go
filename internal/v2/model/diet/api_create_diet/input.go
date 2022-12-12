package api_create_diet

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	ScheduleAt string `json:"schedule_at" binding:"required,datetime=2006-01-02" example:"2022-05-25"`
}
