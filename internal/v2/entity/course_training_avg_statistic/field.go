package course_training_avg_statistic

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //報表id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}
type RateField struct {
	Rate *int `json:"rate,omitempty" gorm:"column:rate" example:"100"` //平均訓練率
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	CourseIDField
	RateField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "course_training_avg_statistics"
}