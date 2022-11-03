package entity

type TrainerAlbum struct {
	ID       int64  `gorm:"column:id"`        // 教練相片id
	UserID   int64  `gorm:"column:user_id"`   // 關聯的用戶id
	Photo    string `gorm:"column:photo"`     // 照片
	CreateAt string `gorm:"column:create_at"` // 創建日期
}

func (TrainerAlbum) TableName() string {
	return "trainer_albums"
}
