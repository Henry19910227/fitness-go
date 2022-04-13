package entity

type ReviewStatistic struct {
	CourseID   int64  `gorm:"column:course_id"`   //課表id
	ScoreTotal int    `gorm:"column:score_total"` //評分累積
	Amount     int    `gorm:"column:amount"`      //評分筆數
	FiveTotal  int    `gorm:"column:five_total"`  //五分總筆數
	FourTotal  int    `gorm:"column:four_total"`  //四分總筆數
	ThreeTotal int    `gorm:"column:three_total"` //三分總筆數
	TwoTotal   int    `gorm:"column:two_total"`   //二分總筆數
	OneTotal   int    `gorm:"column:one_total"`   //一分總筆數
	UpdateAt   string `gorm:"column:update_at"`   //更新時間
}

func (ReviewStatistic) TableName() string {
	return "review_statistics"
}
