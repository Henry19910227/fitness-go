package order

type CMSUserOrdersAPI struct {
	ID          string `gorm:"column:id" json:"id" example:"20220425201813456293"`              // 訂單id
	CourseName  string `gorm:"column:course_name" json:"course_name" example:"Henry課表"`         // 課表名稱
	TrainerName string `gorm:"column:trainer_name" json:"trainer_name" example:"henry"`         // 教練名稱
	CreateAt    string `gorm:"column:create_at" json:"create_at" example:"2022-04-25 20:18:13"` // 創建時間
}
