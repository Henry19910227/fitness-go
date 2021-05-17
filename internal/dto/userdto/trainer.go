package userdto

type CreateTrainerParam struct {
	Name string
	Nickname string
	Phone string
	Email string
}

type CreateTrainerResult struct {
	TrainerID int64
}