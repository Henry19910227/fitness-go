package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // 統計 id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" orm:"column:course_id" binding:"required" example:"10"` //課表id
}
type TotalFinishWorkoutCountField struct {
	TotalFinishWorkoutCount int `json:"total_finish_workout_count" gorm:"column:total_finish_workout_count" binding:"required" example:"5500"` // 歷史訓練總量(完成一次該課表的訓練，重複計算)
}
type UserFinishCountField struct {
	UserFinishCount int `json:"user_finish_count" gorm:"column:user_finish_count" binding:"required" example:"3000"` // 用戶使用人數(完成一次該課表的訓練，不重複計算)
}
type FinishCountAvgField struct {
	FinishCountAvg int `json:"finish_count_avg" gorm:"column:finish_count_avg" binding:"required" example:"25"` // 平均完成訓練次數(同一訓練不重複)
}
type MaleFinishCountField struct {
	MaleFinishCount int `json:"male_finish_count" gorm:"column:male_finish_count" binding:"required" example:"1000"` // 男生使用人數
}
type FemaleFinishCountField struct {
	FemaleFinishCount int `json:"female_finish_count" gorm:"column:female_finish_count" binding:"required" example:"1000"` // 女生使用人數
}
type Age13to17CountField struct {
	Age13to17Count int `json:"age_13_17_count" gorm:"column:age_13_17_count" binding:"required" example:"500"` // 13-17歲使用人數
}
type Age18to24CountField struct {
	Age18to24Count int `json:"age_18_24_count" gorm:"column:age_18_24_count" binding:"required" example:"100"` // 18-24歲使用人數
}
type Age25to34CountField struct {
	Age25to34Count int `json:"age_25_34_count" gorm:"column:age_25_34_count" binding:"required" example:"300"` // 25-34歲使用人數
}
type Age35to44CountField struct {
	Age35to44Count int `json:"age_35_44_count" gorm:"column:age_35_44_count" binding:"required" example:"200"` // 35-44歲使用人數
}
type Age45to54CountField struct {
	Age45to54Count int `json:"age_45_54_count" gorm:"column:age_45_54_count" binding:"required" example:"600"` // 45-54歲使用人數
}
type Age55to64CountField struct {
	Age55to64Count int `json:"age_55_64_count" gorm:"column:age_55_64_count" binding:"required" example:"700"` // 55-64歲使用人數
}
type Age65upCountField struct {
	Age65upCount int `json:"age_65_up_count" gorm:"column:age_65_up_count" binding:"required" example:"550"` // 65+歲使用人數
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
