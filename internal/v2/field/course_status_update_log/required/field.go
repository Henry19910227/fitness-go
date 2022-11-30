package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"2"` // 紀錄 id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
type CourseStatusField struct {
	CourseStatus int `json:"course_status" form:"course_status" gorm:"column:course_status" binding:"required,oneof=1 2 3 4 5" example:"3"` // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
}
type CommentField struct {
	Comment string `json:"comment" gorm:"column:comment" binding:"required,max=500" example:"課表內容太爛必須退審"` // 註解
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
