package model

type Action struct {
	ID int64  `gorm:"column:id"` //動作id
	CourseID int64 `gorm:"column:course_id"` //課表id
	Name string `gorm:"column:name"` //課表名稱
	Source int `gorm:"column:source"` //動作來源(1:系統動作/2:教練自創動作)
	Type int `gorm:"column:type"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `gorm:"column:category"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `gorm:"column:body"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `gorm:"column:equipment"` //器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)
	Intro string `gorm:"column:intro"` //動作介紹
	Cover string `gorm:"column:cover"` //封面
	Video string `gorm:"column:video"` //動作影片
	CreateAt string `gorm:"column:create_at"` // 創建時間
	UpdateAt string `gorm:"column:update_at"` // 更新時間
}

func (Action) TableName() string {
	return "actions"
}

type CreateActionParam struct {
	Name string `gorm:"column:name"` //課表名稱
	Type int `gorm:"column:type"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `gorm:"column:category"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `gorm:"column:body"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `gorm:"column:equipment"` //器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)
	Intro string `gorm:"column:intro"` //動作介紹
}