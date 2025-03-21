package model

type Plan struct {
	ID           int64  `gorm:"column:id"`            //計畫id
	CourseID     int64  `gorm:"column:course_id"`     //課表id
	Name         string `gorm:"column:name"`          //計畫名稱
	WorkoutCount int    `gorm:"column:workout_count"` //訓練數量
	CreateAt     string `gorm:"column:create_at"`     //創建時間
	UpdateAt     string `gorm:"column:update_at"`     //更新時間
}

func (Plan) TableName() string {
	return "plans"
}

type PlanAsset struct {
	ID                 int64  `gorm:"column:id"`                   //計畫id
	CourseID           int64  `gorm:"column:course_id"`            //課表id
	Name               string `gorm:"column:name"`                 //計畫名稱
	WorkoutCount       int    `gorm:"column:workout_count"`        //訓練數量
	FinishWorkoutCount int    `gorm:"column:finish_workout_count"` //完成訓練數量
	CreateAt           string `gorm:"column:create_at"`            //創建時間
	UpdateAt           string `gorm:"column:update_at"`            //更新時間
}

func (PlanAsset) TableName() string {
	return "plans"
}
