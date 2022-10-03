package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
}

type ListInput struct {
	IDs           []int64 `json:"ids"`            //動作id
	SourceList    []int   `json:"source_list"`    //動作來源(1:平台動作/2:教練動作)
	CategoryList  []int   `json:"category_list"`  //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	BodyList      []int   `json:"body_list"`      //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	EquipmentList []int   `json:"equipment_list"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	UserIDOptional
	NameOptional
	TypeOptional
	SourceOptional
	PagingInput
	OrderByInput
}

type DeleteInput struct {
	IDRequired
}

type UserActionListInput struct {
	UserIDOptional
	NameOptional
	Source    []int `json:"source"`    //動作來源(1:平台動作/2:教練動作)
	Category  []int `json:"category"`  //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body      []int `json:"body"`      //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment []int `json:"equipment"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	PagingInput
	OrderByInput
}

type APIGetCMSActionsInput struct {
	PagingInput
}

// APICreateCMSActionInput /cms/action [POST] 創建動作 API
type APICreateCMSActionInput struct {
	Form      APICreateCMSActionForm
	CoverFile *file.Input
	VideoFile *file.Input
}
type APICreateCMSActionForm struct {
	NameRequired
	TypeRequired
	CategoryRequired
	BodyRequired
	EquipmentRequired
	IntroRequired
}

// APIUpdateCMSActionInput /v2/cms/action/{action_id} [PATCH] 更新動作 API
type APIUpdateCMSActionInput struct {
	Uri       APIUpdateCMSActionUri
	Form      APIUpdateCMSActionForm
	CoverFile *file.Input
	VideoFile *file.Input
}
type APIUpdateCMSActionForm struct {
	NameOptional
	IntroOptional
	StatusOptional
}
type APIUpdateCMSActionUri struct {
	IDRequired
}

// APICreateUserActionInput /v2/user/action [POST] 新增個人動作 API
type APICreateUserActionInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Cover  *file.Input
	Video  *file.Input
	Form   APICreateUserActionForm
}
type APICreateUserActionForm struct {
	NameRequired
	TypeRequired
	CategoryRequired
	BodyRequired
	EquipmentRequired
	IntroRequired
}

// APIUpdateUserActionInput /v2/user/action/{action_id} [PATCH] 修改個人動作 API
type APIUpdateUserActionInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Cover  *file.Input
	Video  *file.Input
	Form   APIUpdateUserActionForm
	Uri    APIUpdateUserActionUri
}
type APIUpdateUserActionForm struct {
	NameOptional
	CategoryOptional
	BodyOptional
	EquipmentOptional
	IntroOptional
}
type APIUpdateUserActionUri struct {
	IDRequired
}

// APIGetUserActionsInput /v2/user/actions [GET] 獲取個人動作庫 API
type APIGetUserActionsInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Query  APIGetUserActionsQuery
}
type APIGetUserActionsQuery struct {
	Name      *string `json:"name,omitempty" form:"name" binding:"omitempty,min=1,max=20" example:"槓鈴臥推"`                 //動作名稱(1~20字元)
	Source    *string `json:"source,omitempty" form:"source" binding:"omitempty,action_source" example:"1,3"`             //動作來源(1:平台動作/3:個人動作)
	Category  *string `json:"category,omitempty" form:"category" binding:"omitempty,action_category" example:"1,2,3,4,5"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body      *string `json:"body,omitempty" form:"body" binding:"omitempty,action_body" example:"2,4,6"`                 //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *string `json:"equipment,omitempty" form:"equipment" binding:"omitempty,action_equipment" example:"1,3,5"`  //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	PagingInput
}

// APIDeleteUserActionInput /v2/user/action/{action_id} [DELETE]
type APIDeleteUserActionInput struct {
	UserIDRequired
	Uri APIDeleteUserActionUri
}
type APIDeleteUserActionUri struct {
	IDRequired
}

// APIDeleteUserActionVideoInput /v2/user/action/{action_id}/video
type APIDeleteUserActionVideoInput struct {
	UserIDRequired
	Uri APIDeleteUserActionVideoUri
}
type APIDeleteUserActionVideoUri struct {
	IDRequired
}

// APIGetTrainerActionsInput /v2/trainer/actions [GET] 獲取教練動作庫 API
type APIGetTrainerActionsInput struct {
	UserIDRequired
	Query APIGetUserActionsQuery
}
type APIGetTrainerActionsQuery struct {
	Name      *string `json:"name,omitempty" form:"name" binding:"omitempty,min=1,max=20" example:"槓鈴臥推"`                 //動作名稱(1~20字元)
	Source    *string `json:"source,omitempty" form:"source" binding:"omitempty,action_source" example:"1,2"`             //動作來源(1:平台動作/2:教練)
	Category  *string `json:"category,omitempty" form:"category" binding:"omitempty,action_category" example:"1,2,3,4,5"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body      *string `json:"body,omitempty" form:"body" binding:"omitempty,action_body" example:"2,4,6"`                 //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *string `json:"equipment,omitempty" form:"equipment" binding:"omitempty,action_equipment" example:"1,3,5"`  //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	PagingInput
}

// APICreateTrainerActionInput /v2/trainer/action [POST] 新增教練動作 API
type APICreateTrainerActionInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Cover  *file.Input
	Video  *file.Input
	Form   APICreateUserActionForm
}
type APICreateTrainerActionForm struct {
	NameRequired
	TypeRequired
	CategoryRequired
	BodyRequired
	EquipmentRequired
	IntroRequired
}
