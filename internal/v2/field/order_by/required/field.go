package required

type OrderFieldField struct {
	OrderField string `json:"order_field" form:"order_field" binding:"required" example:"create_at"` // 排序欄位
}

type OrderTypeField struct {
	OrderType string `json:"order_type" form:"order_type" binding:"required,oneof=ASC DESC" example:"DESC"` // 排序類型 (ASC:由低到高/DESC:由高到低)
}
