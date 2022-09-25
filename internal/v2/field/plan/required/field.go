package required

type IDField struct {
	ID int64 `json:"id" uri:"plan_id" gorm:"column:id" binding:"required" example:"1"` //計畫id
}

type CourseIDField struct {
	CourseID int64 `json:"course_id" uri:"course_id" gorm:"column:course_id" binding:"required" example:"1"` //課表id
}

type NameField struct {
	Name string `json:"name" gorm:"column:name" binding:"required,min=1,max=20" example:"第一週增肌計畫"` //計畫名稱
}

type WorkoutCountField struct {
	WorkoutCount int `json:"workout_count" gorm:"column:workout_count" binding:"required" example:"10"` //訓練數量
}

type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}

type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
