package optional

type YearField struct {
	Year *int `json:"year,omitempty" form:"year" gorm:"column:year" binding:"omitempty,max=2500" example:"2022"` //年份
}
type MonthField struct {
	Month *int `json:"month,omitempty" form:"month" gorm:"column:month" binding:"omitempty,min=1,max=12" example:"12"` //月份
}
type TotalField struct {
	Total *int `json:"total,omitempty" gorm:"column:total" binding:"omitempty" example:"1000"` //當月總註冊人數
}
type MaleField struct {
	Male *int `json:"male,omitempty" gorm:"column:male" binding:"omitempty" example:"600"` //男性註冊人數
}
type FemaleField struct {
	Female *int `json:"female,omitempty" gorm:"column:female" binding:"omitempty" example:"400"` //女性註冊人數
}
type BeginnerField struct {
	Beginner *int `json:"beginner,omitempty" gorm:"column:beginner" binding:"omitempty" example:"250"` //入門用戶註冊人數
}
type IntermediateField struct {
	Intermediate *int `json:"intermediate,omitempty" gorm:"column:intermediate" binding:"omitempty" example:"250"` //中階用戶註冊人數
}
type AdvancedField struct {
	Advanced *int `json:"advanced,omitempty" gorm:"column:advanced" binding:"omitempty" example:"250"` //中高階用戶註冊人數
}
type ExpertField struct {
	Expert *int `json:"expert,omitempty" gorm:"column:expert" binding:"omitempty" example:"250"` //專業用戶註冊人數
}
type Age13to17Field struct {
	Age13to17 *int `json:"age_13_17,omitempty" gorm:"column:age_13_17" binding:"omitempty" example:"100"` //13-17歲註冊人數
}
type Age18to24Field struct {
	Age18to24 *int `json:"age_18_24,omitempty" gorm:"column:age_18_24" binding:"omitempty" example:"150"` //18-24歲註冊人數
}
type Age25to34Field struct {
	Age25to34 *int `json:"age_25_34,omitempty" gorm:"column:age_25_34" binding:"omitempty" example:"250"` //25-34歲註冊人數
}
type Age35to44Field struct {
	Age35to44 *int `json:"age_35_44,omitempty" gorm:"column:age_35_44" binding:"omitempty" example:"200"` //35_44歲註冊人數
}
type Age45to54Field struct {
	Age45to54 *int `json:"age_45_54,omitempty" gorm:"column:age_45_54" binding:"omitempty" example:"150"` //45_54歲註冊人數
}
type Age55to64Field struct {
	Age55to64 *int `json:"age_55_64,omitempty" gorm:"column:age_55_64" binding:"omitempty" example:"100"` //55_64歲註冊人數
}
type Age65UpField struct {
	Age65Up *int `json:"age_65_up,omitempty" gorm:"column:age_65_up" binding:"omitempty" example:"50"` //65+歲註冊人數
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
