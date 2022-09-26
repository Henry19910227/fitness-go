package nullable

type SaleIDField struct {
	SaleID *int64 `json:"sale_id" form:"sale_id" gorm:"column:sale_id" binding:"omitempty" example:"3"` // 銷售 id
}
