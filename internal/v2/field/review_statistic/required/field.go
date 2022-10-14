package required

type CourseIDField struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" bind:"required" example:"1"` //課表id
}
type ScoreTotalField struct {
	ScoreTotal int `json:"score_total" gorm:"column:score_total" bind:"required" example:"100"` //評分累積
}
type AmountField struct {
	Amount int `json:"amount" gorm:"column:amount" bind:"required" example:"10"` //評分筆數
}
type FiveTotalField struct {
	FiveTotal int `json:"five_total" gorm:"column:five_total" bind:"required" example:"20"` //五分總筆數
}
type FourTotalField struct {
	FourTotal int `json:"four_total" gorm:"column:four_total" bind:"required" example:"15"` //四分總筆數
}
type ThreeTotalField struct {
	ThreeTotal int `json:"three_total" gorm:"column:three_total" bind:"required" example:"10"` //三分總筆數
}
type TwoTotalField struct {
	TwoTotal int `json:"two_total" gorm:"column:two_total" bind:"required" example:"5"` //二分總筆數
}
type OneTotalField struct {
	OneTotal int `json:"one_total" gorm:"column:one_total" bind:"required" example:"5"` //一分總筆數
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" bind:"required" example:"2022-06-14 00:00:00"` //更新時間
}
