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

type ActionBestPR struct {
	ActionID    int64   `gorm:"column:action_id"`    //動作id
	MaxRM       float64 `gorm:"column:max_rm"`       //最大反覆重量(公斤)
	MaxWeight   float64 `gorm:"column:max_weight"`   //最大重量(公斤)
	MaxReps     int     `gorm:"column:max_reps"`     //最多次數
	MinDuration int     `gorm:"column:min_duration"` //最小時長(秒)
	MaxSpeed    float64 `gorm:"column:max_speed"`    //最高速率
}
