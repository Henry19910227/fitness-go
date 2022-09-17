package user_promote_monthly_statistic

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //報表id
}
type YearField struct {
	Year *int `json:"year,omitempty" gorm:"column:year" example:"2022"` //年份
}
type MonthField struct {
	Month *int `json:"month,omitempty" gorm:"column:month" example:"12"` //月份
}
type TotalField struct {
	Total *int `json:"total,omitempty" gorm:"column:total" example:"1000"` //當月總註冊人數
}
type MaleField struct {
	Male *int `json:"male,omitempty" gorm:"column:male" example:"600"` //男性註冊人數
}
type FemaleField struct {
	Female *int `json:"female,omitempty" gorm:"column:female" example:"400"` //女性註冊人數
}
type Exp1to3Field struct {
	Exp1to3 *int `json:"exp_1_3,omitempty" gorm:"column:exp_1_3" example:"100"` //1-3年經驗晉升教練人數
}
type Exp4to6Field struct {
	Exp4to6 *int `json:"exp_4_6,omitempty" gorm:"column:exp_4_6" example:"100"` //4-6年經驗晉升教練人數
}
type Exp7to10Field struct {
	Exp7to10 *int `json:"exp_7_10,omitempty" gorm:"column:exp_7_10" example:"200"` //7-10年經驗晉升教練人數
}
type Exp11to15Field struct {
	Exp11to15 *int `json:"exp_11_15,omitempty" gorm:"column:exp_11_15" example:"200"` //11-15年經驗晉升教練人數
}
type Exp16to19Field struct {
	Exp16to19 *int `json:"exp_16_19,omitempty" gorm:"column:exp_16_19" example:"200"` //16-19年經驗晉升教練人數
}
type Exp20upField struct {
	Exp20up *int `json:"exp_20_up,omitempty" gorm:"column:exp_20_up" example:"200"` //20+年經驗晉升教練人數
}
type Age13to17Field struct {
	Age13to17 *int `json:"age_13_17,omitempty" gorm:"column:age_13_17" example:"100"` //13-17歲註冊人數
}
type Age18to24Field struct {
	Age18to24 *int `json:"age_18_24,omitempty" gorm:"column:age_18_24" example:"150"` //18-24歲註冊人數
}
type Age25to34Field struct {
	Age25to34 *int `json:"age_25_34,omitempty" gorm:"column:age_25_34" example:"250"` //25-34歲註冊人數
}
type Age35to44Field struct {
	Age35to44 *int `json:"age_35_44,omitempty" gorm:"column:age_35_44" example:"200"` //35_44歲註冊人數
}
type Age45to54Field struct {
	Age45to54 *int `json:"age_45_54,omitempty" gorm:"column:age_45_54" example:"150"` //45_54歲註冊人數
}
type Age55to64Field struct {
	Age55to64 *int `json:"age_55_64,omitempty" gorm:"column:age_55_64" example:"100"` //55_64歲註冊人數
}
type Age65UpField struct {
	Age65Up *int `json:"age_65_up,omitempty" gorm:"column:age_65_up" example:"50"` //65+歲註冊人數
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	YearField
	MonthField
	TotalField
	MaleField
	FemaleField
	Exp1to3Field
	Exp4to6Field
	Exp7to10Field
	Exp11to15Field
	Exp16to19Field
	Exp20upField
	Age13to17Field
	Age18to24Field
	Age25to34Field
	Age35to44Field
	Age45to54Field
	Age55to64Field
	Age65UpField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "user_promote_monthly_statistic"
}

