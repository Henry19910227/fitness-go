package review_statistic

type CourseIDRequired struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" binding:"required" example:"1"` //課表id
}
type ScoreTotalRequired struct {
	ScoreTotal int `json:"score_total" gorm:"column:score_total" binding:"required" example:"100"` //評分累積
}
type AmountRequired struct {
	Amount int `json:"amount" gorm:"column:amount" binding:"required" example:"10"` //評分筆數
}
type FiveTotalRequired struct {
	FiveTotal int `json:"five_total" gorm:"column:five_total" binding:"required" example:"20"` //五分總筆數
}
type FourTotalRequired struct {
	FourTotal int `json:"four_total" gorm:"column:four_total" binding:"required" example:"15"` //四分總筆數
}
type ThreeTotalRequired struct {
	ThreeTotal int `json:"three_total" gorm:"column:three_total" binding:"required" example:"10"` //三分總筆數
}
type TwoTotalRequired struct {
	TwoTotal int `json:"two_total" gorm:"column:two_total" binding:"required" example:"5"` //二分總筆數
}
type OneTotalRequired struct {
	OneTotal int `json:"one_total" gorm:"column:one_total" binding:"required" example:"5"` //一分總筆數
}
type UpdateAtRequired struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
