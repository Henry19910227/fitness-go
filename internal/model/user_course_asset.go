package model

type UserCourseAsset struct {
	ID        int64  `gorm:"column:id"`        // id
	UserID    int64  `gorm:"column:user_id"`   // 用戶id
	CourseID  int64  `gorm:"column:course_id"` // 課表id
	Available int    `gorm:"column:available"` // 是否可用(0:不可用/1:可用)
	CreateAt  string `gorm:"create_at"`        // 創建時間
	UpdateAt  string `gorm:"update_at"`        // 更新時間
}

func (UserCourseAsset) TableName() string {
	return "user_course_assets"
}

type FindUserCourseAssetParam struct {
	UserID   int64 // 用戶id
	CourseID int64 // 課表id
}

type CreateUserCourseAssetParam struct {
	UserID   int64 // 用戶id
	CourseID int64 // 課表id
}
