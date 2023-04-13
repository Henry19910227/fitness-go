package optional

type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty"  form:"course_id" gorm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
type RateField struct {
	Rate *int `json:"rate,omitempty" gorm:"column:rate" binding:"omitempty" example:"100"` //平均訓練率
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
