package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` // 統計 id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" orm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
type TotalFinishWorkoutCountField struct {
	TotalFinishWorkoutCount *int `json:"total_finish_workout_count,omitempty" gorm:"column:total_finish_workout_count" binding:"omitempty" example:"5500"` // 歷史訓練總量(完成一次該課表的訓練，重複計算)
}
type UserFinishCountField struct {
	UserFinishCount *int `json:"user_finish_count,omitempty" gorm:"column:user_finish_count" binding:"omitempty" example:"3000"` // 用戶使用人數(完成一次該課表的訓練，不重複計算)
}
type FinishCountAvgField struct {
	FinishCountAvg *int `json:"finish_count_avg,omitempty" gorm:"column:finish_count_avg" binding:"omitempty" example:"25"` // 平均完成訓練次數(同一訓練不重複)
}
type MaleFinishCountField struct {
	MaleFinishCount *int `json:"male_finish_count,omitempty" gorm:"column:male_finish_count" binding:"omitempty" example:"1000"` // 男生使用人數
}
type FemaleFinishCountField struct {
	FemaleFinishCount *int `json:"female_finish_count,omitempty" gorm:"column:female_finish_count" binding:"omitempty" example:"1000"` // 女生使用人數
}
type Age13to17CountField struct {
	Age13to17Count *int `json:"age_13_17_count,omitempty" gorm:"column:age_13_17_count" binding:"omitempty" example:"500"` // 13-17歲使用人數
}
type Age18to24CountField struct {
	Age18to24Count *int `json:"age_18_24_count,omitempty" gorm:"column:age_18_24_count" binding:"omitempty" example:"100"` // 18-24歲使用人數
}
type Age25to34CountField struct {
	Age25to34Count *int `json:"age_25_34_count,omitempty" gorm:"column:age_25_34_count" binding:"omitempty" example:"300"` // 25-34歲使用人數
}
type Age35to44CountField struct {
	Age35to44Count *int `json:"age_35_44_count,omitempty" gorm:"column:age_35_44_count" binding:"omitempty" example:"200"` // 35-44歲使用人數
}
type Age45to54CountField struct {
	Age45to54Count *int `json:"age_45_54_count,omitempty" gorm:"column:age_45_54_count" binding:"omitempty" example:"600"` // 45-54歲使用人數
}
type Age55to64CountField struct {
	Age55to64Count *int `json:"age_55_64_count,omitempty" gorm:"column:age_55_64_count" binding:"omitempty" example:"700"` // 55-64歲使用人數
}
type Age65upCountField struct {
	Age65upCount *int `json:"age_65_up_count,omitempty" gorm:"column:age_65_up_count" binding:"omitempty" example:"550"` // 65+歲使用人數
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
