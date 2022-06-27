package body_record

type IDRequired struct {
	ID *int64 `json:"id" binding:"required" example:"1"` //主鍵id
}
type UserIDRequired struct {
	UserID *int64 `json:"user_id" binding:"required" example:"10001"` //用戶id
}
type RecordTypeRequired struct {
	RecordType *int `json:"record_type" binding:"required,oneof=1 2 3 4 5 6 7 8 9 10 11 12" example:"2"` //紀錄類型(1:體重紀錄/2:體脂紀錄/3:胸圍紀錄/4:腰圍紀錄/5:臀圍紀錄/6:身高紀錄/7:臂圍紀錄/8:小臂圍紀錄/9:肩圍紀錄/10:大腿圍紀錄/11:小腿圍紀錄/12:頸圍紀錄
}
type ValueRequired struct {
	Value *float64 `json:"value" binding:"required" example:"18.5"` //數值
}
