package api_get_trainer_course_actions

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type PagingInput = paging.Input

// Input /v2/trainer/course/{course_id}/actions [GET] 獲取教練課表動作庫 API
type Input struct {
	userRequired.UserIDField
	Uri   Uri
	Query Query
}
type Uri struct {
	required.CourseIDField
}
type Query struct {
	Name      *string `json:"name,omitempty" form:"name" binding:"omitempty,min=1,max=20" example:"槓鈴臥推"`                 //動作名稱(1~20字元)
	Source    *string `json:"source,omitempty" form:"source" binding:"omitempty,action_source" example:"1,2"`             //動作來源(1:平台動作/2:教練)
	Category  *string `json:"category,omitempty" form:"category" binding:"omitempty,action_category" example:"1,2,3,4,5"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body      *string `json:"body,omitempty" form:"body" binding:"omitempty,action_body" example:"2,4,6"`                 //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *string `json:"equipment,omitempty" form:"equipment" binding:"omitempty,action_equipment" example:"1,3,5"`  //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	PagingInput
}
