package required

type YearField struct {
	Year int `json:"year" gorm:"column:year" binding:"required" example:"2022"` //年份
}
type MonthField struct {
	Month int `json:"month" gorm:"column:month" binding:"required" example:"12"` //月份
}
type TotalField struct {
	Total int `json:"total" gorm:"column:total" binding:"required" example:"1000"` //當月總訂閱人數
}
type MaleField struct {
	Male int `json:"male" gorm:"column:male" binding:"required" example:"600"` //男性訂閱人數
}
type FemaleField struct {
	Female int `json:"female" gorm:"column:female" binding:"required" example:"400"` //女性訂閱人數
}
type Age13to17Field struct {
	Age13to17 int `json:"age_13_17" gorm:"column:age_13_17" binding:"required" example:"100"` //13-17歲訂閱人數
}
type Age18to24Field struct {
	Age18to24 int `json:"age_18_24" gorm:"column:age_18_24" binding:"required" example:"150"` //18-24歲訂閱人數
}
type Age25to34Field struct {
	Age25to34 int `json:"age_25_34" gorm:"column:age_25_34" binding:"required" example:"250"` //25-34歲訂閱人數
}
type Age35to44Field struct {
	Age35to44 int `json:"age_35_44" gorm:"column:age_35_44" binding:"required" example:"200"` //35_44歲訂閱人數
}
type Age45to54Field struct {
	Age45to54 int `json:"age_45_54" gorm:"column:age_45_54" binding:"required" example:"150"` //45_54歲訂閱人數
}
type Age55to64Field struct {
	Age55to64 int `json:"age_55_64" gorm:"column:age_55_64" binding:"required" example:"100"` //55_64歲訂閱人數
}
type Age65UpField struct {
	Age65Up int `json:"age_65_up" gorm:"column:age_65_up" binding:"required" example:"50"` //65+歲訂閱人數
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
