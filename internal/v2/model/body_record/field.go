package body_record

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //主鍵id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type RecordTypeField struct {
	RecordType *int `json:"record_type,omitempty" gorm:"column:record_type" example:"2"` //紀錄類型(1:體重紀錄/2:體脂紀錄/3:胸圍紀錄/4:腰圍紀錄/5:臀圍紀錄/6:身高紀錄/7:臂圍紀錄/8:小臂圍紀錄/9:肩圍紀錄/10:大腿圍紀錄/11:小腿圍紀錄/12:頸圍紀錄
}
type ValueField struct {
	Value *float64 `json:"value,omitempty" gorm:"column:value" example:"18.5"` //數值
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	RecordTypeField
	ValueField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "body_records"
}
