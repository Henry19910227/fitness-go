package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"2"` // 紀錄 id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
type CourseStatusField struct {
	CourseStatus *int `json:"course_status,omitempty" form:"course_status" gorm:"column:course_status" binding:"omitempty,oneof=1 2 3 4 5" example:"3"` // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
}
type CommentField struct {
	Comment *string `json:"comment,omitempty" gorm:"column:comment" binding:"omitempty,max=500" example:"課表內容太爛必須退審"` // 註解
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
