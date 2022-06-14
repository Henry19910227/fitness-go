package entity

type FavoriteTrainer struct {
	UserID    int64  `gorm:"column:user_id"`
	TrainerID int64  `gorm:"column:trainer_id"`
	CreateAt  string `gorm:"column:create_at"`
}

func (FavoriteTrainer) TableName() string {
	return "favorite_trainers"
}
