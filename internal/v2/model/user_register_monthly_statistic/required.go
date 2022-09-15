package user_register_monthly_statistic

type IDRequired struct {
	ID int64 `json:"id" example:"1"` //報表id
}
type YearRequired struct {
	Year int `json:"year" form:"year" binding:"required,max=2500" example:"2022"` //年份
}
type MonthRequired struct {
	Month int `json:"month" form:"month" binding:"required,min=1,max=12" example:"12"` //月份
}
type TotalRequired struct {
	Total int `json:"total" example:"1000"` //當月總註冊人數
}
type MaleRequired struct {
	Male int `json:"male" example:"600"` //男性註冊人數
}
type FemaleRequired struct {
	Female int `json:"female" example:"400"` //女性註冊人數
}
type BeginnerRequired struct {
	Beginner int `json:"beginner" example:"400"` //入門用戶註冊人數
}
type IntermediateRequired struct {
	Intermediate int `json:"intermediate" example:"400"` //中階用戶註冊人數
}
type AdvancedRequired struct {
	Advanced int `json:"advanced" example:"400"` //中高階用戶註冊人數
}
type ExpertRequired struct {
	Expert int `json:"expert" example:"400"` //專業用戶註冊人數
}
type Age13to17Required struct {
	Age13to17 int `json:"age_13_17" example:"100"` //13-17歲註冊人數
}
type Age18to24Required struct {
	Age18to24 int `json:"age_18_24" example:"150"` //18-24歲註冊人數
}
type Age25to34Required struct {
	Age25to34 int `json:"age_25_34" example:"250"` //25-34歲註冊人數
}
type Age35to44Required struct {
	Age35to44 int `json:"age_35_44" example:"200"` //35_44歲註冊人數
}
type Age45to54Required struct {
	Age45to54 int `json:"age_45_54" example:"150"` //45_54歲註冊人數
}
type Age55to64Required struct {
	Age55to64 int `json:"age_55_64" example:"100"` //55_64歲註冊人數
}
type Age65UpRequired struct {
	Age65Up int `json:"age_65_up" example:"50"` //65+歲註冊人數
}
