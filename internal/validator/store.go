package validator

type SearchCourseProductsQuery struct {
	Name *string `form:"name" binding:"omitempty,min=1,max=20" example:"增肌課表"` //課表名稱(1~20字元)
	OrderType *string `form:"order_type" binding:"omitempty,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)-單選
	Score *int `form:"score" binding:"omitempty,min=1,max=5" example:"5"` // 評價(1~5分)-單選
	Level []int `form:"level" binding:"omitempty,level_inspect" example:"3"` // 強度(1:初級/2:中級/3:中高級/4:高級)-複選
	Category []int `form:"category" binding:"omitempty,category_inspect" example:"3"` // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選
	Suit []int `form:"suit" binding:"omitempty,suit_inspect" example:"7"` // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選
	Equipment []int `form:"equipment" binding:"omitempty,equipment_inspect" example:"5"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選
	Place []int `form:"place" binding:"omitempty,place_inspect" example:"1,2,3"` // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選
	TrainTarget []int `form:"train_target" binding:"omitempty,trainer_target_inspect" example:"1"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選
	BodyTarget []int `form:"body_target" binding:"omitempty,body_target_inspect" example:"2"` // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選
	SaleType []int `form:"sale_type" binding:"omitempty,sale_type_inspect" example:"1"` // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)-複選
	TrainerSex []string `form:"trainer_sex" binding:"omitempty,trainer_sex_inspect" example:"m"` // 教練性別(m:男性/f:女性)-複選
	TrainerSkill  []int  `form:"trainer_skill" binding:"omitempty,trainer_skill_inspect" example:"5"`  // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	Page int `form:"page" binding:"required,min=1" example:"1"` // 頁數
	Size int `form:"size" binding:"required,min=1" example:"5"` // 筆數
}