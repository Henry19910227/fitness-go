package dto

type Diet struct {
	ID         int64   `json:"id" example:"1"` // rda
	RDA        *RDA    `json:"rda,omitempty"`  // rda
	ScheduleAt string  `json:"schedule_at" example:"2022-05-25 11:00:00"`
	CreateAt   *string `json:"create_at,omitempty" example:"2022-05-25 11:00:00"`
	UpdateAt   *string `json:"update_at,omitempty" example:"2022-05-25 11:00:00"`
	Meals      []*Meal `json:"meals,omitempty"`
}
