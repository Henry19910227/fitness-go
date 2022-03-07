package validator

import "github.com/Henry19910227/fitness-go/internal/dto"

type WorkoutIDUri struct {
	WorkoutID int64 `uri:"workout_id" binding:"required" example:"1"`
}

type CreateWorkoutBody struct {
	Name              string `json:"name" binding:"required,min=1,max=20" example:"胸肌訓練"`
	WorkoutTemplateID *int64 `json:"workout_template_id" binding:"omitempty" example:"1"` // 訓練模板ID
}

type CreateWorkoutLogBody struct {
	Duration       int                  `json:"duration" binding:"required" example:"3600"`                 // 訓練時長(秒)
	Intensity      int                  `json:"intensity" binding:"required,oneof=1 2 3 4 5 6" example:"4"` // 訓練強度(1:輕鬆/2:適中/3:稍難/4:很累)
	Place          int                  `json:"place" binding:"required,oneof=1 2 3 4 5" example:"1"`       // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	WorkoutSetLogs []*dto.WorkoutSetLog `json:"workout_set_logs"`                                           // 訓練組記錄
}

type UpdateWorkoutBody struct {
	Name      *string `json:"name" binding:"omitempty,min=1,max=20" example:"胸肌訓練"`
	Equipment *string `json:"equipment" binding:"omitempty,equipment,min=0,max=10" example:"2,3,7"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
