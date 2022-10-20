package optional

type OrderFieldField struct {
	OrderField *string `json:"order_field,omitempty" form:"order_field" binding:"omitempty" example:"create_at"` // 排序欄位
}

type OrderTypeField struct {
	OrderType *string `json:"order_type,omitempty" form:"order_type" binding:"omitempty,oneof=ASC DESC" example:"DESC"` // 排序類型 (ASC:由低到高/DESC:由高到低)
}
