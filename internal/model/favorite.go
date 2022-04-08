package model

type FavoriteCourse struct {
	UserID   int64  `gorm:"column:user_id"`
	CourseID int64  `gorm:"column:course_id"`
	CreateAt string `gorm:"column:create_at"`
}

func (FavoriteCourse) TableName() string {
	return "favorite_courses"
}

type FavoriteTrainer struct {
	UserID    int64  `gorm:"column:user_id"`
	TrainerID int64  `gorm:"column:trainer_id"`
	CreateAt  string `gorm:"column:create_at"`
}

func (FavoriteTrainer) TableName() string {
	return "favorite_trainers"
}
