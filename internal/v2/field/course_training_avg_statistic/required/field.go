package required

type CourseIDField struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
type RateField struct {
	Rate int `json:"rate" gorm:"column:rate" binding:"required" example:"100"` //平均訓練率
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
