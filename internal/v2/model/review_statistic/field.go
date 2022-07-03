package review_statistic

type CourseIDField struct {
	CourseID  *int64  `json:"course_id,omitempty" gorm:"column:course_id" example:"1"`   //課表id
}
type ScoreTotalField struct {
	ScoreTotal *int   `json:"score_total,omitempty" gorm:"column:score_total" example:"100"` //評分累積
}
type AmountField struct {
	Amount     *int    `json:"amount,omitempty" gorm:"column:amount" example:"10"`      //評分筆數
}
type FiveTotalField struct {
	FiveTotal  *int    `json:"five_total,omitempty" gorm:"column:five_total" example:"20"`  //五分總筆數
}
type FourTotalField struct {
	FourTotal  *int    `json:"four_total,omitempty" gorm:"column:four_total" example:"15"`  //四分總筆數
}
type ThreeTotalField struct {
	ThreeTotal  *int    `json:"three_total,omitempty" gorm:"column:three_total" example:"10"`  //三分總筆數
}
type TwoTotalField struct {
	TwoTotal  *int    `json:"two_total,omitempty" gorm:"column:two_total" example:"5"`  //二分總筆數
}
type OneTotalField struct {
	OneTotal  *int    `json:"one_total,omitempty" gorm:"column:one_total" example:"5"`  //一分總筆數
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	CourseIDField
	ScoreTotalField
	AmountField
	FiveTotalField
	FourTotalField
	ThreeTotalField
	TwoTotalField
	OneTotalField
	UpdateAtField
}
func (Table) TableName() string {
	return "review_statistics"
}