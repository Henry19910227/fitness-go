package optional

type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" bind:"omitempty" example:"1"` //課表id
}
type ScoreTotalField struct {
	ScoreTotal *int `json:"score_total,omitempty" gorm:"column:score_total" bind:"omitempty" example:"100"` //評分累積
}
type AmountField struct {
	Amount *int `json:"amount,omitempty" gorm:"column:amount" bind:"omitempty" example:"10"` //評分筆數
}
type FiveTotalField struct {
	FiveTotal *int `json:"five_total,omitempty" gorm:"column:five_total" bind:"omitempty" example:"20"` //五分總筆數
}
type FourTotalField struct {
	FourTotal *int `json:"four_total,omitempty" gorm:"column:four_total" bind:"omitempty" example:"15"` //四分總筆數
}
type ThreeTotalField struct {
	ThreeTotal *int `json:"three_total,omitempty" gorm:"column:three_total" bind:"omitempty" example:"10"` //三分總筆數
}
type TwoTotalField struct {
	TwoTotal *int `json:"two_total,omitempty" gorm:"column:two_total" bind:"omitempty" example:"5"` //二分總筆數
}
type OneTotalField struct {
	OneTotal *int `json:"one_total,omitempty" gorm:"column:one_total" bind:"omitempty" example:"5"` //一分總筆數
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" bind:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
