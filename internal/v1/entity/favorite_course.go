package entity

type FavoriteCourse struct {
	UserID   int64   `gorm:"column:user_id"`
	CourseID int64   `gorm:"column:course_id"`
	CreateAt string  `gorm:"column:create_at"`
}

func (FavoriteCourse) TableName() string {
	return "favorite_courses"
}
