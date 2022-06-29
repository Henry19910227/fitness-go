package body_record

type IDOptional struct {
	ID *int64 `json:"id,omitempty" example:"1"` //紀錄id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" example:"10001"` //用戶id
}
type RecordTypeOptional struct {
	RecordType *int `json:"record_type,omitempty" example:"2"` //紀錄類型(1:體重紀錄/2:體脂紀錄/3:胸圍紀錄/4:腰圍紀錄/5:臀圍紀錄/6:身高紀錄/7:臂圍紀錄/8:小臂圍紀錄/9:肩圍紀錄/10:大腿圍紀錄/11:小腿圍紀錄/12:頸圍紀錄
}
type ValueOptional struct {
	Value *float64 `json:"value,omitempty" binding:"omitempty" example:"18.5"` //數值
}
