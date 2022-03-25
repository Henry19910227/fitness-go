package dto

type WorkoutSetLog struct {
	WorkoutSet WorkoutSet `json:"workout_set"`           //訓練組
	Weight     float64    `json:"weight" example:"10"`   //重量(公斤)
	Reps       int        `json:"reps" example:"5"`      //次數
	Distance   float64    `json:"distance" example:"1"`  //距離(公里)
	Duration   int        `json:"duration" example:"30"` //時長(秒)
	Incline    float64    `json:"incline" example:"5"`   //坡度
}

type WorkoutSetLogParam struct {
	WorkoutSetID int64   `json:"workout_set_id" example:"1"` //訓練組id
	Weight       float64 `json:"weight" example:"10"`        //重量(公斤)
	Reps         int     `json:"reps" example:"5"`           //次數
	Distance     float64 `json:"distance" example:"1"`       //距離(公里)
	Duration     int     `json:"duration" example:"30"`      //時長(秒)
	Incline      float64 `json:"incline" example:"5"`        //坡度
}
