package model

type ActionPR struct {
	ID       int64   `gorm:"column:id"`        //PR紀錄 id
	ActionID int64   `gorm:"column:action_id"` //動作id
	Weight   float64 `gorm:"column:weight"`    //重量(公斤)
	Reps     int     `gorm:"column:reps"`      //次數
	Distance float64 `gorm:"column:distance"`  //距離(公里)
	Duration int     `gorm:"column:duration"`  //時長(秒)
	Incline  float64 `gorm:"column:incline"`   //坡度
}

type CreateActionPRParam struct {
	ActionID int64   `gorm:"column:action_id"` //動作id
	Weight   float64 `gorm:"column:weight"`    //重量(公斤)
	Reps     int     `gorm:"column:reps"`      //次數
	Distance float64 `gorm:"column:distance"`  //距離(公里)
	Duration int     `gorm:"column:duration"`  //時長(秒)
	Incline  float64 `gorm:"column:incline"`   //坡度
}
