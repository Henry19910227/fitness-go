package required

type IDField struct {
	ID string `json:"id" gorm:"column:id" binding:"required" example:"1"` // 訂單id
}
type VersionField struct {
	Version string `json:"version" gorm:"column:version" binding:"required" example:"1.0.0"` // android 版本號
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` // 創建時間
}
