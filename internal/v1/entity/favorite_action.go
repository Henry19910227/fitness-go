package entity

type FavoriteAction struct {
	UserID   int64  `gorm:"column:user_id"`
	ActionID int64  `gorm:"column:action_id"`
	CreateAt string `gorm:"column:create_at"`
}

func (FavoriteAction) TableName() string {
	return "favorite_actions"
}
