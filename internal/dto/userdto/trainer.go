package userdto

type SwitchTrainerModeResult struct {
	Trainer struct{
		Name string `json:"name" example:"王小明"`
		Nickname string `json:"nickname" example:"Henry"`
		Phone string `json:"phone" example:"0978820789"`
		Email string `json:"email" example:"henry@gmail.com"`
	}
	Token string `json:"token" example:""`
}

type CreateTrainerParam struct {
	Name string
	Nickname string
	Phone string
	Email string
}

type CreateTrainerResult struct {
	TrainerID int64 `json:"trainer_id" example:"1"`
}