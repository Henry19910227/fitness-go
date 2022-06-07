package validator

type CreateDietBody struct {
	ScheduleAt string `json:"schedule_at" binding:"required,datetime=2006-01-02" example:"2022-05-25"`
}

type GetDietQuery struct {
	ScheduleAt string `form:"schedule_at" binding:"required,datetime=2006-01-02" example:"2022-05-25"`
}
