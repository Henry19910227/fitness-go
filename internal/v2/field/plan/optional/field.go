package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" uri:"plan_id" gorm:"column:id" binding:"omitempty" example:"1"` //計畫id
}

type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" uri:"course_id" gorm:"column:course_id" binding:"omitempty" example:"1"` //課表id
}

type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" binding:"omitempty" example:"第一週增肌計畫"` //計畫名稱
}

type WorkoutCountField struct {
	WorkoutCount *int `json:"workout_count,omitempty" gorm:"column:workout_count" binding:"omitempty" example:"10"` //訓練數量
}

type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}

type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
