package validator

type PagingQuery struct {
	Page int `form:"page" binding:"required,min=1" example:"1"`
	Size int `form:"size" binding:"required,min=1" example:"5"`
}

type OrderByQuery struct {
	OrderType  *string `form:"order_type" binding:"omitempty,oneof=ASC DESC" example:"DESC"` // 排序類型 (ASC:由低到高/DESC:由高到低)
	OrderField *string `form:"order_field" binding:"omitempty" example:"create_at"`          // 排序欄位
}
