package order_by

type OrderFieldField struct {
	OrderField string `form:"order_field" binding:"required" example:"create_at"` // 排序欄位
}

type OrderTypeField struct {
	OrderType string `form:"order_type" binding:"required,oneof=ASC DESC" example:"DESC"` // 排序類型 (ASC:由低到高/DESC:由高到低)
}
